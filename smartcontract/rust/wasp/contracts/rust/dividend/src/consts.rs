// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

#![allow(dead_code)]

use wasmlib::*;

pub const SC_NAME: &str = "dividend";
pub const SC_HNAME: ScHname = ScHname(0xcce2e239);

pub const PARAM_ADDRESS: &str = "address";
pub const PARAM_FACTOR: &str = "factor";

pub const VAR_MEMBERS: &str = "members";
pub const VAR_TOTAL_FACTOR: &str = "totalFactor";

pub const FUNC_DIVIDE: &str = "divide";
pub const FUNC_MEMBER: &str = "member";

pub const HFUNC_DIVIDE: ScHname = ScHname(0xc7878107);
pub const HFUNC_MEMBER: ScHname = ScHname(0xc07da2cb);
