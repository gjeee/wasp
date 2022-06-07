// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

import * as wasmlib from "wasmlib";
import * as sc from "./index";

const exportMap: wasmlib.ScExportMap = {
	names: [
		sc.FuncApprove,
		sc.FuncBurn,
		sc.FuncInit,
		sc.FuncMint,
		sc.FuncSafeTransferFrom,
		sc.FuncSetApprovalForAll,
		sc.FuncTransferFrom,
		sc.ViewBalanceOf,
		sc.ViewGetApproved,
		sc.ViewIsApprovedForAll,
		sc.ViewName,
		sc.ViewOwnerOf,
		sc.ViewSymbol,
		sc.ViewTokenURI,
	],
	funcs: [
		funcApproveThunk,
		funcBurnThunk,
		funcInitThunk,
		funcMintThunk,
		funcSafeTransferFromThunk,
		funcSetApprovalForAllThunk,
		funcTransferFromThunk,
	],
	views: [
		viewBalanceOfThunk,
		viewGetApprovedThunk,
		viewIsApprovedForAllThunk,
		viewNameThunk,
		viewOwnerOfThunk,
		viewSymbolThunk,
		viewTokenURIThunk,
	],
};

export function on_call(index: i32): void {
	wasmlib.WasmVMHost.connect();
	wasmlib.ScExports.call(index, exportMap);
}

export function on_load(): void {
	wasmlib.WasmVMHost.connect();
	wasmlib.ScExports.export(exportMap);
}

function funcApproveThunk(ctx: wasmlib.ScFuncContext): void {
	ctx.log("erc721.funcApprove");
	let f = new sc.ApproveContext();
	ctx.require(f.params.tokenID().exists(), "missing mandatory tokenID");
	sc.funcApprove(ctx, f);
	ctx.log("erc721.funcApprove ok");
}

function funcBurnThunk(ctx: wasmlib.ScFuncContext): void {
	ctx.log("erc721.funcBurn");
	let f = new sc.BurnContext();
	ctx.require(f.params.tokenID().exists(), "missing mandatory tokenID");
	sc.funcBurn(ctx, f);
	ctx.log("erc721.funcBurn ok");
}

function funcInitThunk(ctx: wasmlib.ScFuncContext): void {
	ctx.log("erc721.funcInit");
	let f = new sc.InitContext();
	ctx.require(f.params.name().exists(), "missing mandatory name");
	ctx.require(f.params.symbol().exists(), "missing mandatory symbol");
	sc.funcInit(ctx, f);
	ctx.log("erc721.funcInit ok");
}

function funcMintThunk(ctx: wasmlib.ScFuncContext): void {
	ctx.log("erc721.funcMint");
	let f = new sc.MintContext();
	ctx.require(f.params.tokenID().exists(), "missing mandatory tokenID");
	sc.funcMint(ctx, f);
	ctx.log("erc721.funcMint ok");
}

function funcSafeTransferFromThunk(ctx: wasmlib.ScFuncContext): void {
	ctx.log("erc721.funcSafeTransferFrom");
	let f = new sc.SafeTransferFromContext();
	ctx.require(f.params.from().exists(), "missing mandatory from");
	ctx.require(f.params.to().exists(), "missing mandatory to");
	ctx.require(f.params.tokenID().exists(), "missing mandatory tokenID");
	sc.funcSafeTransferFrom(ctx, f);
	ctx.log("erc721.funcSafeTransferFrom ok");
}

function funcSetApprovalForAllThunk(ctx: wasmlib.ScFuncContext): void {
	ctx.log("erc721.funcSetApprovalForAll");
	let f = new sc.SetApprovalForAllContext();
	ctx.require(f.params.approval().exists(), "missing mandatory approval");
	ctx.require(f.params.operator().exists(), "missing mandatory operator");
	sc.funcSetApprovalForAll(ctx, f);
	ctx.log("erc721.funcSetApprovalForAll ok");
}

function funcTransferFromThunk(ctx: wasmlib.ScFuncContext): void {
	ctx.log("erc721.funcTransferFrom");
	let f = new sc.TransferFromContext();
	ctx.require(f.params.from().exists(), "missing mandatory from");
	ctx.require(f.params.to().exists(), "missing mandatory to");
	ctx.require(f.params.tokenID().exists(), "missing mandatory tokenID");
	sc.funcTransferFrom(ctx, f);
	ctx.log("erc721.funcTransferFrom ok");
}

function viewBalanceOfThunk(ctx: wasmlib.ScViewContext): void {
	ctx.log("erc721.viewBalanceOf");
	let f = new sc.BalanceOfContext();
	const results = new wasmlib.ScDict([]);
	f.results = new sc.MutableBalanceOfResults(results.asProxy());
	ctx.require(f.params.owner().exists(), "missing mandatory owner");
	sc.viewBalanceOf(ctx, f);
	ctx.results(results);
	ctx.log("erc721.viewBalanceOf ok");
}

function viewGetApprovedThunk(ctx: wasmlib.ScViewContext): void {
	ctx.log("erc721.viewGetApproved");
	let f = new sc.GetApprovedContext();
	const results = new wasmlib.ScDict([]);
	f.results = new sc.MutableGetApprovedResults(results.asProxy());
	ctx.require(f.params.tokenID().exists(), "missing mandatory tokenID");
	sc.viewGetApproved(ctx, f);
	ctx.results(results);
	ctx.log("erc721.viewGetApproved ok");
}

function viewIsApprovedForAllThunk(ctx: wasmlib.ScViewContext): void {
	ctx.log("erc721.viewIsApprovedForAll");
	let f = new sc.IsApprovedForAllContext();
	const results = new wasmlib.ScDict([]);
	f.results = new sc.MutableIsApprovedForAllResults(results.asProxy());
	ctx.require(f.params.operator().exists(), "missing mandatory operator");
	ctx.require(f.params.owner().exists(), "missing mandatory owner");
	sc.viewIsApprovedForAll(ctx, f);
	ctx.results(results);
	ctx.log("erc721.viewIsApprovedForAll ok");
}

function viewNameThunk(ctx: wasmlib.ScViewContext): void {
	ctx.log("erc721.viewName");
	let f = new sc.NameContext();
	const results = new wasmlib.ScDict([]);
	f.results = new sc.MutableNameResults(results.asProxy());
	sc.viewName(ctx, f);
	ctx.results(results);
	ctx.log("erc721.viewName ok");
}

function viewOwnerOfThunk(ctx: wasmlib.ScViewContext): void {
	ctx.log("erc721.viewOwnerOf");
	let f = new sc.OwnerOfContext();
	const results = new wasmlib.ScDict([]);
	f.results = new sc.MutableOwnerOfResults(results.asProxy());
	ctx.require(f.params.tokenID().exists(), "missing mandatory tokenID");
	sc.viewOwnerOf(ctx, f);
	ctx.results(results);
	ctx.log("erc721.viewOwnerOf ok");
}

function viewSymbolThunk(ctx: wasmlib.ScViewContext): void {
	ctx.log("erc721.viewSymbol");
	let f = new sc.SymbolContext();
	const results = new wasmlib.ScDict([]);
	f.results = new sc.MutableSymbolResults(results.asProxy());
	sc.viewSymbol(ctx, f);
	ctx.results(results);
	ctx.log("erc721.viewSymbol ok");
}

function viewTokenURIThunk(ctx: wasmlib.ScViewContext): void {
	ctx.log("erc721.viewTokenURI");
	let f = new sc.TokenURIContext();
	const results = new wasmlib.ScDict([]);
	f.results = new sc.MutableTokenURIResults(results.asProxy());
	ctx.require(f.params.tokenID().exists(), "missing mandatory tokenID");
	sc.viewTokenURI(ctx, f);
	ctx.results(results);
	ctx.log("erc721.viewTokenURI ok");
}
