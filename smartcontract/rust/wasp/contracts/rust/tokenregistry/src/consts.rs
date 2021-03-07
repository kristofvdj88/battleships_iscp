// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

#![allow(dead_code)]

use wasmlib::*;

pub const SC_NAME: &str = "tokenregistry";
pub const SC_HNAME: ScHname = ScHname(0xe1ba0c78);

pub const PARAM_COLOR: &str = "color";
pub const PARAM_DESCRIPTION: &str = "description";
pub const PARAM_USER_DEFINED: &str = "userDefined";

pub const VAR_COLOR_LIST: &str = "colorList";
pub const VAR_REGISTRY: &str = "registry";

pub const FUNC_MINT_SUPPLY: &str = "mintSupply";
pub const FUNC_TRANSFER_OWNERSHIP: &str = "transferOwnership";
pub const FUNC_UPDATE_METADATA: &str = "updateMetadata";
pub const VIEW_GET_INFO: &str = "getInfo";

pub const HFUNC_MINT_SUPPLY: ScHname = ScHname(0x564349a7);
pub const HFUNC_TRANSFER_OWNERSHIP: ScHname = ScHname(0xbb9eb5af);
pub const HFUNC_UPDATE_METADATA: ScHname = ScHname(0xa26b23b6);
pub const HVIEW_GET_INFO: ScHname = ScHname(0xcfedba5f);
