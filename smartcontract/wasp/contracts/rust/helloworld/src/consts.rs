// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

#![allow(dead_code)]

use wasmlib::*;

pub const SC_NAME: &str = "helloworld";
pub const SC_DESCRIPTION: &str = "The ubiquitous hello world demo";
pub const SC_HNAME: ScHname = ScHname(0x0683223c);

pub const VAR_HELLO_WORLD: &str = "helloWorld";

pub const FUNC_HELLO_WORLD: &str = "helloWorld";
pub const VIEW_GET_HELLO_WORLD: &str = "getHelloWorld";

pub const HFUNC_HELLO_WORLD: ScHname = ScHname(0x9d042e65);
pub const HVIEW_GET_HELLO_WORLD: ScHname = ScHname(0x210439ce);
