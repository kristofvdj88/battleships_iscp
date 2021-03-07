// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

use wasmlib::*;

//@formatter:off
pub struct Donation {
    pub amount:    i64,       // amount donated
    pub donator:   ScAgentId, // who donated
    pub error:     String,    // error to be reported to donator if anything goes wrong
    pub feedback:  String,    // the feedback for the person donated to
    pub timestamp: i64,       // when the donation took place
}
//@formatter:on

impl Donation {
    pub fn from_bytes(bytes: &[u8]) -> Donation {
        let mut decode = BytesDecoder::new(bytes);
        Donation {
            amount: decode.int(),
            donator: decode.agent_id(),
            error: decode.string(),
            feedback: decode.string(),
            timestamp: decode.int(),
        }
    }

    pub fn to_bytes(&self) -> Vec<u8> {
        let mut encode = BytesEncoder::new();
        encode.int(self.amount);
        encode.agent_id(&self.donator);
        encode.string(&self.error);
        encode.string(&self.feedback);
        encode.int(self.timestamp);
        return encode.data();
    }
}
