package libtest

import (
	"fmt"
	"testing"

	"github.com/brunoamancio/IOTA-SmartContracts/tests/testutils"
	"github.com/brunoamancio/IOTA-SmartContracts/tests/testutils/testconstants"
	"github.com/iotaledger/goshimmer/dapps/valuetransfers/packages/address/signaturescheme"
	"github.com/iotaledger/goshimmer/dapps/valuetransfers/packages/balance"
	"github.com/iotaledger/wasp/packages/coretypes"
	"github.com/iotaledger/wasp/packages/kv/codec"
	"github.com/iotaledger/wasp/packages/solo"
	"github.com/stretchr/testify/require"
)

//  -----------------------------------------------  //
//  See code samples in Tests/testutils/codesamples  //
//  -----------------------------------------------  //

func TestCreateGame(t *testing.T, userWallet signaturescheme.SignatureScheme) {
	contractWasmFilePath := testutils.MustGetContractWasmFilePath(t, testconstants.ContractName) // You can use if file is in SmartContract/pkg

	env := solo.New(t, testconstants.Debug, testconstants.StackTrace)
	chainName := testconstants.ContractName + "Chain"
	chain := env.NewChain(nil, chainName)

	// Uploads wasm of SC and deploys it into chain
	err := chain.DeployWasmContract(nil, testconstants.ContractName, contractWasmFilePath)
	require.NoError(t, err)

	// Loads contract information
	contract, err := chain.FindContract(testconstants.ContractName)
	require.NoError(t, err)
	require.NotNil(t, contract)
	require.Equal(t, testconstants.ContractName, contract.Name)

	// global ID of the deployed contract
	contractID := coretypes.NewContractID(chain.ChainID, coretypes.Hn(testconstants.ContractName))
	// contract id in the form of the agent ID
	contractAgentID := coretypes.NewAgentIDFromContractID(contractID)

	// Create a request to the "create_game" function endpoint of the SC and post the request to the L1 Tangle
	req := solo.NewCallParams(testconstants.ContractName, "create_game", "createGameRequestKey", testconstants.CreateGameRequest).
			    WithTransfer(balance.ColorIOTA, 100)
	_, err = chain.PostRequest(req, userWallet)
	require.NoError(t, err)

	// Assert if the stake of 100 tokens is on chain, in the account of the contractAgentID
	chain.AssertAccountBalance(contractAgentID, balance.ColorIOTA, 100)	

	// Create a request to the "get_game" view endpoint of the SC and post the request to the L2 chain
	res, err := chain.CallView(testconstants.ContractName, "get_game", "getGameRequestKey", testconstants.GetGameRequest)	
	require.NoError(t, err)
	returnedString, exists, err := codec.DecodeString(res.MustGet("gameStateResponseKey"))
	require.NoError(t, err)
	require.True(t, exists)
	fmt.Println(returnedString)
}
