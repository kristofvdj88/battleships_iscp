// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// types encapsulating immutable host objects

use std::convert::TryInto;

use crate::context::*;
use crate::hashtypes::*;
use crate::host::*;
use crate::keys::*;

// proxy object for immutable ScAddress in host map
pub struct ScImmutableAddress {
    obj_id: i32,
    key_id: Key32,
}

impl ScImmutableAddress {
    // check if object exists in host map
    pub fn exists(&self) -> bool {
        exists(self.obj_id, self.key_id, TYPE_ADDRESS)
    }

    // human-readable string representation
    pub fn to_string(&self) -> String {
        self.value().to_string()
    }

    // get value from host map
    pub fn value(&self) -> ScAddress {
        ScAddress::from_bytes(&get_bytes(self.obj_id, self.key_id, TYPE_ADDRESS))
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

// immutable array of ScAddress
pub struct ScImmutableAddressArray {
    pub(crate) obj_id: i32
}

impl ScImmutableAddressArray {
    // index 0..length(), exclusive
    pub fn get_address(&self, index: i32) -> ScImmutableAddress {
        ScImmutableAddress { obj_id: self.obj_id, key_id: Key32(index) }
    }

    // number of items in array
    pub fn length(&self) -> i32 {
        get_length(self.obj_id)
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

// proxy object for immutable ScAgentId in host map
pub struct ScImmutableAgentId {
    obj_id: i32,
    key_id: Key32,
}

impl ScImmutableAgentId {
    // check if object exists in host map
    pub fn exists(&self) -> bool {
        exists(self.obj_id, self.key_id, TYPE_AGENT_ID)
    }

    // human-readable string representation
    pub fn to_string(&self) -> String {
        self.value().to_string()
    }

    // get value from host map
    pub fn value(&self) -> ScAgentId {
        ScAgentId::from_bytes(&get_bytes(self.obj_id, self.key_id, TYPE_AGENT_ID))
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

// immutable array of ScAgentId
pub struct ScImmutableAgentIdArray {
    pub(crate) obj_id: i32
}

impl ScImmutableAgentIdArray {
    // index 0..length(), exclusive
    pub fn get_agent_id(&self, index: i32) -> ScImmutableAgentId {
        ScImmutableAgentId { obj_id: self.obj_id, key_id: Key32(index) }
    }

    // number of items in array
    pub fn length(&self) -> i32 {
        get_length(self.obj_id)
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

// proxy object for immutable bytes array in host map
pub struct ScImmutableBytes {
    obj_id: i32,
    key_id: Key32,
}

impl ScImmutableBytes {
    // check if object exists in host map
    pub fn exists(&self) -> bool {
        exists(self.obj_id, self.key_id, TYPE_BYTES)
    }

    // human-readable string representation
    pub fn to_string(&self) -> String {
        base58_encode(&self.value())
    }

    // get value from host map
    pub fn value(&self) -> Vec<u8> {
        get_bytes(self.obj_id, self.key_id, TYPE_BYTES)
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

// immutable array of byte array
pub struct ScImmutableBytesArray {
    pub(crate) obj_id: i32
}

impl ScImmutableBytesArray {
    // index 0..length(), exclusive
    pub fn get_bytes(&self, index: i32) -> ScImmutableBytes {
        ScImmutableBytes { obj_id: self.obj_id, key_id: Key32(index) }
    }

    // number of items in array
    pub fn length(&self) -> i32 {
        get_length(self.obj_id)
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

// proxy object for immutable ScChainId in host map
pub struct ScImmutableChainId {
    obj_id: i32,
    key_id: Key32,
}

impl ScImmutableChainId {
    // check if object exists in host map
    pub fn exists(&self) -> bool {
        exists(self.obj_id, self.key_id, TYPE_CHAIN_ID)
    }

    // human-readable string representation
    pub fn to_string(&self) -> String {
        self.value().to_string()
    }

    // get value from host map
    pub fn value(&self) -> ScChainId {
        ScChainId::from_bytes(&get_bytes(self.obj_id, self.key_id, TYPE_CHAIN_ID))
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

// proxy object for immutable ScColor in host map
pub struct ScImmutableColor {
    obj_id: i32,
    key_id: Key32,
}

impl ScImmutableColor {
    // check if object exists in host map
    pub fn exists(&self) -> bool {
        exists(self.obj_id, self.key_id, TYPE_COLOR)
    }

    // human-readable string representation
    pub fn to_string(&self) -> String {
        self.value().to_string()
    }

    // get value from host map
    pub fn value(&self) -> ScColor {
        ScColor::from_bytes(&get_bytes(self.obj_id, self.key_id, TYPE_COLOR))
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

// immutable array of ScColor
pub struct ScImmutableColorArray {
    pub(crate) obj_id: i32
}

impl ScImmutableColorArray {
    // index 0..length(), exclusive
    pub fn get_color(&self, index: i32) -> ScImmutableColor {
        ScImmutableColor { obj_id: self.obj_id, key_id: Key32(index) }
    }

    // number of items in array
    pub fn length(&self) -> i32 {
        get_length(self.obj_id)
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

// proxy object for immutable ScContractId in host map
pub struct ScImmutableContractId {
    obj_id: i32,
    key_id: Key32,
}

impl ScImmutableContractId {
    // check if object exists in host map
    pub fn exists(&self) -> bool {
        exists(self.obj_id, self.key_id, TYPE_CONTRACT_ID)
    }

    // human-readable string representation
    pub fn to_string(&self) -> String {
        self.value().to_string()
    }

    // get value from host map
    pub fn value(&self) -> ScContractId {
        ScContractId::from_bytes(&get_bytes(self.obj_id, self.key_id, TYPE_CONTRACT_ID))
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

// proxy object for immutable ScHash in host map
pub struct ScImmutableHash {
    obj_id: i32,
    key_id: Key32,
}

impl ScImmutableHash {
    // check if object exists in host map
    pub fn exists(&self) -> bool {
        exists(self.obj_id, self.key_id, TYPE_HASH)
    }

    // human-readable string representation
    pub fn to_string(&self) -> String {
        self.value().to_string()
    }

    // get value from host map
    pub fn value(&self) -> ScHash {
        ScHash::from_bytes(&get_bytes(self.obj_id, self.key_id, TYPE_HASH))
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

// immutable array of ScHash
pub struct ScImmutableHashArray {
    pub(crate) obj_id: i32
}

impl ScImmutableHashArray {
    // index 0..length(), exclusive
    pub fn get_hash(&self, index: i32) -> ScImmutableHash {
        ScImmutableHash { obj_id: self.obj_id, key_id: Key32(index) }
    }

    // number of items in array
    pub fn length(&self) -> i32 {
        get_length(self.obj_id)
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

// proxy object for immutable ScHname in host map
pub struct ScImmutableHname {
    obj_id: i32,
    key_id: Key32,
}

impl ScImmutableHname {
    // check if object exists in host map
    pub fn exists(&self) -> bool {
        exists(self.obj_id, self.key_id, TYPE_HNAME)
    }

    // human-readable string representation
    pub fn to_string(&self) -> String {
        self.value().to_string()
    }

    // get value from host map
    pub fn value(&self) -> ScHname {
        ScHname::from_bytes(&get_bytes(self.obj_id, self.key_id, TYPE_HNAME))
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

// proxy object for immutable int64 in host map
pub struct ScImmutableInt {
    obj_id: i32,
    key_id: Key32,
}

impl ScImmutableInt {
    // check if object exists in host map
    pub fn exists(&self) -> bool {
        exists(self.obj_id, self.key_id, TYPE_INT)
    }

    // human-readable string representation
    pub fn to_string(&self) -> String {
        self.value().to_string()
    }

    // get value from host map
    pub fn value(&self) -> i64 {
        let bytes = get_bytes(self.obj_id, self.key_id, TYPE_INT);
        i64::from_le_bytes(bytes.try_into().expect("invalid i64 length"))
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

// immutable array of int64
pub struct ScImmutableIntArray {
    pub(crate) obj_id: i32
}

impl ScImmutableIntArray {
    // index 0..length(), exclusive
    pub fn get_int(&self, index: i32) -> ScImmutableInt {
        ScImmutableInt { obj_id: self.obj_id, key_id: Key32(index) }
    }

    // number of items in array
    pub fn length(&self) -> i32 {
        get_length(self.obj_id)
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

pub struct ScImmutableMap {
    pub(crate) obj_id: i32
}

impl ScImmutableMap {
    // get proxy for immutable ScAddress field specified by key
    pub fn get_address<T: MapKey + ?Sized>(&self, key: &T) -> ScImmutableAddress {
        ScImmutableAddress { obj_id: self.obj_id, key_id: key.get_id() }
    }

    // get proxy for ScImmutableAddressArray specified by key
    pub fn get_address_array<T: MapKey + ?Sized>(&self, key: &T) -> ScImmutableAddressArray {
        let arr_id = get_object_id(self.obj_id, key.get_id(), TYPE_ADDRESS | TYPE_ARRAY);
        ScImmutableAddressArray { obj_id: arr_id }
    }

    // get proxy for immutable ScAgentId field specified by key
    pub fn get_agent_id<T: MapKey + ?Sized>(&self, key: &T) -> ScImmutableAgentId {
        ScImmutableAgentId { obj_id: self.obj_id, key_id: key.get_id() }
    }

    // get proxy for ScImmutableAgentIdArray specified by key
    pub fn get_agent_id_array<T: MapKey + ?Sized>(&self, key: &T) -> ScImmutableAgentIdArray {
        let arr_id = get_object_id(self.obj_id, key.get_id(), TYPE_AGENT_ID | TYPE_ARRAY);
        ScImmutableAgentIdArray { obj_id: arr_id }
    }

    // get proxy for immutable bytes array field specified by key
    pub fn get_bytes<T: MapKey + ?Sized>(&self, key: &T) -> ScImmutableBytes {
        ScImmutableBytes { obj_id: self.obj_id, key_id: key.get_id() }
    }

    // get proxy for ScImmutableBytesArray specified by key
    pub fn get_bytes_array<T: MapKey + ?Sized>(&self, key: &T) -> ScImmutableBytesArray {
        let arr_id = get_object_id(self.obj_id, key.get_id(), TYPE_BYTES | TYPE_ARRAY);
        ScImmutableBytesArray { obj_id: arr_id }
    }

    // get proxy for immutable ScChainId field specified by key
    pub fn get_chain_id<T: MapKey + ?Sized>(&self, key: &T) -> ScImmutableChainId {
        ScImmutableChainId { obj_id: self.obj_id, key_id: key.get_id() }
    }

    // get proxy for immutable ScColor field specified by key
    pub fn get_color<T: MapKey + ?Sized>(&self, key: &T) -> ScImmutableColor {
        ScImmutableColor { obj_id: self.obj_id, key_id: key.get_id() }
    }

    // get proxy for ScImmutableColorArray specified by key
    pub fn get_color_array<T: MapKey + ?Sized>(&self, key: &T) -> ScImmutableColorArray {
        let arr_id = get_object_id(self.obj_id, key.get_id(), TYPE_COLOR | TYPE_ARRAY);
        ScImmutableColorArray { obj_id: arr_id }
    }

    // get proxy for immutable ScContractId field specified by key
    pub fn get_contract_id<T: MapKey + ?Sized>(&self, key: &T) -> ScImmutableContractId {
        ScImmutableContractId { obj_id: self.obj_id, key_id: key.get_id() }
    }

    // get proxy for immutable ScHash field specified by key
    pub fn get_hash<T: MapKey + ?Sized>(&self, key: &T) -> ScImmutableHash {
        ScImmutableHash { obj_id: self.obj_id, key_id: key.get_id() }
    }

    // get proxy for ScImmutableHashArray specified by key
    pub fn get_hash_array<T: MapKey + ?Sized>(&self, key: &T) -> ScImmutableHashArray {
        let arr_id = get_object_id(self.obj_id, key.get_id(), TYPE_HASH | TYPE_ARRAY);
        ScImmutableHashArray { obj_id: arr_id }
    }

    // get proxy for immutable ScHname field specified by key
    pub fn get_hname<T: MapKey + ?Sized>(&self, key: &T) -> ScImmutableHname {
        ScImmutableHname { obj_id: self.obj_id, key_id: key.get_id() }
    }

    // get proxy for immutable int64 field specified by key
    pub fn get_int<T: MapKey + ?Sized>(&self, key: &T) -> ScImmutableInt {
        ScImmutableInt { obj_id: self.obj_id, key_id: key.get_id() }
    }

    // get proxy for ScImmutableIntArray specified by key
    pub fn get_int_array<T: MapKey + ?Sized>(&self, key: &T) -> ScImmutableIntArray {
        let arr_id = get_object_id(self.obj_id, key.get_id(), TYPE_INT | TYPE_ARRAY);
        ScImmutableIntArray { obj_id: arr_id }
    }

    // get proxy for ScImmutableMap specified by key
    pub fn get_map<T: MapKey + ?Sized>(&self, key: &T) -> ScImmutableMap {
        let map_id = get_object_id(self.obj_id, key.get_id(), TYPE_MAP);
        ScImmutableMap { obj_id: map_id }
    }

    // get proxy for ScImmutableMapArray specified by key
    pub fn get_map_array<T: MapKey + ?Sized>(&self, key: &T) -> ScImmutableMapArray {
        let arr_id = get_object_id(self.obj_id, key.get_id(), TYPE_MAP | TYPE_ARRAY);
        ScImmutableMapArray { obj_id: arr_id }
    }

    // get proxy for immutable UTF-8 text string field specified by key
    pub fn get_string<T: MapKey + ?Sized>(&self, key: &T) -> ScImmutableString {
        ScImmutableString { obj_id: self.obj_id, key_id: key.get_id() }
    }

    // get proxy for ScImmutableStringArray specified by key
    pub fn get_string_array<T: MapKey + ?Sized>(&self, key: &T) -> ScImmutableStringArray {
        let arr_id = get_object_id(self.obj_id, key.get_id(), TYPE_STRING | TYPE_ARRAY);
        ScImmutableStringArray { obj_id: arr_id }
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

// immutable array of ScMap
pub struct ScImmutableMapArray {
    pub(crate) obj_id: i32
}

impl ScImmutableMapArray {
    // index 0..length(), exclusive
    pub fn get_map(&self, index: i32) -> ScImmutableMap {
        let map_id = get_object_id(self.obj_id, Key32(index), TYPE_MAP);
        ScImmutableMap { obj_id: map_id }
    }

    // number of items in array
    pub fn length(&self) -> i32 {
        get_length(self.obj_id)
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

// proxy object for immutable UTF-8 text string in host map
pub struct ScImmutableString {
    obj_id: i32,
    key_id: Key32,
}

impl ScImmutableString {
    // check if object exists in host map
    pub fn exists(&self) -> bool {
        exists(self.obj_id, self.key_id, TYPE_STRING)
    }

    // human-readable string representation
    pub fn to_string(&self) -> String {
        self.value()
    }

    // get value from host map
    pub fn value(&self) -> String {
        let bytes = get_bytes(self.obj_id, self.key_id, TYPE_STRING);
        unsafe { String::from_utf8_unchecked(bytes) }
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

// immutable array of UTF-8 text string
pub struct ScImmutableStringArray {
    pub(crate) obj_id: i32
}

impl ScImmutableStringArray {
    // index 0..length(), exclusive
    pub fn get_string(&self, index: i32) -> ScImmutableString {
        ScImmutableString { obj_id: self.obj_id, key_id: Key32(index) }
    }

    // number of items in array
    pub fn length(&self) -> i32 {
        get_length(self.obj_id)
    }
}
