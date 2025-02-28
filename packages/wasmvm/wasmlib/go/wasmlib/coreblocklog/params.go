// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

//nolint:revive
package coreblocklog

import "github.com/iotaledger/wasp/packages/wasmvm/wasmlib/go/wasmlib/wasmtypes"

type ImmutableGetBlockInfoParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableGetBlockInfoParams) BlockIndex() wasmtypes.ScImmutableUint32 {
	return wasmtypes.NewScImmutableUint32(s.proxy.Root(ParamBlockIndex))
}

type MutableGetBlockInfoParams struct {
	proxy wasmtypes.Proxy
}

func (s MutableGetBlockInfoParams) BlockIndex() wasmtypes.ScMutableUint32 {
	return wasmtypes.NewScMutableUint32(s.proxy.Root(ParamBlockIndex))
}

type ImmutableGetEventsForBlockParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableGetEventsForBlockParams) BlockIndex() wasmtypes.ScImmutableUint32 {
	return wasmtypes.NewScImmutableUint32(s.proxy.Root(ParamBlockIndex))
}

type MutableGetEventsForBlockParams struct {
	proxy wasmtypes.Proxy
}

func (s MutableGetEventsForBlockParams) BlockIndex() wasmtypes.ScMutableUint32 {
	return wasmtypes.NewScMutableUint32(s.proxy.Root(ParamBlockIndex))
}

type ImmutableGetEventsForContractParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableGetEventsForContractParams) ContractHname() wasmtypes.ScImmutableHname {
	return wasmtypes.NewScImmutableHname(s.proxy.Root(ParamContractHname))
}

func (s ImmutableGetEventsForContractParams) FromBlock() wasmtypes.ScImmutableUint32 {
	return wasmtypes.NewScImmutableUint32(s.proxy.Root(ParamFromBlock))
}

func (s ImmutableGetEventsForContractParams) ToBlock() wasmtypes.ScImmutableUint32 {
	return wasmtypes.NewScImmutableUint32(s.proxy.Root(ParamToBlock))
}

type MutableGetEventsForContractParams struct {
	proxy wasmtypes.Proxy
}

func (s MutableGetEventsForContractParams) ContractHname() wasmtypes.ScMutableHname {
	return wasmtypes.NewScMutableHname(s.proxy.Root(ParamContractHname))
}

func (s MutableGetEventsForContractParams) FromBlock() wasmtypes.ScMutableUint32 {
	return wasmtypes.NewScMutableUint32(s.proxy.Root(ParamFromBlock))
}

func (s MutableGetEventsForContractParams) ToBlock() wasmtypes.ScMutableUint32 {
	return wasmtypes.NewScMutableUint32(s.proxy.Root(ParamToBlock))
}

type ImmutableGetEventsForRequestParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableGetEventsForRequestParams) RequestID() wasmtypes.ScImmutableRequestID {
	return wasmtypes.NewScImmutableRequestID(s.proxy.Root(ParamRequestID))
}

type MutableGetEventsForRequestParams struct {
	proxy wasmtypes.Proxy
}

func (s MutableGetEventsForRequestParams) RequestID() wasmtypes.ScMutableRequestID {
	return wasmtypes.NewScMutableRequestID(s.proxy.Root(ParamRequestID))
}

type ImmutableGetRequestIDsForBlockParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableGetRequestIDsForBlockParams) BlockIndex() wasmtypes.ScImmutableUint32 {
	return wasmtypes.NewScImmutableUint32(s.proxy.Root(ParamBlockIndex))
}

type MutableGetRequestIDsForBlockParams struct {
	proxy wasmtypes.Proxy
}

func (s MutableGetRequestIDsForBlockParams) BlockIndex() wasmtypes.ScMutableUint32 {
	return wasmtypes.NewScMutableUint32(s.proxy.Root(ParamBlockIndex))
}

type ImmutableGetRequestReceiptParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableGetRequestReceiptParams) RequestID() wasmtypes.ScImmutableRequestID {
	return wasmtypes.NewScImmutableRequestID(s.proxy.Root(ParamRequestID))
}

type MutableGetRequestReceiptParams struct {
	proxy wasmtypes.Proxy
}

func (s MutableGetRequestReceiptParams) RequestID() wasmtypes.ScMutableRequestID {
	return wasmtypes.NewScMutableRequestID(s.proxy.Root(ParamRequestID))
}

type ImmutableGetRequestReceiptsForBlockParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableGetRequestReceiptsForBlockParams) BlockIndex() wasmtypes.ScImmutableUint32 {
	return wasmtypes.NewScImmutableUint32(s.proxy.Root(ParamBlockIndex))
}

type MutableGetRequestReceiptsForBlockParams struct {
	proxy wasmtypes.Proxy
}

func (s MutableGetRequestReceiptsForBlockParams) BlockIndex() wasmtypes.ScMutableUint32 {
	return wasmtypes.NewScMutableUint32(s.proxy.Root(ParamBlockIndex))
}

type ImmutableIsRequestProcessedParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableIsRequestProcessedParams) RequestID() wasmtypes.ScImmutableRequestID {
	return wasmtypes.NewScImmutableRequestID(s.proxy.Root(ParamRequestID))
}

type MutableIsRequestProcessedParams struct {
	proxy wasmtypes.Proxy
}

func (s MutableIsRequestProcessedParams) RequestID() wasmtypes.ScMutableRequestID {
	return wasmtypes.NewScMutableRequestID(s.proxy.Root(ParamRequestID))
}
