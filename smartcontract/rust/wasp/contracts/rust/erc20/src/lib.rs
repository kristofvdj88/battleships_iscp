// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

use consts::*;
use erc20::*;
use wasmlib::*;

mod consts;
mod erc20;

#[no_mangle]
fn on_load() {
    let exports = ScExports::new();
    exports.add_func(FUNC_APPROVE, func_approve);
    exports.add_func(FUNC_INIT, func_init);
    exports.add_func(FUNC_TRANSFER, func_transfer);
    exports.add_func(FUNC_TRANSFER_FROM, func_transfer_from);
    exports.add_view(VIEW_ALLOWANCE, view_allowance);
    exports.add_view(VIEW_BALANCE_OF, view_balance_of);
    exports.add_view(VIEW_TOTAL_SUPPLY, view_total_supply);
}
