package libtest

import (
	"testing"

	"github.com/brunoamancio/IOTA-SmartContracts/tests/testutils"
	"github.com/brunoamancio/IOTA-SmartContracts/tests/testutils/testconstants"
	notsolo "github.com/brunoamancio/NotSolo"
	"github.com/iotaledger/goshimmer/dapps/valuetransfers/packages/balance"
	"github.com/stretchr/testify/require"
)

//  -----------------------------------------------  //
//  See code samples in Tests/testutils/codesamples  //
//  -----------------------------------------------  //

func TestCreateGame(t *testing.T) {
	contractWasmFilePath := testutils.MustGetContractWasmFilePath(t, testconstants.ContractName) // You can use if file is in SmartContract/pkg

	// Setup Solo environment to create SC chain
	notSolo := notsolo.New(t)
	chainName := testconstants.ContractName + "Chain"
	chain := notSolo.Chain.NewChain(nil, chainName)

	// Uploads wasm of SC and deploys it into chain
	notSolo.Chain.DeployWasmContract(chain, nil, testconstants.ContractName, contractWasmFilePath)

	// Loads contract information
	contract, err := notSolo.Chain.GetContractRecord(chain, testconstants.ContractName)
	require.NoError(t, err)
	require.NotNil(t, contract)
	require.Equal(t, testconstants.ContractName, contract.Name)

	// Create a user's wallet (private key) and request 1337 iotas from the faucet.
	// It corresponds to L1 address
	userWallet := notSolo.SigScheme.NewSignatureSchemeWithFunds()

	// Create a request to the "create_game" function endpoint of the SC and post the request (to the L1 Tangle)
	notSolo.Request.MustPostWithTransfer(userWallet, balance.ColorIOTA, 100,
		chain, testconstants.ContractName, "create_game",
		"createGameRequestKey", testconstants.CreateGameRequest)

	// Assert if the stake of 100 tokens is on chain, in the account of the contract
	notSolo.Chain.RequireContractBalance(chain, testconstants.ContractName, balance.ColorIOTA, 100)

	// Post a request to the "get_game" view endpoint of the SC (to the L2 chain). A Responce will be received
	response := notSolo.Request.MustView(chain, testconstants.ContractName, "get_game", "getGameRequestKey", testconstants.GetGameRequest)
	gameStateResponseKey := notSolo.Data.MustGetString(response["gameStateResponseKey"])
	t.Logf(gameStateResponseKey)
}
