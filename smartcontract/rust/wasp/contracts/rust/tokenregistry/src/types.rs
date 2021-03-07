// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

use wasmlib::*;

//@formatter:off
pub struct Token {
    pub created:      i64,       // creation timestamp
    pub description:  String,    // description what minted token represents
    pub minted_by:    ScAgentId, // original minter
    pub owner:        ScAgentId, // current owner
    pub supply:       i64,       // amount of tokens originally minted
    pub updated:      i64,       // last update timestamp
    pub user_defined: String,    // any user defined text
}
//@formatter:on

impl Token {
    pub fn from_bytes(bytes: &[u8]) -> Token {
        let mut decode = BytesDecoder::new(bytes);
        Token {
            created: decode.int(),
            description: decode.string(),
            minted_by: decode.agent_id(),
            owner: decode.agent_id(),
            supply: decode.int(),
            updated: decode.int(),
            user_defined: decode.string(),
        }
    }

    pub fn to_bytes(&self) -> Vec<u8> {
        let mut encode = BytesEncoder::new();
        encode.int(self.created);
        encode.string(&self.description);
        encode.agent_id(&self.minted_by);
        encode.agent_id(&self.owner);
        encode.int(self.supply);
        encode.int(self.updated);
        encode.string(&self.user_defined);
        return encode.data();
    }
}
