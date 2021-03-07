// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

#![allow(dead_code)]

use wasmlib::*;

pub const SC_NAME: &str = "erc20";
pub const SC_DESCRIPTION: &str = "ERC-20 PoC for IOTA Smart Contracts";
pub const SC_HNAME: ScHname = ScHname(0x200e3733);

pub const PARAM_ACCOUNT: &str = "ac";
pub const PARAM_AMOUNT: &str = "am";
pub const PARAM_CREATOR: &str = "c";
pub const PARAM_DELEGATION: &str = "d";
pub const PARAM_RECIPIENT: &str = "r";
pub const PARAM_SUPPLY: &str = "s";

pub const VAR_BALANCES: &str = "b";
pub const VAR_SUPPLY: &str = "s";

pub const FUNC_APPROVE: &str = "approve";
pub const FUNC_INIT: &str = "init";
pub const FUNC_TRANSFER: &str = "transfer";
pub const FUNC_TRANSFER_FROM: &str = "transferFrom";
pub const VIEW_ALLOWANCE: &str = "allowance";
pub const VIEW_BALANCE_OF: &str = "balanceOf";
pub const VIEW_TOTAL_SUPPLY: &str = "totalSupply";

pub const HFUNC_APPROVE: ScHname = ScHname(0xa0661268);
pub const HFUNC_INIT: ScHname = ScHname(0x1f44d644);
pub const HFUNC_TRANSFER: ScHname = ScHname(0xa15da184);
pub const HFUNC_TRANSFER_FROM: ScHname = ScHname(0xd5e0a602);
pub const HVIEW_ALLOWANCE: ScHname = ScHname(0x5e16006a);
pub const HVIEW_BALANCE_OF: ScHname = ScHname(0x67ef8df4);
pub const HVIEW_TOTAL_SUPPLY: ScHname = ScHname(0x9505e6ca);
