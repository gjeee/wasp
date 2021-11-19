package vmcontext

import (
	"math/big"
	"time"

	"github.com/iotaledger/wasp/packages/vm/vmcontext/vmtxbuilder"

	iotago "github.com/iotaledger/iota.go/v3"
	"github.com/iotaledger/wasp/packages/hashing"
)

func (vmctx *VMContext) BuildTransactionEssence(stateHash hashing.HashValue, timestamp time.Time) *iotago.TransactionEssence {
	return &iotago.TransactionEssence{
		Inputs:  vmctx.txbuilder.Inputs(),
		Outputs: vmctx.txbuilder.Outputs(stateHash[:]),
		Payload: nil, // ??
	}
}

func (vmctx *VMContext) createTxBuilderSnapshot(id int) *vmtxbuilder.AnchorTransactionBuilder {
	return vmctx.txbuilder.Clone()
}

func (vmctx *VMContext) restoreTxBuilderSnapshot(snapshot *vmtxbuilder.AnchorTransactionBuilder) {
	vmctx.txbuilder = snapshot
}

func (vmctx *VMContext) loadNativeTokensOnChain(id iotago.NativeTokenID) (*big.Int, iotago.UTXOInput) {
	// calls `accounts` and `blocklog` to find UTXO ID for a specific token ID, if any
	panic("not implemented")
}
