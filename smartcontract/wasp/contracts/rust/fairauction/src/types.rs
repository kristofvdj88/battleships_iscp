// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

use wasmlib::*;

//@formatter:off
pub struct Auction {
    pub color:          ScColor,   // color of tokens for sale
    pub creator:        ScAgentId, // issuer of start_auction transaction
    pub deposit:        i64,       // deposit by auction owner to cover the SC fees
    pub description:    String,    // auction description
    pub duration:       i64,       // auction duration in minutes
    pub highest_bid:    i64,       // the current highest bid amount
    pub highest_bidder: ScAgentId, // the current highest bidder
    pub minimum_bid:    i64,       // minimum bid amount
    pub num_tokens:     i64,       // number of tokens for sale
    pub owner_margin:   i64,       // auction owner's margin in promilles
    pub when_started:   i64,       // timestamp when auction started
}
//@formatter:on

impl Auction {
    pub fn from_bytes(bytes: &[u8]) -> Auction {
        let mut decode = BytesDecoder::new(bytes);
        Auction {
            color: decode.color(),
            creator: decode.agent_id(),
            deposit: decode.int(),
            description: decode.string(),
            duration: decode.int(),
            highest_bid: decode.int(),
            highest_bidder: decode.agent_id(),
            minimum_bid: decode.int(),
            num_tokens: decode.int(),
            owner_margin: decode.int(),
            when_started: decode.int(),
        }
    }

    pub fn to_bytes(&self) -> Vec<u8> {
        let mut encode = BytesEncoder::new();
        encode.color(&self.color);
        encode.agent_id(&self.creator);
        encode.int(self.deposit);
        encode.string(&self.description);
        encode.int(self.duration);
        encode.int(self.highest_bid);
        encode.agent_id(&self.highest_bidder);
        encode.int(self.minimum_bid);
        encode.int(self.num_tokens);
        encode.int(self.owner_margin);
        encode.int(self.when_started);
        return encode.data();
    }
}

//@formatter:off
pub struct Bid {
    pub amount:    i64, // cumulative amount of bids from same bidder
    pub index:     i64, // index of bidder in bidder list
    pub timestamp: i64, // timestamp of most recent bid
}
//@formatter:on

impl Bid {
    pub fn from_bytes(bytes: &[u8]) -> Bid {
        let mut decode = BytesDecoder::new(bytes);
        Bid {
            amount: decode.int(),
            index: decode.int(),
            timestamp: decode.int(),
        }
    }

    pub fn to_bytes(&self) -> Vec<u8> {
        let mut encode = BytesEncoder::new();
        encode.int(self.amount);
        encode.int(self.index);
        encode.int(self.timestamp);
        return encode.data();
    }
}
