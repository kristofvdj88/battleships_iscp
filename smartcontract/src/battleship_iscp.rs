use battleship::*;
use iota_sc_utils::*;
use wasmlib::*;

use crate::consts::*;
use crate::helpers::*;
use crate::structures::*;

pub fn create_game(ctx: &ScFuncContext) {
    ctx.log("create_game");

    // Reads argument called "createGameRequestKey" and tries to deserialize it to the CreateGameRequest structure
    let req: Result<CreateGameRequest, serde_json::Error> =
        get_struct::<ScFuncContext, CreateGameRequest>(PARAM_CREATE_GAME_REQUEST, ctx);
    if !req.is_ok() {
        ctx.log("Request could not be deserialized");
        return;
    }
    ctx.log("Request successfully deserialized");
    let req_body = req.unwrap();

    // Gets the caller agent identifier from the context
    let caller_agent_id = ctx.caller().to_string();
    // Gets the provided stake that will be used in this game
    let stake = ctx.incoming().balance(&ScColor::IOTA) * 2;

    // Create a new game
    let game = Game::new(&caller_agent_id, &req_body.player_name, stake);
    ctx.log(
        &format!(
            "Game {} was created by player {}",
            game.id, req_body.player_name
        )[..],
    );

    // Save game to SC state as a JSON string
    let game_str = serde_json::to_string(&game).unwrap();
    ctx.state()
        .get_string(ACTIVE_GAME_STATE)
        .set_value(&game_str);
}

pub fn join_game(ctx: &ScFuncContext) {
    ctx.log("join_game");

    // Reads argument called "joinGameRequestKey" and tries to deserialize it to the JoinGameRequest structure
    let req: Result<JoinGameRequest, serde_json::Error> =
        get_struct::<ScFuncContext, JoinGameRequest>(PARAM_JOIN_GAME_REQUEST, ctx);
    if !req.is_ok() {
        ctx.log("Request could not be deserialized");
        return;
    }
    ctx.log("Request successfully deserialized");
    let req_body = req.unwrap();

    // Get the game from the SC state
    let game_str = ctx.state().get_string(ACTIVE_GAME_STATE).value();
    let mut game: Game = serde_json::from_str(&game_str).unwrap();

    // Check if the SC game state id matches the request game id, return if not
    if game.id != req_body.game_id {
        ctx.log(&format!("Game {} not found.", req_body.game_id)[..]);
        return;
    }

    // Check if the provided stake of the new player is correct
    let stake = ctx.incoming().balance(&ScColor::IOTA) * 2;
    if game.stake != stake {
        ctx.log(&format!("Player {} requested to join game {}, but supplied an incorrect stake. The stake of this game equals {}", &req_body.player_name, game.id, stake.to_string())[..]);
        return;
    }

    let caller_agent_id = ctx.caller().to_string();
    // Let the player join the game
    let game_ref = game.join(&caller_agent_id, &req_body.player_name);
    ctx.log(
        &format!(
            "Game {} was joined by player {}",
            game_ref.id, req_body.player_name
        )[..],
    );

    // Save game to SC state as a JSON string
    let game_str = serde_json::to_string(game_ref).unwrap();
    ctx.state()
        .get_string(ACTIVE_GAME_STATE)
        .set_value(&game_str);
}

pub fn init_field(ctx: &ScFuncContext) {
    ctx.log("init_field");

    // Reads argument called "initFieldRequestKey" and tries to deserialize it to the InitFieldRequest structure
    let req: Result<InitFieldRequest, serde_json::Error> =
        get_struct::<ScFuncContext, InitFieldRequest>(PARAM_INIT_FIELD_REQUEST, ctx);
    if !req.is_ok() {
        ctx.log("Request could not be deserialized");
        return;
    }
    ctx.log("Request successfully deserialized");
    let req_body = req.unwrap();

    // Get the game from the SC state
    let game_str = ctx.state().get_string(ACTIVE_GAME_STATE).value();
    let mut game: Game = serde_json::from_str(&game_str).unwrap();

    if !validate_request(ctx, &game, req_body.game_id) {
        ctx.log("Request could not be validated. Terminating request.");
        return;
    }

    // Initialize the game field of the request player and start the game if it's ready
    let caller_agent_id = ctx.caller().to_string();
    game.init(caller_agent_id, req_body.field);
    if game.is_ready() {
        game.start();
    }

    // Save game to SC state as a JSON string
    let game_str = serde_json::to_string(&game).unwrap();
    ctx.state()
        .get_string(ACTIVE_GAME_STATE)
        .set_value(&game_str);
}

pub fn make_move(ctx: &ScFuncContext) {
    ctx.log("make_move");

    // Reads argument called "moveRequestKey" and tries to deserialize it to the MoveRequest structure
    let req: Result<MoveRequest, serde_json::Error> =
        get_struct::<ScFuncContext, MoveRequest>(PARAM_MOVE_REQUEST, ctx);
    if !req.is_ok() {
        ctx.log("Request could not be deserialized");
        return;
    }
    ctx.log("Request successfully deserialized");
    let req_body = req.unwrap();

    // Get the game from the SC state
    let game_str = ctx.state().get_string(ACTIVE_GAME_STATE).value();
    let mut game: Game = serde_json::from_str(&game_str).unwrap();

    if !validate_request(ctx, &game, req_body.game_id) {
        ctx.log("Request could not be validated. Terminating request.");
        return;
    }

    // Make the request player's move
    let caller_agent_id = ctx.caller().to_string();
    game.player_move(caller_agent_id, req_body.point);

    // Check if the game has been settled and proclaim a winner if so
    if game
        .mediator
        .player_a
        .as_ref()
        .unwrap()
        .own_field
        .sunked_ships
        == utils::ALL_SHIPS
    {
        let winner_id = game.mediator.player_a.as_ref().unwrap().id.to_string();
        game.won(&winner_id);
    }
    if game
        .mediator
        .player_b
        .as_ref()
        .unwrap()
        .own_field
        .sunked_ships
        == utils::ALL_SHIPS
    {
        let winner_id = game.mediator.player_b.as_ref().unwrap().id.to_string();
        game.won(&winner_id);
    }

    // Send all staked tokens of this game to the winner
    ctx.log(&format!("Player {} has won game {}.", game.winner_id, game.id)[..]);
    let winner_id_vec = ctx.utility().base58_decode(&game.winner_id[..]);
    let winner_id_arr = vector_as_u8_array(winner_id_vec);
    let winner = ScAgentId::from_bytes(&winner_id_arr);
    let bal = ctx.balances().balance(&ScColor::IOTA);
    if bal >= game.stake {
        ctx.transfer_to_address(
            &winner.address(),
            ScTransfers::new(&ScColor::IOTA, game.stake),
        )
    }

    // Remove the game from the active game state
    ctx.state().get_string(ACTIVE_GAME_STATE).set_value("");

    // Save the game to the settled games state
    let mut settled_game_str = ctx.state().get_string(SETTLED_GAMES_STATE).value();
    let mut settled_games: SettledGames = serde_json::from_str(&settled_game_str).unwrap();
    settled_games.games.push(game);
    settled_game_str = serde_json::to_string(&settled_games).unwrap();
    ctx.state()
        .get_string(SETTLED_GAMES_STATE)
        .set_value(&settled_game_str);
}

pub fn quit_game(ctx: &ScFuncContext) {
    ctx.log("quit_game");

    // Reads argument called "quitGameRequestKey" and tries to deserialize it to the QuitGameRequest structure
    let req: Result<QuitGameRequest, serde_json::Error> =
        get_struct::<ScFuncContext, QuitGameRequest>(PARAM_QUIT_GAME_REQUEST, ctx);
    if !req.is_ok() {
        ctx.log("Request could not be deserialized");
        return;
    }
    ctx.log("Request successfully deserialized");
    let req_body = req.unwrap();

    // Get the game from the SC state
    let game_str = ctx.state().get_string(ACTIVE_GAME_STATE).value();
    let mut game: Game = serde_json::from_str(&game_str).unwrap();

    // Check if the SC game state id matches the request game id, return if not
    if game.id == req_body.game_id {
        ctx.log(&format!("Game {} not found.", req_body.game_id)[..]);
    }

    // Check if the SC game state player id matches the request caller_agent_id, return if not
    if !game.mediator.player_a.is_some() && !game.mediator.player_b.is_some() {
        ctx.log(&format!("Game {} has no players.", game.id)[..]);
        return;
    }
    let caller_agent_id = ctx.caller().to_string();
    if game.mediator.player_a.as_ref().unwrap().id != caller_agent_id
        && game.mediator.player_b.as_ref().unwrap().id != caller_agent_id
    {
        ctx.log(&format!("Player {} not found in game {}.", caller_agent_id, game.id)[..]);
        return;
    }

    // Settle the game
    ctx.log(
        &format!(
            "Player {} wants to forfait game {}.",
            caller_agent_id, game.id
        )[..],
    );
    if game.mediator.player_a.as_ref().unwrap().id == caller_agent_id {
        let winner_id = game.mediator.player_b.as_ref().unwrap().id.to_string();
        game.won(&winner_id);
    }
    if game.mediator.player_b.as_ref().unwrap().id == caller_agent_id {
        let winner_id = game.mediator.player_a.as_ref().unwrap().id.to_string();
        game.won(&winner_id);
    }

    // Remove the game from the active game state
    ctx.state().get_string(ACTIVE_GAME_STATE).set_value("");

    // Save the game to the settled games state
    let mut settled_game_str = ctx.state().get_string(SETTLED_GAMES_STATE).value();
    let mut settled_games: SettledGames = serde_json::from_str(&settled_game_str).unwrap();
    settled_games.games.push(game);
    settled_game_str = serde_json::to_string(&settled_games).unwrap();
    ctx.state()
        .get_string(SETTLED_GAMES_STATE)
        .set_value(&settled_game_str);
}

// Public view
pub fn get_game(ctx: &ScViewContext) {
    ctx.log("get_game");

    // Reads argument called "quitGameRequestKey" and tries to deserialize it to the QuitGameRequest structure
    let req: Result<GetGameRequest, serde_json::Error> =
        get_struct::<ScViewContext, GetGameRequest>(PARAM_QUIT_GAME_REQUEST, ctx);
    if !req.is_ok() {
        ctx.log("Request could not be deserialized");
        return;
    }
    ctx.log("Request successfully deserialized");
    let req_body = req.unwrap();

    // Get the game from the SC state
    let game_str = ctx.state().get_string(ACTIVE_GAME_STATE).value();
    let game: Game = serde_json::from_str(&game_str).unwrap();

    // validate if request game id matches state game id
    if req_body.game_id != game.id {
        ctx.log(&format!("Game {} not found.", req_body.game_id)[..]);
        ctx.results()
            .get_string(PARAM_GAME_STATE_RESPONSE)
            .set_value("");
        return;
    }

    // Attach the game state to the responce with the key "gameStateResponseKey"
    ctx.results()
        .get_string(PARAM_GAME_STATE_RESPONSE)
        .set_value(&game_str);
}

// Private functions
fn validate_request(ctx: &ScFuncContext, game: &Game, req_game_id: String) -> bool {
    // Check if the SC game state id matches the request game id, return if not
    if game.id != req_game_id {
        ctx.log(&format!("Game {} not found.", req_game_id)[..]);
        return false;
    }

    // Check if the SC game contains players
    if !game.mediator.player_a.is_some() && !game.mediator.player_b.is_some() {
        ctx.log(&format!("Game {} has no players.", game.id)[..]);
        return false;
    }

    // Check if game is already settled
    if game.is_settled() {
        ctx.log(
            &format!(
                "Game {} has already been settled. It was won by {}.",
                game.id, game.winner_id
            )[..],
        );
        return false;
    }

    // Check if a SC game state player id matches the request caller_agent_id, return if not
    let caller_agent_id = ctx.caller().to_string();
    if (game.mediator.player_a.is_some()
        && game.mediator.player_a.as_ref().unwrap().id != caller_agent_id)
        && (game.mediator.player_b.is_some()
            && game.mediator.player_b.as_ref().unwrap().id != caller_agent_id)
    {
        ctx.log(&format!("Player {} not found in game {}.", caller_agent_id, game.id)[..]);
        return false;
    }

    true
}

fn settle_game(ctx: ScFuncContext, game: Game) {
    // Send all staked tokens of this game to the winner
    ctx.log(&format!("Player {} has won game {}.", game.winner_id, game.id)[..]);
    let winner_id_vec = ctx.utility().base58_decode(&game.winner_id[..]);
    let winner_id_arr = vector_as_u8_array(winner_id_vec);
    let winner = ScAgentId::from_bytes(&winner_id_arr);
    let bal = ctx.balances().balance(&ScColor::IOTA);
    if bal >= game.stake {
        ctx.transfer_to_address(
            &winner.address(),
            ScTransfers::new(&ScColor::IOTA, game.stake),
        )
    }

    //TODO: add functionality to be able to host several active games at once and remove usage of SETTLED_GAMES_STATE
    // Remove the game from the active game state
    ctx.state().get_string(ACTIVE_GAME_STATE).set_value("");

    // Save the game to the settled games state
    let mut settled_game_str = ctx.state().get_string(SETTLED_GAMES_STATE).value();
    let mut settled_games: SettledGames = serde_json::from_str(&settled_game_str).unwrap();
    settled_games.games.push(game);
    settled_game_str = serde_json::to_string(&settled_games).unwrap();
    ctx.state()
        .get_string(SETTLED_GAMES_STATE)
        .set_value(&settled_game_str);
}
