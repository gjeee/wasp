// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

import * as wasmtypes from "wasmlib/wasmtypes";
import * as sc from "./index";

export class ImmutableApproveParams extends wasmtypes.ScProxy {
	// target account, clear approval when not present
	approved(): wasmtypes.ScImmutableAgentID {
		return new wasmtypes.ScImmutableAgentID(this.proxy.root(sc.ParamApproved));
	}

	// token ID
	tokenID(): wasmtypes.ScImmutableHash {
		return new wasmtypes.ScImmutableHash(this.proxy.root(sc.ParamTokenID));
	}
}

export class MutableApproveParams extends wasmtypes.ScProxy {
	// target account, clear approval when not present
	approved(): wasmtypes.ScMutableAgentID {
		return new wasmtypes.ScMutableAgentID(this.proxy.root(sc.ParamApproved));
	}

	// token ID
	tokenID(): wasmtypes.ScMutableHash {
		return new wasmtypes.ScMutableHash(this.proxy.root(sc.ParamTokenID));
	}
}

export class ImmutableBurnParams extends wasmtypes.ScProxy {
	// token ID
	tokenID(): wasmtypes.ScImmutableHash {
		return new wasmtypes.ScImmutableHash(this.proxy.root(sc.ParamTokenID));
	}
}

export class MutableBurnParams extends wasmtypes.ScProxy {
	// token ID
	tokenID(): wasmtypes.ScMutableHash {
		return new wasmtypes.ScMutableHash(this.proxy.root(sc.ParamTokenID));
	}
}

export class ImmutableInitParams extends wasmtypes.ScProxy {
	// creator/owner of the initial supply
	name(): wasmtypes.ScImmutableString {
		return new wasmtypes.ScImmutableString(this.proxy.root(sc.ParamName));
	}

	// initial token supply
	symbol(): wasmtypes.ScImmutableString {
		return new wasmtypes.ScImmutableString(this.proxy.root(sc.ParamSymbol));
	}
}

export class MutableInitParams extends wasmtypes.ScProxy {
	// creator/owner of the initial supply
	name(): wasmtypes.ScMutableString {
		return new wasmtypes.ScMutableString(this.proxy.root(sc.ParamName));
	}

	// initial token supply
	symbol(): wasmtypes.ScMutableString {
		return new wasmtypes.ScMutableString(this.proxy.root(sc.ParamSymbol));
	}
}

export class ImmutableMintParams extends wasmtypes.ScProxy {
	// New token id
	tokenID(): wasmtypes.ScImmutableHash {
		return new wasmtypes.ScImmutableHash(this.proxy.root(sc.ParamTokenID));
	}

	// Optional token URI that overrides default
	tokenURI(): wasmtypes.ScImmutableString {
		return new wasmtypes.ScImmutableString(this.proxy.root(sc.ParamTokenURI));
	}
}

export class MutableMintParams extends wasmtypes.ScProxy {
	// New token id
	tokenID(): wasmtypes.ScMutableHash {
		return new wasmtypes.ScMutableHash(this.proxy.root(sc.ParamTokenID));
	}

	// Optional token URI that overrides default
	tokenURI(): wasmtypes.ScMutableString {
		return new wasmtypes.ScMutableString(this.proxy.root(sc.ParamTokenURI));
	}
}

export class ImmutableSafeTransferFromParams extends wasmtypes.ScProxy {
	// extra data to pass to SC
	data(): wasmtypes.ScImmutableBytes {
		return new wasmtypes.ScImmutableBytes(this.proxy.root(sc.ParamData));
	}

	// from account
	from(): wasmtypes.ScImmutableAgentID {
		return new wasmtypes.ScImmutableAgentID(this.proxy.root(sc.ParamFrom));
	}

	// to account, which is SC
	to(): wasmtypes.ScImmutableAgentID {
		return new wasmtypes.ScImmutableAgentID(this.proxy.root(sc.ParamTo));
	}

	// token ID
	tokenID(): wasmtypes.ScImmutableHash {
		return new wasmtypes.ScImmutableHash(this.proxy.root(sc.ParamTokenID));
	}
}

export class MutableSafeTransferFromParams extends wasmtypes.ScProxy {
	// extra data to pass to SC
	data(): wasmtypes.ScMutableBytes {
		return new wasmtypes.ScMutableBytes(this.proxy.root(sc.ParamData));
	}

	// from account
	from(): wasmtypes.ScMutableAgentID {
		return new wasmtypes.ScMutableAgentID(this.proxy.root(sc.ParamFrom));
	}

	// to account, which is SC
	to(): wasmtypes.ScMutableAgentID {
		return new wasmtypes.ScMutableAgentID(this.proxy.root(sc.ParamTo));
	}

	// token ID
	tokenID(): wasmtypes.ScMutableHash {
		return new wasmtypes.ScMutableHash(this.proxy.root(sc.ParamTokenID));
	}
}

export class ImmutableSetApprovalForAllParams extends wasmtypes.ScProxy {
	approval(): wasmtypes.ScImmutableBool {
		return new wasmtypes.ScImmutableBool(this.proxy.root(sc.ParamApproval));
	}

	// target operator of account
	operator(): wasmtypes.ScImmutableAgentID {
		return new wasmtypes.ScImmutableAgentID(this.proxy.root(sc.ParamOperator));
	}
}

export class MutableSetApprovalForAllParams extends wasmtypes.ScProxy {
	approval(): wasmtypes.ScMutableBool {
		return new wasmtypes.ScMutableBool(this.proxy.root(sc.ParamApproval));
	}

	// target operator of account
	operator(): wasmtypes.ScMutableAgentID {
		return new wasmtypes.ScMutableAgentID(this.proxy.root(sc.ParamOperator));
	}
}

export class ImmutableTransferFromParams extends wasmtypes.ScProxy {
	// from account
	from(): wasmtypes.ScImmutableAgentID {
		return new wasmtypes.ScImmutableAgentID(this.proxy.root(sc.ParamFrom));
	}

	// to account
	to(): wasmtypes.ScImmutableAgentID {
		return new wasmtypes.ScImmutableAgentID(this.proxy.root(sc.ParamTo));
	}

	// token ID
	tokenID(): wasmtypes.ScImmutableHash {
		return new wasmtypes.ScImmutableHash(this.proxy.root(sc.ParamTokenID));
	}
}

export class MutableTransferFromParams extends wasmtypes.ScProxy {
	// from account
	from(): wasmtypes.ScMutableAgentID {
		return new wasmtypes.ScMutableAgentID(this.proxy.root(sc.ParamFrom));
	}

	// to account
	to(): wasmtypes.ScMutableAgentID {
		return new wasmtypes.ScMutableAgentID(this.proxy.root(sc.ParamTo));
	}

	// token ID
	tokenID(): wasmtypes.ScMutableHash {
		return new wasmtypes.ScMutableHash(this.proxy.root(sc.ParamTokenID));
	}
}

export class ImmutableBalanceOfParams extends wasmtypes.ScProxy {
	// account owner
	owner(): wasmtypes.ScImmutableAgentID {
		return new wasmtypes.ScImmutableAgentID(this.proxy.root(sc.ParamOwner));
	}
}

export class MutableBalanceOfParams extends wasmtypes.ScProxy {
	// account owner
	owner(): wasmtypes.ScMutableAgentID {
		return new wasmtypes.ScMutableAgentID(this.proxy.root(sc.ParamOwner));
	}
}

export class ImmutableGetApprovedParams extends wasmtypes.ScProxy {
	tokenID(): wasmtypes.ScImmutableHash {
		return new wasmtypes.ScImmutableHash(this.proxy.root(sc.ParamTokenID));
	}
}

export class MutableGetApprovedParams extends wasmtypes.ScProxy {
	tokenID(): wasmtypes.ScMutableHash {
		return new wasmtypes.ScMutableHash(this.proxy.root(sc.ParamTokenID));
	}
}

export class ImmutableIsApprovedForAllParams extends wasmtypes.ScProxy {
	operator(): wasmtypes.ScImmutableAgentID {
		return new wasmtypes.ScImmutableAgentID(this.proxy.root(sc.ParamOperator));
	}

	owner(): wasmtypes.ScImmutableAgentID {
		return new wasmtypes.ScImmutableAgentID(this.proxy.root(sc.ParamOwner));
	}
}

export class MutableIsApprovedForAllParams extends wasmtypes.ScProxy {
	operator(): wasmtypes.ScMutableAgentID {
		return new wasmtypes.ScMutableAgentID(this.proxy.root(sc.ParamOperator));
	}

	owner(): wasmtypes.ScMutableAgentID {
		return new wasmtypes.ScMutableAgentID(this.proxy.root(sc.ParamOwner));
	}
}

export class ImmutableOwnerOfParams extends wasmtypes.ScProxy {
	tokenID(): wasmtypes.ScImmutableHash {
		return new wasmtypes.ScImmutableHash(this.proxy.root(sc.ParamTokenID));
	}
}

export class MutableOwnerOfParams extends wasmtypes.ScProxy {
	tokenID(): wasmtypes.ScMutableHash {
		return new wasmtypes.ScMutableHash(this.proxy.root(sc.ParamTokenID));
	}
}

export class ImmutableTokenURIParams extends wasmtypes.ScProxy {
	tokenID(): wasmtypes.ScImmutableHash {
		return new wasmtypes.ScImmutableHash(this.proxy.root(sc.ParamTokenID));
	}
}

export class MutableTokenURIParams extends wasmtypes.ScProxy {
	tokenID(): wasmtypes.ScMutableHash {
		return new wasmtypes.ScMutableHash(this.proxy.root(sc.ParamTokenID));
	}
}
