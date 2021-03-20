use wasmlib::*;
use battleship_iscp::*;

mod consts;
mod battleship_iscp;
mod structures;
mod helpers;



#[no_mangle]
fn on_load() {
    let exports = ScExports::new();

    // SC functions
    exports.add_func("create_game", create_game);
    exports.add_func("join_game", join_game);
    exports.add_func("init_field", init_field);
    exports.add_func("make_move", make_move);
    exports.add_func("quit_game", quit_game);
    exports.add_func("contract_creator_only_function", contract_creator_only_function);
    exports.add_func("chain_owner_only_function", chain_owner_only_function);

    // SC Views
    //exports.add_view("getGames", get_games);
    exports.add_view("getGame", get_game);
}