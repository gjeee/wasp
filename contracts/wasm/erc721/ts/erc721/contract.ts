// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

import * as wasmlib from "wasmlib";
import * as sc from "./index";

export class ApproveCall {
	func: wasmlib.ScFunc;
	params: sc.MutableApproveParams = new sc.MutableApproveParams(wasmlib.ScView.nilProxy);
	public constructor(ctx: wasmlib.ScFuncCallContext) {
		this.func = new wasmlib.ScFunc(ctx, sc.HScName, sc.HFuncApprove);
	}
}

export class ApproveContext {
	events: sc.Erc721Events = new sc.Erc721Events();
	params: sc.ImmutableApproveParams = new sc.ImmutableApproveParams(wasmlib.paramsProxy());
	state: sc.MutableErc721State = new sc.MutableErc721State(wasmlib.ScState.proxy());
}

export class BurnCall {
	func: wasmlib.ScFunc;
	params: sc.MutableBurnParams = new sc.MutableBurnParams(wasmlib.ScView.nilProxy);
	public constructor(ctx: wasmlib.ScFuncCallContext) {
		this.func = new wasmlib.ScFunc(ctx, sc.HScName, sc.HFuncBurn);
	}
}

export class BurnContext {
	events: sc.Erc721Events = new sc.Erc721Events();
	params: sc.ImmutableBurnParams = new sc.ImmutableBurnParams(wasmlib.paramsProxy());
	state: sc.MutableErc721State = new sc.MutableErc721State(wasmlib.ScState.proxy());
}

export class InitCall {
	func: wasmlib.ScInitFunc;
	params: sc.MutableInitParams = new sc.MutableInitParams(wasmlib.ScView.nilProxy);
	public constructor(ctx: wasmlib.ScFuncCallContext) {
		this.func = new wasmlib.ScInitFunc(ctx, sc.HScName, sc.HFuncInit);
	}
}

export class InitContext {
	events: sc.Erc721Events = new sc.Erc721Events();
	params: sc.ImmutableInitParams = new sc.ImmutableInitParams(wasmlib.paramsProxy());
	state: sc.MutableErc721State = new sc.MutableErc721State(wasmlib.ScState.proxy());
}

export class MintCall {
	func: wasmlib.ScFunc;
	params: sc.MutableMintParams = new sc.MutableMintParams(wasmlib.ScView.nilProxy);
	public constructor(ctx: wasmlib.ScFuncCallContext) {
		this.func = new wasmlib.ScFunc(ctx, sc.HScName, sc.HFuncMint);
	}
}

export class MintContext {
	events: sc.Erc721Events = new sc.Erc721Events();
	params: sc.ImmutableMintParams = new sc.ImmutableMintParams(wasmlib.paramsProxy());
	state: sc.MutableErc721State = new sc.MutableErc721State(wasmlib.ScState.proxy());
}

export class SafeTransferFromCall {
	func: wasmlib.ScFunc;
	params: sc.MutableSafeTransferFromParams = new sc.MutableSafeTransferFromParams(wasmlib.ScView.nilProxy);
	public constructor(ctx: wasmlib.ScFuncCallContext) {
		this.func = new wasmlib.ScFunc(ctx, sc.HScName, sc.HFuncSafeTransferFrom);
	}
}

export class SafeTransferFromContext {
	events: sc.Erc721Events = new sc.Erc721Events();
	params: sc.ImmutableSafeTransferFromParams = new sc.ImmutableSafeTransferFromParams(wasmlib.paramsProxy());
	state: sc.MutableErc721State = new sc.MutableErc721State(wasmlib.ScState.proxy());
}

export class SetApprovalForAllCall {
	func: wasmlib.ScFunc;
	params: sc.MutableSetApprovalForAllParams = new sc.MutableSetApprovalForAllParams(wasmlib.ScView.nilProxy);
	public constructor(ctx: wasmlib.ScFuncCallContext) {
		this.func = new wasmlib.ScFunc(ctx, sc.HScName, sc.HFuncSetApprovalForAll);
	}
}

export class SetApprovalForAllContext {
	events: sc.Erc721Events = new sc.Erc721Events();
	params: sc.ImmutableSetApprovalForAllParams = new sc.ImmutableSetApprovalForAllParams(wasmlib.paramsProxy());
	state: sc.MutableErc721State = new sc.MutableErc721State(wasmlib.ScState.proxy());
}

export class TransferFromCall {
	func: wasmlib.ScFunc;
	params: sc.MutableTransferFromParams = new sc.MutableTransferFromParams(wasmlib.ScView.nilProxy);
	public constructor(ctx: wasmlib.ScFuncCallContext) {
		this.func = new wasmlib.ScFunc(ctx, sc.HScName, sc.HFuncTransferFrom);
	}
}

export class TransferFromContext {
	events: sc.Erc721Events = new sc.Erc721Events();
	params: sc.ImmutableTransferFromParams = new sc.ImmutableTransferFromParams(wasmlib.paramsProxy());
	state: sc.MutableErc721State = new sc.MutableErc721State(wasmlib.ScState.proxy());
}

export class BalanceOfCall {
	func: wasmlib.ScView;
	params: sc.MutableBalanceOfParams = new sc.MutableBalanceOfParams(wasmlib.ScView.nilProxy);
	results: sc.ImmutableBalanceOfResults = new sc.ImmutableBalanceOfResults(wasmlib.ScView.nilProxy);
	public constructor(ctx: wasmlib.ScViewCallContext) {
		this.func = new wasmlib.ScView(ctx, sc.HScName, sc.HViewBalanceOf);
	}
}

export class BalanceOfContext {
	params: sc.ImmutableBalanceOfParams = new sc.ImmutableBalanceOfParams(wasmlib.paramsProxy());
	results: sc.MutableBalanceOfResults = new sc.MutableBalanceOfResults(wasmlib.ScView.nilProxy);
	state: sc.ImmutableErc721State = new sc.ImmutableErc721State(wasmlib.ScState.proxy());
}

export class GetApprovedCall {
	func: wasmlib.ScView;
	params: sc.MutableGetApprovedParams = new sc.MutableGetApprovedParams(wasmlib.ScView.nilProxy);
	results: sc.ImmutableGetApprovedResults = new sc.ImmutableGetApprovedResults(wasmlib.ScView.nilProxy);
	public constructor(ctx: wasmlib.ScViewCallContext) {
		this.func = new wasmlib.ScView(ctx, sc.HScName, sc.HViewGetApproved);
	}
}

export class GetApprovedContext {
	params: sc.ImmutableGetApprovedParams = new sc.ImmutableGetApprovedParams(wasmlib.paramsProxy());
	results: sc.MutableGetApprovedResults = new sc.MutableGetApprovedResults(wasmlib.ScView.nilProxy);
	state: sc.ImmutableErc721State = new sc.ImmutableErc721State(wasmlib.ScState.proxy());
}

export class IsApprovedForAllCall {
	func: wasmlib.ScView;
	params: sc.MutableIsApprovedForAllParams = new sc.MutableIsApprovedForAllParams(wasmlib.ScView.nilProxy);
	results: sc.ImmutableIsApprovedForAllResults = new sc.ImmutableIsApprovedForAllResults(wasmlib.ScView.nilProxy);
	public constructor(ctx: wasmlib.ScViewCallContext) {
		this.func = new wasmlib.ScView(ctx, sc.HScName, sc.HViewIsApprovedForAll);
	}
}

export class IsApprovedForAllContext {
	params: sc.ImmutableIsApprovedForAllParams = new sc.ImmutableIsApprovedForAllParams(wasmlib.paramsProxy());
	results: sc.MutableIsApprovedForAllResults = new sc.MutableIsApprovedForAllResults(wasmlib.ScView.nilProxy);
	state: sc.ImmutableErc721State = new sc.ImmutableErc721State(wasmlib.ScState.proxy());
}

export class NameCall {
	func: wasmlib.ScView;
	results: sc.ImmutableNameResults = new sc.ImmutableNameResults(wasmlib.ScView.nilProxy);
	public constructor(ctx: wasmlib.ScViewCallContext) {
		this.func = new wasmlib.ScView(ctx, sc.HScName, sc.HViewName);
	}
}

export class NameContext {
	results: sc.MutableNameResults = new sc.MutableNameResults(wasmlib.ScView.nilProxy);
	state: sc.ImmutableErc721State = new sc.ImmutableErc721State(wasmlib.ScState.proxy());
}

export class OwnerOfCall {
	func: wasmlib.ScView;
	params: sc.MutableOwnerOfParams = new sc.MutableOwnerOfParams(wasmlib.ScView.nilProxy);
	results: sc.ImmutableOwnerOfResults = new sc.ImmutableOwnerOfResults(wasmlib.ScView.nilProxy);
	public constructor(ctx: wasmlib.ScViewCallContext) {
		this.func = new wasmlib.ScView(ctx, sc.HScName, sc.HViewOwnerOf);
	}
}

export class OwnerOfContext {
	params: sc.ImmutableOwnerOfParams = new sc.ImmutableOwnerOfParams(wasmlib.paramsProxy());
	results: sc.MutableOwnerOfResults = new sc.MutableOwnerOfResults(wasmlib.ScView.nilProxy);
	state: sc.ImmutableErc721State = new sc.ImmutableErc721State(wasmlib.ScState.proxy());
}

export class SymbolCall {
	func: wasmlib.ScView;
	results: sc.ImmutableSymbolResults = new sc.ImmutableSymbolResults(wasmlib.ScView.nilProxy);
	public constructor(ctx: wasmlib.ScViewCallContext) {
		this.func = new wasmlib.ScView(ctx, sc.HScName, sc.HViewSymbol);
	}
}

export class SymbolContext {
	results: sc.MutableSymbolResults = new sc.MutableSymbolResults(wasmlib.ScView.nilProxy);
	state: sc.ImmutableErc721State = new sc.ImmutableErc721State(wasmlib.ScState.proxy());
}

export class TokenURICall {
	func: wasmlib.ScView;
	params: sc.MutableTokenURIParams = new sc.MutableTokenURIParams(wasmlib.ScView.nilProxy);
	results: sc.ImmutableTokenURIResults = new sc.ImmutableTokenURIResults(wasmlib.ScView.nilProxy);
	public constructor(ctx: wasmlib.ScViewCallContext) {
		this.func = new wasmlib.ScView(ctx, sc.HScName, sc.HViewTokenURI);
	}
}

export class TokenURIContext {
	params: sc.ImmutableTokenURIParams = new sc.ImmutableTokenURIParams(wasmlib.paramsProxy());
	results: sc.MutableTokenURIResults = new sc.MutableTokenURIResults(wasmlib.ScView.nilProxy);
	state: sc.ImmutableErc721State = new sc.ImmutableErc721State(wasmlib.ScState.proxy());
}

export class ScFuncs {
	static approve(ctx: wasmlib.ScFuncCallContext): ApproveCall {
		const f = new ApproveCall(ctx);
		f.params = new sc.MutableApproveParams(wasmlib.newCallParamsProxy(f.func));
		return f;
	}

	static burn(ctx: wasmlib.ScFuncCallContext): BurnCall {
		const f = new BurnCall(ctx);
		f.params = new sc.MutableBurnParams(wasmlib.newCallParamsProxy(f.func));
		return f;
	}

	static init(ctx: wasmlib.ScFuncCallContext): InitCall {
		const f = new InitCall(ctx);
		f.params = new sc.MutableInitParams(wasmlib.newCallParamsProxy(f.func));
		return f;
	}

	static mint(ctx: wasmlib.ScFuncCallContext): MintCall {
		const f = new MintCall(ctx);
		f.params = new sc.MutableMintParams(wasmlib.newCallParamsProxy(f.func));
		return f;
	}

	static safeTransferFrom(ctx: wasmlib.ScFuncCallContext): SafeTransferFromCall {
		const f = new SafeTransferFromCall(ctx);
		f.params = new sc.MutableSafeTransferFromParams(wasmlib.newCallParamsProxy(f.func));
		return f;
	}

	static setApprovalForAll(ctx: wasmlib.ScFuncCallContext): SetApprovalForAllCall {
		const f = new SetApprovalForAllCall(ctx);
		f.params = new sc.MutableSetApprovalForAllParams(wasmlib.newCallParamsProxy(f.func));
		return f;
	}

	static transferFrom(ctx: wasmlib.ScFuncCallContext): TransferFromCall {
		const f = new TransferFromCall(ctx);
		f.params = new sc.MutableTransferFromParams(wasmlib.newCallParamsProxy(f.func));
		return f;
	}

	static balanceOf(ctx: wasmlib.ScViewCallContext): BalanceOfCall {
		const f = new BalanceOfCall(ctx);
		f.params = new sc.MutableBalanceOfParams(wasmlib.newCallParamsProxy(f.func));
		f.results = new sc.ImmutableBalanceOfResults(wasmlib.newCallResultsProxy(f.func));
		return f;
	}

	static getApproved(ctx: wasmlib.ScViewCallContext): GetApprovedCall {
		const f = new GetApprovedCall(ctx);
		f.params = new sc.MutableGetApprovedParams(wasmlib.newCallParamsProxy(f.func));
		f.results = new sc.ImmutableGetApprovedResults(wasmlib.newCallResultsProxy(f.func));
		return f;
	}

	static isApprovedForAll(ctx: wasmlib.ScViewCallContext): IsApprovedForAllCall {
		const f = new IsApprovedForAllCall(ctx);
		f.params = new sc.MutableIsApprovedForAllParams(wasmlib.newCallParamsProxy(f.func));
		f.results = new sc.ImmutableIsApprovedForAllResults(wasmlib.newCallResultsProxy(f.func));
		return f;
	}

	static name(ctx: wasmlib.ScViewCallContext): NameCall {
		const f = new NameCall(ctx);
		f.results = new sc.ImmutableNameResults(wasmlib.newCallResultsProxy(f.func));
		return f;
	}

	static ownerOf(ctx: wasmlib.ScViewCallContext): OwnerOfCall {
		const f = new OwnerOfCall(ctx);
		f.params = new sc.MutableOwnerOfParams(wasmlib.newCallParamsProxy(f.func));
		f.results = new sc.ImmutableOwnerOfResults(wasmlib.newCallResultsProxy(f.func));
		return f;
	}

	static symbol(ctx: wasmlib.ScViewCallContext): SymbolCall {
		const f = new SymbolCall(ctx);
		f.results = new sc.ImmutableSymbolResults(wasmlib.newCallResultsProxy(f.func));
		return f;
	}

	static tokenURI(ctx: wasmlib.ScViewCallContext): TokenURICall {
		const f = new TokenURICall(ctx);
		f.params = new sc.MutableTokenURIParams(wasmlib.newCallParamsProxy(f.func));
		f.results = new sc.ImmutableTokenURIResults(wasmlib.newCallResultsProxy(f.func));
		return f;
	}
}
