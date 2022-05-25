// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

import * as wasmtypes from "wasmlib/wasmtypes";
import * as sc from "./index";

export class ImmutableFoundryCreateNewResults extends wasmtypes.ScProxy {
	foundrySN(): wasmtypes.ScImmutableUint32 {
		return new wasmtypes.ScImmutableUint32(this.proxy.root(sc.ResultFoundrySN));
	}
}

export class MutableFoundryCreateNewResults extends wasmtypes.ScProxy {
	foundrySN(): wasmtypes.ScMutableUint32 {
		return new wasmtypes.ScMutableUint32(this.proxy.root(sc.ResultFoundrySN));
	}
}

export class ArrayOfImmutableNftID extends wasmtypes.ScProxy {

	length(): u32 {
		return this.proxy.length();
	}

	getNftID(index: u32): wasmtypes.ScImmutableNftID {
		return new wasmtypes.ScImmutableNftID(this.proxy.index(index));
	}
}

export class ImmutableAccountNFTsResults extends wasmtypes.ScProxy {
	nftIDs(): sc.ArrayOfImmutableNftID {
		return new sc.ArrayOfImmutableNftID(this.proxy.root(sc.ResultNftIDs));
	}
}

export class ArrayOfMutableNftID extends wasmtypes.ScProxy {

	appendNftID(): wasmtypes.ScMutableNftID {
		return new wasmtypes.ScMutableNftID(this.proxy.append());
	}

	clear(): void {
		this.proxy.clearArray();
	}

	length(): u32 {
		return this.proxy.length();
	}

	getNftID(index: u32): wasmtypes.ScMutableNftID {
		return new wasmtypes.ScMutableNftID(this.proxy.index(index));
	}
}

export class MutableAccountNFTsResults extends wasmtypes.ScProxy {
	nftIDs(): sc.ArrayOfMutableNftID {
		return new sc.ArrayOfMutableNftID(this.proxy.root(sc.ResultNftIDs));
	}
}

export class MapAgentIDToImmutableBool extends wasmtypes.ScProxy {

	getBool(key: wasmtypes.ScAgentID): wasmtypes.ScImmutableBool {
		return new wasmtypes.ScImmutableBool(this.proxy.key(wasmtypes.agentIDToBytes(key)));
	}
}

export class ImmutableAccountsResults extends wasmtypes.ScProxy {
	allAccounts(): sc.MapAgentIDToImmutableBool {
		return new sc.MapAgentIDToImmutableBool(this.proxy);
	}
}

export class MapAgentIDToMutableBool extends wasmtypes.ScProxy {

	clear(): void {
		this.proxy.clearMap();
	}

	getBool(key: wasmtypes.ScAgentID): wasmtypes.ScMutableBool {
		return new wasmtypes.ScMutableBool(this.proxy.key(wasmtypes.agentIDToBytes(key)));
	}
}

export class MutableAccountsResults extends wasmtypes.ScProxy {
	allAccounts(): sc.MapAgentIDToMutableBool {
		return new sc.MapAgentIDToMutableBool(this.proxy);
	}
}

export class MapTokenIDToImmutableBigInt extends wasmtypes.ScProxy {

	getBigInt(key: wasmtypes.ScTokenID): wasmtypes.ScImmutableBigInt {
		return new wasmtypes.ScImmutableBigInt(this.proxy.key(wasmtypes.tokenIDToBytes(key)));
	}
}

export class ImmutableBalanceResults extends wasmtypes.ScProxy {
	balances(): sc.MapTokenIDToImmutableBigInt {
		return new sc.MapTokenIDToImmutableBigInt(this.proxy);
	}
}

export class MapTokenIDToMutableBigInt extends wasmtypes.ScProxy {

	clear(): void {
		this.proxy.clearMap();
	}

	getBigInt(key: wasmtypes.ScTokenID): wasmtypes.ScMutableBigInt {
		return new wasmtypes.ScMutableBigInt(this.proxy.key(wasmtypes.tokenIDToBytes(key)));
	}
}

export class MutableBalanceResults extends wasmtypes.ScProxy {
	balances(): sc.MapTokenIDToMutableBigInt {
		return new sc.MapTokenIDToMutableBigInt(this.proxy);
	}
}

export class ImmutableFoundryOutputResults extends wasmtypes.ScProxy {
	foundryOutputBin(): wasmtypes.ScImmutableBytes {
		return new wasmtypes.ScImmutableBytes(this.proxy.root(sc.ResultFoundryOutputBin));
	}
}

export class MutableFoundryOutputResults extends wasmtypes.ScProxy {
	foundryOutputBin(): wasmtypes.ScMutableBytes {
		return new wasmtypes.ScMutableBytes(this.proxy.root(sc.ResultFoundryOutputBin));
	}
}

export class ImmutableGetAccountNonceResults extends wasmtypes.ScProxy {
	accountNonce(): wasmtypes.ScImmutableUint64 {
		return new wasmtypes.ScImmutableUint64(this.proxy.root(sc.ResultAccountNonce));
	}
}

export class MutableGetAccountNonceResults extends wasmtypes.ScProxy {
	accountNonce(): wasmtypes.ScMutableUint64 {
		return new wasmtypes.ScMutableUint64(this.proxy.root(sc.ResultAccountNonce));
	}
}

export class MapTokenIDToImmutableBool extends wasmtypes.ScProxy {

	getBool(key: wasmtypes.ScTokenID): wasmtypes.ScImmutableBool {
		return new wasmtypes.ScImmutableBool(this.proxy.key(wasmtypes.tokenIDToBytes(key)));
	}
}

export class ImmutableGetNativeTokenIDRegistryResults extends wasmtypes.ScProxy {
	mapping(): sc.MapTokenIDToImmutableBool {
		return new sc.MapTokenIDToImmutableBool(this.proxy);
	}
}

export class MapTokenIDToMutableBool extends wasmtypes.ScProxy {

	clear(): void {
		this.proxy.clearMap();
	}

	getBool(key: wasmtypes.ScTokenID): wasmtypes.ScMutableBool {
		return new wasmtypes.ScMutableBool(this.proxy.key(wasmtypes.tokenIDToBytes(key)));
	}
}

export class MutableGetNativeTokenIDRegistryResults extends wasmtypes.ScProxy {
	mapping(): sc.MapTokenIDToMutableBool {
		return new sc.MapTokenIDToMutableBool(this.proxy);
	}
}

export class ImmutableNftDataResults extends wasmtypes.ScProxy {
	nftData(): wasmtypes.ScImmutableBytes {
		return new wasmtypes.ScImmutableBytes(this.proxy.root(sc.ResultNftData));
	}
}

export class MutableNftDataResults extends wasmtypes.ScProxy {
	nftData(): wasmtypes.ScMutableBytes {
		return new wasmtypes.ScMutableBytes(this.proxy.root(sc.ResultNftData));
	}
}

export class ImmutableTotalAssetsResults extends wasmtypes.ScProxy {
	assets(): sc.MapTokenIDToImmutableBigInt {
		return new sc.MapTokenIDToImmutableBigInt(this.proxy);
	}
}

export class MutableTotalAssetsResults extends wasmtypes.ScProxy {
	assets(): sc.MapTokenIDToMutableBigInt {
		return new sc.MapTokenIDToMutableBigInt(this.proxy);
	}
}
