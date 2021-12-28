// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

#![allow(dead_code)]
#![allow(unused_imports)]

use wasmlib::*;
use wasmlib::host::*;

use crate::*;
use crate::keys::*;
use crate::typedefs::*;

pub struct MapHashToImmutableAgentID {
	pub(crate) obj_id: i32,
}

impl MapHashToImmutableAgentID {
    pub fn get_agent_id(&self, key: &ScHash) -> ScImmutableAgentID {
        ScImmutableAgentID::new(self.obj_id, key.get_key_id())
    }
}

pub struct MapAgentIDToImmutableOperators {
	pub(crate) obj_id: i32,
}

impl MapAgentIDToImmutableOperators {
    pub fn get_operators(&self, key: &ScAgentID) -> ImmutableOperators {
        let sub_id = get_object_id(self.obj_id, key.get_key_id(), TYPE_MAP);
        ImmutableOperators { obj_id: sub_id }
    }
}

pub struct MapAgentIDToImmutableUint64 {
	pub(crate) obj_id: i32,
}

impl MapAgentIDToImmutableUint64 {
    pub fn get_uint64(&self, key: &ScAgentID) -> ScImmutableUint64 {
        ScImmutableUint64::new(self.obj_id, key.get_key_id())
    }
}

#[derive(Clone, Copy)]
pub struct ImmutableErc721State {
    pub(crate) id: i32,
}

impl ImmutableErc721State {
    pub fn approved_accounts(&self) -> MapHashToImmutableAgentID {
		let map_id = get_object_id(self.id, idx_map(IDX_STATE_APPROVED_ACCOUNTS), TYPE_MAP);
		MapHashToImmutableAgentID { obj_id: map_id }
	}

    pub fn approved_operators(&self) -> MapAgentIDToImmutableOperators {
		let map_id = get_object_id(self.id, idx_map(IDX_STATE_APPROVED_OPERATORS), TYPE_MAP);
		MapAgentIDToImmutableOperators { obj_id: map_id }
	}

    pub fn balances(&self) -> MapAgentIDToImmutableUint64 {
		let map_id = get_object_id(self.id, idx_map(IDX_STATE_BALANCES), TYPE_MAP);
		MapAgentIDToImmutableUint64 { obj_id: map_id }
	}

    pub fn name(&self) -> ScImmutableString {
		ScImmutableString::new(self.id, idx_map(IDX_STATE_NAME))
	}

    pub fn owners(&self) -> MapHashToImmutableAgentID {
		let map_id = get_object_id(self.id, idx_map(IDX_STATE_OWNERS), TYPE_MAP);
		MapHashToImmutableAgentID { obj_id: map_id }
	}

    pub fn symbol(&self) -> ScImmutableString {
		ScImmutableString::new(self.id, idx_map(IDX_STATE_SYMBOL))
	}
}

pub struct MapHashToMutableAgentID {
	pub(crate) obj_id: i32,
}

impl MapHashToMutableAgentID {
    pub fn clear(&self) {
        clear(self.obj_id);
    }

    pub fn get_agent_id(&self, key: &ScHash) -> ScMutableAgentID {
        ScMutableAgentID::new(self.obj_id, key.get_key_id())
    }
}

pub struct MapAgentIDToMutableOperators {
	pub(crate) obj_id: i32,
}

impl MapAgentIDToMutableOperators {
    pub fn clear(&self) {
        clear(self.obj_id);
    }

    pub fn get_operators(&self, key: &ScAgentID) -> MutableOperators {
        let sub_id = get_object_id(self.obj_id, key.get_key_id(), TYPE_MAP);
        MutableOperators { obj_id: sub_id }
    }
}

pub struct MapAgentIDToMutableUint64 {
	pub(crate) obj_id: i32,
}

impl MapAgentIDToMutableUint64 {
    pub fn clear(&self) {
        clear(self.obj_id);
    }

    pub fn get_uint64(&self, key: &ScAgentID) -> ScMutableUint64 {
        ScMutableUint64::new(self.obj_id, key.get_key_id())
    }
}

#[derive(Clone, Copy)]
pub struct MutableErc721State {
    pub(crate) id: i32,
}

impl MutableErc721State {
    pub fn approved_accounts(&self) -> MapHashToMutableAgentID {
		let map_id = get_object_id(self.id, idx_map(IDX_STATE_APPROVED_ACCOUNTS), TYPE_MAP);
		MapHashToMutableAgentID { obj_id: map_id }
	}

    pub fn approved_operators(&self) -> MapAgentIDToMutableOperators {
		let map_id = get_object_id(self.id, idx_map(IDX_STATE_APPROVED_OPERATORS), TYPE_MAP);
		MapAgentIDToMutableOperators { obj_id: map_id }
	}

    pub fn balances(&self) -> MapAgentIDToMutableUint64 {
		let map_id = get_object_id(self.id, idx_map(IDX_STATE_BALANCES), TYPE_MAP);
		MapAgentIDToMutableUint64 { obj_id: map_id }
	}

    pub fn name(&self) -> ScMutableString {
		ScMutableString::new(self.id, idx_map(IDX_STATE_NAME))
	}

    pub fn owners(&self) -> MapHashToMutableAgentID {
		let map_id = get_object_id(self.id, idx_map(IDX_STATE_OWNERS), TYPE_MAP);
		MapHashToMutableAgentID { obj_id: map_id }
	}

    pub fn symbol(&self) -> ScMutableString {
		ScMutableString::new(self.id, idx_map(IDX_STATE_SYMBOL))
	}
}
