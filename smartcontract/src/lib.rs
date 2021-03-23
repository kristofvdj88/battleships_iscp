use battleship_iscp::*;
use iota_sc_utils::wasmlib::ScExports;

mod battleship_iscp;
mod consts;
mod helpers;
mod structures;

#[no_mangle]
fn on_load() {
    let exports = ScExports::new();

    // SC functions
    exports.add_func("create_game", create_game);
    exports.add_func("join_game", join_game);
    exports.add_func("init_field", init_field);
    exports.add_func("make_move", make_move);
    exports.add_func("quit_game", quit_game);

    // SC Views
    //exports.add_view("getGames", get_games);
    exports.add_view("getGame", get_game);
}
