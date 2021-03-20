package testconstants

const (
	// ContractName is defined in smartcontract/rust/Cargo.toml > package > name
	ContractName = "battleship_iscp"
	// Debug is used by Solo. 'true' for logging level 'debug', otherwise 'info'
	Debug = false
	// StackTrace is used by Solo. 'true' if stack trace must be printed in case of errors
	StackTrace = false

	CreateGameRequest = "{\"player_name\": \"Admiral Earl Howe\"}"
	JoinGameRequest = "{\"game_id\": \"001\", \"player_name\": \"Jeanbon Saint-André\"}"
	InitFieldRequest1 = "{\"game_id\": \"001\", \"field\": [[Empty, Empty, Bound, Bound, Bound, Empty, Bound, Bound, Bound, Empty, Empty, Empty], [Empty, Empty, Bound, Ship, Bound, Empty, Bound, Ship, Bound, Empty, Empty, Empty], [Empty, Empty, Bound, Ship, Bound, Bound, Bound, Ship, Bound, Empty, Empty, Empty], [Empty, Empty, Bound, Ship, Bound, Ship, Bound, Bound, Bound, Bound, Bound, Bound], [Empty, Empty, Bound, Ship, Bound, Ship, Bound, Bound, Ship, Bound, Ship, Bound], [Empty, Empty, Bound, Bound, Bound, Ship, Bound, Bound, Bound, Bound, Bound, Bound], [Empty, Bound, Bound, Bound, Bound, Bound, Bound, Bound, Bound, Bound, Ship, Bound], [Empty, Bound, Ship, Bound, Bound, Bound, Bound, Ship, Bound, Bound, Bound, Bound], [Empty, Bound, Ship, Bound, Ship, Bound, Bound, Ship, Bound, Bound, Bound, Empty], [Empty, Bound, Ship, Bound, Bound, Bound, Bound, Bound, Bound, Ship, Bound, Empty], [Empty, Bound, Bound, Bound, Empty, Empty, Empty, Empty, Bound, Ship, Bound, Empty], [Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Bound, Bound, Bound, Empty]]}"
	InitFieldRequest2 = "{\"game_id\": \"001\", \"field\": [[Empty, Bound, Bound, Bound, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty], [Empty, Bound, Ship, Bound, Bound, Bound, Bound, Bound, Empty, Bound, Bound, Bound], [Empty, Bound, Ship, Bound, Ship, Bound, Ship, Bound, Bound, Bound, Ship, Bound], [Bound, Bound, Bound, Bound, Ship, Bound, Ship, Bound, Ship, Bound, Bound, Bound], [Bound, Ship, Bound, Bound, Bound, Bound, Ship, Bound, Bound, Bound, Empty, Empty], [Bound, Bound, Bound, Bound, Bound, Bound, Bound, Bound, Bound, Bound, Bound, Empty], [Empty, Bound, Ship, Bound, Ship, Bound, Bound, Bound, Bound, Ship, Bound, Empty], [Empty, Bound, Ship, Bound, Bound, Bound, Bound, Ship, Bound, Ship, Bound, Empty], [Empty, Bound, Ship, Bound, Empty, Empty, Bound, Ship, Bound, Ship, Bound, Empty], [Empty, Bound, Bound, Bound, Empty, Empty, Bound, Bound, Bound, Ship, Bound, Empty], [Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Bound, Bound, Bound, Empty], [Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty]]}"
	MakeMoveRequest = ""
	QuitGameRequest = "{\"game_id\": \"001\"}"
	GetGameRequest = "{\"game_id\": \"001\"}"
)
