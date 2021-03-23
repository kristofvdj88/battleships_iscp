package testconstants

const (
	// ContractName is defined in smartcontract/rust/Cargo.toml > package > name
	ContractName = "battleship_iscp"
	// Debug is used by Solo. 'true' for logging level 'debug', otherwise 'info'
	Debug = false
	// StackTrace is used by Solo. 'true' if stack trace must be printed in case of errors
	StackTrace = false

	/* INTERESTING FACT: Calls to a smart contract require 1 EXTRA iota token to be sent to the chain it is located in.
	   It is colored with the chain's color and corresponds to the request. That is how the protocol locates the backlog of
	   requests to be processed. Basically, it works as a flag. After the request is processed, the token is uncolored
	   and sent to the chain owner's account in the chain.
	*/
	IotaTokensConsumedByRequest = 1

	/* INTERESTING FACT: Creating a chain requires 2 iota tokens. They are colored with the chain's color,
	   1 is sent to the chain's address in the value tangle, and the other is used exactly as iotaTokensConsumedByRequest.
	*/
	IotaTokensConsumedByChain = 1

	// Used to fund address in NewSignatureSchemeWithFunds. // Defined in iotaledger/wasp/packages/testutiltestutil.RequestFundsAmount.
	InitialWalletFunds = 1337

	// AccountsContractName sets the name of the Accounts contract, which is a root contract present in every chain
	AccountsContractName = "accounts"

	CreateGameRequest = "{\"player_name\": \"Admiral Earl Howe\"}"
	JoinGameRequest   = "{\"game_id\": \"001\", \"player_name\": \"Jeanbon Saint-André\"}"
	InitFieldRequest1 = "{\"game_id\": \"001\", \"field\": [[Empty, Empty, Bound, Bound, Bound, Empty, Bound, Bound, Bound, Empty, Empty, Empty], [Empty, Empty, Bound, Ship, Bound, Empty, Bound, Ship, Bound, Empty, Empty, Empty], [Empty, Empty, Bound, Ship, Bound, Bound, Bound, Ship, Bound, Empty, Empty, Empty], [Empty, Empty, Bound, Ship, Bound, Ship, Bound, Bound, Bound, Bound, Bound, Bound], [Empty, Empty, Bound, Ship, Bound, Ship, Bound, Bound, Ship, Bound, Ship, Bound], [Empty, Empty, Bound, Bound, Bound, Ship, Bound, Bound, Bound, Bound, Bound, Bound], [Empty, Bound, Bound, Bound, Bound, Bound, Bound, Bound, Bound, Bound, Ship, Bound], [Empty, Bound, Ship, Bound, Bound, Bound, Bound, Ship, Bound, Bound, Bound, Bound], [Empty, Bound, Ship, Bound, Ship, Bound, Bound, Ship, Bound, Bound, Bound, Empty], [Empty, Bound, Ship, Bound, Bound, Bound, Bound, Bound, Bound, Ship, Bound, Empty], [Empty, Bound, Bound, Bound, Empty, Empty, Empty, Empty, Bound, Ship, Bound, Empty], [Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Bound, Bound, Bound, Empty]]}"
	InitFieldRequest2 = "{\"game_id\": \"001\", \"field\": [[Empty, Bound, Bound, Bound, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty], [Empty, Bound, Ship, Bound, Bound, Bound, Bound, Bound, Empty, Bound, Bound, Bound], [Empty, Bound, Ship, Bound, Ship, Bound, Ship, Bound, Bound, Bound, Ship, Bound], [Bound, Bound, Bound, Bound, Ship, Bound, Ship, Bound, Ship, Bound, Bound, Bound], [Bound, Ship, Bound, Bound, Bound, Bound, Ship, Bound, Bound, Bound, Empty, Empty], [Bound, Bound, Bound, Bound, Bound, Bound, Bound, Bound, Bound, Bound, Bound, Empty], [Empty, Bound, Ship, Bound, Ship, Bound, Bound, Bound, Bound, Ship, Bound, Empty], [Empty, Bound, Ship, Bound, Bound, Bound, Bound, Ship, Bound, Ship, Bound, Empty], [Empty, Bound, Ship, Bound, Empty, Empty, Bound, Ship, Bound, Ship, Bound, Empty], [Empty, Bound, Bound, Bound, Empty, Empty, Bound, Bound, Bound, Ship, Bound, Empty], [Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Bound, Bound, Bound, Empty], [Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty]]}"
	MakeMoveRequest   = ""
	QuitGameRequest   = "{\"game_id\": \"001\"}"
	GetGameRequest    = "{\"game_id\": \"001\"}"
)
