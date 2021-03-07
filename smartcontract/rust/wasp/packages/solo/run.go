// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package solo

import (
	"fmt"
	"github.com/iotaledger/goshimmer/dapps/waspconn/packages/waspconn"
	"github.com/iotaledger/wasp/packages/coretypes"
	"github.com/iotaledger/wasp/packages/hashing"
	"github.com/iotaledger/wasp/packages/kv/dict"
	"github.com/iotaledger/wasp/packages/sctransaction"
	"github.com/iotaledger/wasp/packages/state"
	"github.com/iotaledger/wasp/packages/vm"
	"github.com/iotaledger/wasp/packages/vm/runvm"
	"github.com/stretchr/testify/require"
	"strings"
	"sync"
)

func (ch *Chain) runBatch(batch []vm.RequestRefWithFreeTokens, trace string) (dict.Dict, error) {
	ch.Log.Debugf("runBatch ('%s')", trace)
	ch.runVMMutex.Lock()
	defer ch.runVMMutex.Unlock()

	// solidify arguments
	for _, reqRef := range batch {
		if ok, err := reqRef.RequestSection().SolidifyArgs(ch.Env.registry); err != nil || !ok {
			return nil, fmt.Errorf("solo inconsistency: failed to solidify request args")
		}
	}

	task := &vm.VMTask{
		Processors:         ch.proc,
		ChainID:            ch.ChainID,
		Color:              ch.ChainColor,
		Entropy:            hashing.RandomHash(nil),
		ValidatorFeeTarget: ch.ValidatorFeeTarget,
		Balances:           waspconn.OutputsToBalances(ch.Env.utxoDB.GetAddressOutputs(ch.ChainAddress)),
		Requests:           batch,
		Timestamp:          ch.Env.LogicalTime().UnixNano(),
		VirtualState:       ch.State.Clone(),
		Log:                ch.Log,
	}
	var err error
	var wg sync.WaitGroup
	var callRes dict.Dict
	var callErr error
	task.OnFinish = func(callResult dict.Dict, callError error, err error) {
		require.NoError(ch.Env.T, err)
		callRes = callResult
		callErr = callError
		wg.Done()
	}

	wg.Add(1)
	err = runvm.RunComputationsAsync(task)
	require.NoError(ch.Env.T, err)

	wg.Wait()
	task.ResultTransaction.Sign(ch.ChainSigScheme)

	ch.settleStateTransition(task.VirtualState, task.ResultBlock, task.ResultTransaction)
	return callRes, callErr
}

func (ch *Chain) settleStateTransition(newState state.VirtualState, block state.Block, stateTx *sctransaction.Transaction) {
	err := ch.Env.utxoDB.AddTransaction(stateTx.Transaction)
	require.NoError(ch.Env.T, err)

	err = newState.ApplyBlock(block)
	require.NoError(ch.Env.T, err)

	err = newState.CommitToDb(block)
	require.NoError(ch.Env.T, err)

	prevBlockIndex := ch.StateTx.MustState().BlockIndex()

	ch.StateTx = stateTx
	ch.State = newState

	ch.Log.Infof("state transition #%d --> #%d. Requests in the block: %d. Posted: %d",
		prevBlockIndex, ch.State.BlockIndex(), len(block.RequestIDs()), len(ch.StateTx.Requests()))
	ch.Log.Debugf("Batch processed: %s", batchShortStr(block.RequestIDs()))

	ch.Env.ClockStep()

	// dispatch requests among chains
	ch.Env.glbMutex.Lock()
	defer ch.Env.glbMutex.Unlock()

	reqRefByChain := make(map[coretypes.ChainID][]sctransaction.RequestRef)
	for i, rsect := range ch.StateTx.Requests() {
		chid := rsect.Target().ChainID()
		_, ok := reqRefByChain[chid]
		if !ok {
			reqRefByChain[chid] = make([]sctransaction.RequestRef, 0)
		}
		reqRefByChain[chid] = append(reqRefByChain[chid], sctransaction.RequestRef{
			Tx:    stateTx,
			Index: uint16(i),
		})
	}
	for chid, reqs := range reqRefByChain {
		chain, ok := ch.Env.chains[chid]
		if !ok {
			ch.Log.Infof("dispatching requests. Unknown chain: %s", chid.String())
			continue
		}
		chain.chPosted.Add(len(reqs))
		for _, reqRef := range reqs {
			chain.chInRequest <- reqRef
		}
	}
}

func batchShortStr(reqIds []*coretypes.RequestID) string {
	ret := make([]string, len(reqIds))
	for i, r := range reqIds {
		ret[i] = r.Short()
	}
	return fmt.Sprintf("[%s]", strings.Join(ret, ","))
}
