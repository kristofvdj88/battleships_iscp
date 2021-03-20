// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// statemgr package implements object which is responsible for the smart contract
// ledger state to be synchronized and validated
package statemgr

import (
	"time"

	valuetransaction "github.com/iotaledger/goshimmer/dapps/valuetransfers/packages/transaction"
	"github.com/iotaledger/hive.go/logger"
	"github.com/iotaledger/wasp/packages/chain"
	"github.com/iotaledger/wasp/packages/hashing"
	"github.com/iotaledger/wasp/packages/sctransaction"
	"github.com/iotaledger/wasp/packages/state"
	"github.com/iotaledger/wasp/packages/util"
)

type stateManager struct {
	chain chain.Chain

	// becomes true after initially loaded state is validated.
	// after that it is always true
	solidStateValid bool

	// flag pingPong[idx] if ping-pong message was received from the peer idx
	pingPong              []bool
	deadlineForPongQuorum time.Time

	// pending batches of state updates are candidates to confirmation by the state transaction
	// which leads to the state transition
	// the map key is hash of the variable state which is a result of applying the
	// block of state updates to the solid variable state
	pendingBlocks map[hashing.HashValue]*pendingBlock

	// last variable state stored in the database
	// it may be nil at bootstrap when origin variable state is calculated
	solidState state.VirtualState

	// state transaction with +1 state index from the state index of solid variable state
	// it may be nil if does not exist or not fetched yet
	nextStateTransaction *sctransaction.Transaction

	// state transaction which approves current state
	approvingTransaction *sctransaction.Transaction

	// was state transition message of the current state sent to the consensus operator
	consensusNotifiedOnStateTransition bool

	// largest state index evidenced by other messages. If this index is more than 1 step ahead
	// of the solid variable state, it means the state of the smart contract in the current node
	// falls behind the state of the smart contract, i.e. it is not synced
	largestEvidencedStateIndex uint32

	// the timeout deadline for sync inquiries
	syncMessageDeadline time.Time

	// current block being synced
	syncedBatch *syncedBatch

	// for the pseudo-random sequence of peers
	permutation *util.Permutation16

	// logger
	log *logger.Logger

	// Channels for accepting external events.
	evidenceStateIndexCh         chan uint32
	eventStateIndexPingPongMsgCh chan *chain.StateIndexPingPongMsg
	eventGetBlockMsgCh           chan *chain.GetBlockMsg
	eventBlockHeaderMsgCh        chan *chain.BlockHeaderMsg
	eventStateUpdateMsgCh        chan *chain.StateUpdateMsg
	eventStateTransactionMsgCh   chan *chain.StateTransactionMsg
	eventPendingBlockMsgCh       chan chain.PendingBlockMsg
	eventTimerMsgCh              chan chain.TimerTick
	closeCh                      chan bool
}

type syncedBatch struct {
	msgCounter   uint16
	stateIndex   uint32
	stateUpdates []state.StateUpdate
	stateTxId    valuetransaction.ID
}

type pendingBlock struct {
	// block of state updates, not validated yet
	block state.Block
	// resulting variable state after applied the block to the solidState
	nextState state.VirtualState
	// state transaction request deadline. For committed batches only
	stateTransactionRequestDeadline time.Time
}

func New(c chain.Chain, log *logger.Logger) chain.StateManager {
	ret := &stateManager{
		chain:                        c,
		pingPong:                     make([]bool, c.Size()),
		pendingBlocks:                make(map[hashing.HashValue]*pendingBlock),
		permutation:                  util.NewPermutation16(c.NumPeers(), nil),
		log:                          log.Named("s"),
		evidenceStateIndexCh:         make(chan uint32),
		eventStateIndexPingPongMsgCh: make(chan *chain.StateIndexPingPongMsg),
		eventGetBlockMsgCh:           make(chan *chain.GetBlockMsg),
		eventBlockHeaderMsgCh:        make(chan *chain.BlockHeaderMsg),
		eventStateUpdateMsgCh:        make(chan *chain.StateUpdateMsg),
		eventStateTransactionMsgCh:   make(chan *chain.StateTransactionMsg),
		eventPendingBlockMsgCh:       make(chan chain.PendingBlockMsg),
		eventTimerMsgCh:              make(chan chain.TimerTick),
		closeCh:                      make(chan bool),
	}
	go ret.initLoadState()

	return ret
}

func (sm *stateManager) Close() {
	close(sm.closeCh)
}

// initial loading of the solid state
func (sm *stateManager) initLoadState() {
	var err error
	var batch state.Block
	var stateExists bool

	sm.solidState, batch, stateExists, err = state.LoadSolidState(sm.chain.ID())
	if err != nil {
		sm.log.Errorf("initLoadState: %v", err)
		sm.chain.Dismiss()
		return
	}

	if stateExists {
		// state loaded, will be waiting for it to be confirmed from the tangle
		sm.addPendingBlock(batch)
		sm.largestEvidencedStateIndex = sm.solidState.BlockIndex()

		h := sm.solidState.Hash()
		txh := batch.StateTransactionID()
		sm.log.Debugw("solid state has been loaded",
			"state index", sm.solidState.BlockIndex(),
			"state hash", h.String(),
			"approving tx", txh.String(),
		)
	} else {
		// pre-origin state. Origin block is empty block.
		// Will be waiting for the origin transaction to arrive
		sm.addPendingBlock(state.MustNewOriginBlock(sm.chain.Color()))

		sm.log.Info("solid state does not exist: WAITING FOR THE ORIGIN TRANSACTION")
	}

	sm.chain.SetReadyStateManager() // Open msg queue for the committee
	sm.recvLoop()                   // Start to process external events.
}

func (sm *stateManager) recvLoop() {
	for {
		select {
		case msg, ok := <-sm.evidenceStateIndexCh:
			if ok {
				sm.evidenceStateIndex(msg)
			}
		case msg, ok := <-sm.eventStateIndexPingPongMsgCh:
			if ok {
				sm.eventStateIndexPingPongMsg(msg)
			}
		case msg, ok := <-sm.eventGetBlockMsgCh:
			if ok {
				sm.eventGetBlockMsg(msg)
			}
		case msg, ok := <-sm.eventBlockHeaderMsgCh:
			if ok {
				sm.eventBlockHeaderMsg(msg)
			}
		case msg, ok := <-sm.eventStateUpdateMsgCh:
			if ok {
				sm.eventStateUpdateMsg(msg)
			}
		case msg, ok := <-sm.eventStateTransactionMsgCh:
			if ok {
				sm.eventStateTransactionMsg(msg)
			}
		case msg, ok := <-sm.eventPendingBlockMsgCh:
			if ok {
				sm.eventPendingBlockMsg(msg)
			}
		case msg, ok := <-sm.eventTimerMsgCh:
			if ok {
				sm.eventTimerMsg(msg)
			}
		case <-sm.closeCh:
			return
		}
	}
}
