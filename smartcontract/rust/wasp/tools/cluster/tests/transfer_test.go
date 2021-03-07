package tests

import (
	"github.com/iotaledger/wasp/packages/coretypes/cbalances"
	"testing"
	"time"

	"github.com/iotaledger/goshimmer/dapps/valuetransfers/packages/balance"
	"github.com/iotaledger/wasp/client/chainclient"
	"github.com/iotaledger/wasp/packages/coretypes"
	"github.com/iotaledger/wasp/packages/testutil"
	"github.com/iotaledger/wasp/packages/vm/core/accounts"
)

func TestDepositWithdraw(t *testing.T) {
	setup(t, "test_cluster")

	chain, err := clu.DeployDefaultChain()
	check(err, t)

	testOwner := wallet.WithIndex(1)
	mySigScheme := testOwner.SigScheme()
	myAddress := testOwner.Address()

	err = requestFunds(clu, myAddress, "myAddress")
	check(err, t)
	if !clu.VerifyAddressBalances(myAddress, testutil.RequestFundsAmount, map[balance.Color]int64{
		balance.ColorIOTA: testutil.RequestFundsAmount,
	}, "myAddress begin") {
		t.Fail()
	}
	if !clu.VerifyAddressBalances(chain.OriginatorAddress(), testutil.RequestFundsAmount-2, map[balance.Color]int64{
		balance.ColorIOTA: testutil.RequestFundsAmount - 2,
	}, "originatorAddress begin") {
		t.Fail()
	}
	if !clu.VerifyAddressBalances(chain.ChainAddress(), 2, map[balance.Color]int64{
		chain.Color:       1,
		balance.ColorIOTA: 1,
	}, "chainAddress begin") {
		t.Fail()
	}
	checkLedger(t, chain)

	myAgentID := coretypes.NewAgentIDFromAddress(*myAddress)
	origAgentId := coretypes.NewAgentIDFromAddress(*chain.OriginatorAddress())

	checkBalanceOnChain(t, chain, origAgentId, balance.ColorIOTA, 1)
	checkBalanceOnChain(t, chain, myAgentID, balance.ColorIOTA, 0)
	checkLedger(t, chain)

	// deposit some iotas to the chain
	depositIotas := int64(42)
	chClient := chainclient.New(clu.Level1Client(), clu.WaspClient(0), chain.ChainID, mySigScheme)
	reqTx, err := chClient.PostRequest(accounts.Interface.Hname(), coretypes.Hn(accounts.FuncDeposit), chainclient.PostRequestParams{
		Transfer: cbalances.NewIotasOnly(depositIotas),
	})
	check(err, t)
	err = chain.CommitteeMultiClient().WaitUntilAllRequestsProcessed(reqTx, 30*time.Second)
	check(err, t)
	checkLedger(t, chain)
	checkBalanceOnChain(t, chain, myAgentID, balance.ColorIOTA, depositIotas+1) // 1 iota from request
	checkBalanceOnChain(t, chain, origAgentId, balance.ColorIOTA, 1)

	if !clu.VerifyAddressBalances(myAddress, testutil.RequestFundsAmount-depositIotas-1, map[balance.Color]int64{
		balance.ColorIOTA: testutil.RequestFundsAmount - depositIotas - 1,
	}, "myAddress after deposit") {
		t.Fail()
	}

	// withdraw iotas back
	reqTx3, err := chClient.PostRequest(accounts.Interface.Hname(), coretypes.Hn(accounts.FuncWithdrawToAddress))
	check(err, t)
	err = chain.CommitteeMultiClient().WaitUntilAllRequestsProcessed(reqTx3, 30*time.Second)
	check(err, t)

	check(err, t)
	checkLedger(t, chain)
	checkBalanceOnChain(t, chain, myAgentID, balance.ColorIOTA, 0)

	if !clu.VerifyAddressBalances(myAddress, testutil.RequestFundsAmount, map[balance.Color]int64{
		balance.ColorIOTA: testutil.RequestFundsAmount,
	}, "myAddress after withdraw") {
		t.Fail()
	}
}
