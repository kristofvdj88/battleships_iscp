// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

use wasmlib::*;

//@formatter:off
pub struct Member {
    pub address: ScAddress, // address of dividend recipient
    pub factor:  i64,       // relative division factor
}
//@formatter:on

impl Member {
    pub fn from_bytes(bytes: &[u8]) -> Member {
        let mut decode = BytesDecoder::new(bytes);
        Member {
            address: decode.address(),
            factor: decode.int(),
        }
    }

    pub fn to_bytes(&self) -> Vec<u8> {
        let mut encode = BytesEncoder::new();
        encode.address(&self.address);
        encode.int(self.factor);
        return encode.data();
    }
}
