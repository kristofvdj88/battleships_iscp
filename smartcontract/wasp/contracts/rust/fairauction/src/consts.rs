// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

#![allow(dead_code)]

use wasmlib::*;

pub const SC_NAME: &str = "fairauction";
pub const SC_HNAME: ScHname = ScHname(0x1b5c43b1);

pub const PARAM_COLOR: &str = "color";
pub const PARAM_DESCRIPTION: &str = "description";
pub const PARAM_DURATION: &str = "duration";
pub const PARAM_MINIMUM_BID: &str = "minimumBid";
pub const PARAM_OWNER_MARGIN: &str = "ownerMargin";

pub const VAR_AUCTIONS: &str = "auctions";
pub const VAR_BIDDER_LIST: &str = "bidderList";
pub const VAR_BIDDERS: &str = "bidders";
pub const VAR_COLOR: &str = "color";
pub const VAR_CREATOR: &str = "creator";
pub const VAR_DEPOSIT: &str = "deposit";
pub const VAR_DESCRIPTION: &str = "description";
pub const VAR_DURATION: &str = "duration";
pub const VAR_HIGHEST_BID: &str = "highestBid";
pub const VAR_HIGHEST_BIDDER: &str = "highestBidder";
pub const VAR_INFO: &str = "info";
pub const VAR_MINIMUM_BID: &str = "minimumBid";
pub const VAR_NUM_TOKENS: &str = "numTokens";
pub const VAR_OWNER_MARGIN: &str = "ownerMargin";
pub const VAR_WHEN_STARTED: &str = "whenStarted";

pub const FUNC_FINALIZE_AUCTION: &str = "finalizeAuction";
pub const FUNC_PLACE_BID: &str = "placeBid";
pub const FUNC_SET_OWNER_MARGIN: &str = "setOwnerMargin";
pub const FUNC_START_AUCTION: &str = "startAuction";
pub const VIEW_GET_INFO: &str = "getInfo";

pub const HFUNC_FINALIZE_AUCTION: ScHname = ScHname(0x8d534ddc);
pub const HFUNC_PLACE_BID: ScHname = ScHname(0x9bd72fa9);
pub const HFUNC_SET_OWNER_MARGIN: ScHname = ScHname(0x1774461a);
pub const HFUNC_START_AUCTION: ScHname = ScHname(0xd5b7bacb);
pub const HVIEW_GET_INFO: ScHname = ScHname(0xcfedba5f);
