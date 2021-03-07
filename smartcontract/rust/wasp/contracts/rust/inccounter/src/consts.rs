// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

#![allow(dead_code)]

use wasmlib::*;

pub const SC_NAME: &str = "inccounter";
pub const SC_HNAME: ScHname = ScHname(0xaf2438e9);

pub const PARAM_COUNTER: &str = "counter";
pub const PARAM_NUM_REPEATS: &str = "numRepeats";

pub const VAR_COUNTER: &str = "counter";
pub const VAR_INT1: &str = "int1";
pub const VAR_INT_ARRAY1: &str = "intArray1";
pub const VAR_NUM_REPEATS: &str = "numRepeats";
pub const VAR_STRING1: &str = "string1";
pub const VAR_STRING_ARRAY1: &str = "stringArray1";

pub const FUNC_CALL_INCREMENT: &str = "callIncrement";
pub const FUNC_CALL_INCREMENT_RECURSE5X: &str = "callIncrementRecurse5x";
pub const FUNC_INCREMENT: &str = "increment";
pub const FUNC_INIT: &str = "init";
pub const FUNC_LOCAL_STATE_INTERNAL_CALL: &str = "localStateInternalCall";
pub const FUNC_LOCAL_STATE_POST: &str = "localStatePost";
pub const FUNC_LOCAL_STATE_SANDBOX_CALL: &str = "localStateSandboxCall";
pub const FUNC_POST_INCREMENT: &str = "postIncrement";
pub const FUNC_REPEAT_MANY: &str = "repeatMany";
pub const FUNC_WHEN_MUST_INCREMENT: &str = "whenMustIncrement";
pub const VIEW_GET_COUNTER: &str = "getCounter";

pub const HFUNC_CALL_INCREMENT: ScHname = ScHname(0xeb5dcacd);
pub const HFUNC_CALL_INCREMENT_RECURSE5X: ScHname = ScHname(0x8749fbff);
pub const HFUNC_INCREMENT: ScHname = ScHname(0xd351bd12);
pub const HFUNC_INIT: ScHname = ScHname(0x1f44d644);
pub const HFUNC_LOCAL_STATE_INTERNAL_CALL: ScHname = ScHname(0xecfc5d33);
pub const HFUNC_LOCAL_STATE_POST: ScHname = ScHname(0x3fd54d13);
pub const HFUNC_LOCAL_STATE_SANDBOX_CALL: ScHname = ScHname(0x7bd22c53);
pub const HFUNC_POST_INCREMENT: ScHname = ScHname(0x81c772f5);
pub const HFUNC_REPEAT_MANY: ScHname = ScHname(0x4ff450d3);
pub const HFUNC_WHEN_MUST_INCREMENT: ScHname = ScHname(0xb4c3e7a6);
pub const HVIEW_GET_COUNTER: ScHname = ScHname(0xb423e607);
