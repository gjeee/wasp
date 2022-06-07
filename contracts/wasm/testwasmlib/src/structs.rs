// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

#![allow(dead_code)]
#![allow(unused_imports)]

use wasmlib::*;

#[derive(Clone)]
pub struct Location {
    pub x : i32,
    pub y : i32,
}

impl Location {
    pub fn from_bytes(bytes: &[u8]) -> Location {
        let mut dec = WasmDecoder::new(bytes);
        Location {
            x : int32_decode(&mut dec),
            y : int32_decode(&mut dec),
        }
    }

    pub fn to_bytes(&self) -> Vec<u8> {
        let mut enc = WasmEncoder::new();
		int32_encode(&mut enc, self.x);
		int32_encode(&mut enc, self.y);
        enc.buf()
    }
}

#[derive(Clone)]
pub struct ImmutableLocation {
    pub(crate) proxy: Proxy,
}

impl ImmutableLocation {
    pub fn exists(&self) -> bool {
        self.proxy.exists()
    }

    pub fn value(&self) -> Location {
        Location::from_bytes(&self.proxy.get())
    }
}

#[derive(Clone)]
pub struct MutableLocation {
    pub(crate) proxy: Proxy,
}

impl MutableLocation {
    pub fn delete(&self) {
        self.proxy.delete();
    }

    pub fn exists(&self) -> bool {
        self.proxy.exists()
    }

    pub fn set_value(&self, value: &Location) {
        self.proxy.set(&value.to_bytes());
    }

    pub fn value(&self) -> Location {
        Location::from_bytes(&self.proxy.get())
    }
}
