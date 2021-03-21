// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package testcore

import (
	"github.com/iotaledger/goshimmer/packages/ledgerstate"
	"github.com/iotaledger/wasp/packages/coretypes"
	"github.com/iotaledger/wasp/packages/solo"
	"github.com/iotaledger/wasp/packages/vm/core/accounts"
	"github.com/iotaledger/wasp/packages/vm/core/blob"
	"github.com/iotaledger/wasp/packages/vm/core/eventlog"
	"github.com/iotaledger/wasp/packages/vm/core/root"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestInit(t *testing.T) {
	env := solo.New(t, false, false)
	chain := env.NewChain(nil, "chain1")

	chain.AssertAccountBalance(&chain.OriginatorAgentID, ledgerstate.ColorIOTA, 1)
	env.AssertAddressBalance(chain.OriginatorAddress, ledgerstate.ColorIOTA, solo.Saldo-2)

	checkFees(chain, blob.Interface.Name, 0, 0)
	checkFees(chain, root.Interface.Name, 0, 0)
	checkFees(chain, accounts.Interface.Name, 0, 0)
	checkFees(chain, eventlog.Interface.Name, 0, 0)
}

func TestBase(t *testing.T) {
	env := solo.New(t, false, false)
	chain := env.NewChain(nil, "chain1")

	req := solo.NewCallParams(root.Interface.Name, root.FuncSetContractFee,
		root.ParamHname, blob.Interface.Hname(),
		root.ParamOwnerFee, 1,
	)
	_, err := chain.PostRequestSync(req, nil)
	require.NoError(t, err)

	chain.AssertAccountBalance(&chain.OriginatorAgentID, ledgerstate.ColorIOTA, 2)
	env.AssertAddressBalance(chain.OriginatorAddress, ledgerstate.ColorIOTA, solo.Saldo-3)

	checkFees(chain, blob.Interface.Name, 1, 0)
}

func TestFeeIsEnough1(t *testing.T) {
	env := solo.New(t, false, false)
	chain := env.NewChain(nil, "chain1")

	req := solo.NewCallParams(root.Interface.Name, root.FuncSetContractFee,
		root.ParamHname, blob.Interface.Hname(),
		root.ParamOwnerFee, 1,
	)
	_, err := chain.PostRequestSync(req, nil)
	require.NoError(t, err)

	chain.AssertAccountBalance(&chain.OriginatorAgentID, ledgerstate.ColorIOTA, 2)
	env.AssertAddressBalance(chain.OriginatorAddress, ledgerstate.ColorIOTA, solo.Saldo-3)

	checkFees(chain, blob.Interface.Name, 1, 0)

	_, err = chain.UploadBlob(nil,
		blob.VarFieldVMType, "dummyType",
		blob.VarFieldProgramBinary, "dummyBinary",
	)
	require.NoError(t, err)

	chain.AssertAccountBalance(&chain.OriginatorAgentID, ledgerstate.ColorIOTA, 3)
	env.AssertAddressBalance(chain.OriginatorAddress, ledgerstate.ColorIOTA, solo.Saldo-5)
}

func TestFeeIsEnough2(t *testing.T) {
	env := solo.New(t, false, false)
	chain := env.NewChain(nil, "chain1")

	req := solo.NewCallParams(root.Interface.Name, root.FuncSetContractFee,
		root.ParamHname, blob.Interface.Hname(),
		root.ParamOwnerFee, 2,
	)
	_, err := chain.PostRequestSync(req, nil)
	require.NoError(t, err)

	chain.AssertAccountBalance(&chain.OriginatorAgentID, ledgerstate.ColorIOTA, 2)
	env.AssertAddressBalance(chain.OriginatorAddress, ledgerstate.ColorIOTA, solo.Saldo-3)

	checkFees(chain, blob.Interface.Name, 2, 0)

	user, userAddr := env.NewKeyPairWithFunds()
	userAgentID := coretypes.NewAgentID(userAddr, 0)
	_, err = chain.UploadBlob(user,
		blob.VarFieldVMType, "dummyType",
		blob.VarFieldProgramBinary, "dummyBinary",
	)
	require.NoError(t, err)

	chain.AssertAccountBalance(&chain.OriginatorAgentID, ledgerstate.ColorIOTA, 4)
	env.AssertAddressBalance(chain.OriginatorAddress, ledgerstate.ColorIOTA, solo.Saldo-3)

	chain.AssertAccountBalance(userAgentID, ledgerstate.ColorIOTA, 1)
	env.AssertAddressBalance(userAddr, ledgerstate.ColorIOTA, solo.Saldo-3)
}
