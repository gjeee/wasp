package state

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/iotaledger/wasp/packages/coretypes"
	"github.com/iotaledger/wasp/packages/kv"
	"github.com/iotaledger/wasp/packages/kv/dict"
	"github.com/iotaledger/wasp/packages/kv/optimism"
	"github.com/iotaledger/wasp/packages/webapi/httperrors"
	"github.com/iotaledger/wasp/packages/webapi/routes"
	"github.com/iotaledger/wasp/packages/webapi/webapiutil"
	"github.com/iotaledger/wasp/plugins/chains"
	"github.com/labstack/echo/v4"
	"github.com/pangpanglabs/echoswagger/v2"
)

func AddEndpoints(server echoswagger.ApiRouter) {
	dictExample := dict.Dict{
		kv.Key("key1"): []byte("value1"),
	}.JSONDict()

	server.GET(routes.CallView(":chainID", ":contractHname", ":fname"), handleCallView).
		SetSummary("Call a view function on a contract").
		AddParamPath("", "chainID", "ChainID (base58-encoded)").
		AddParamPath("", "contractHname", "Contract Hname").
		AddParamPath("getInfo", "fname", "Function name").
		AddParamBody(dictExample, "params", "Parameters", false).
		AddResponse(http.StatusOK, "Result", dictExample, nil)

	server.GET(routes.StateGet(":chainID", ":key"), handleStateGet).
		SetSummary("Fetch the raw value associated with the given key in the chain state").
		AddParamPath("", "chainID", "ChainID (base58-encoded)").
		AddParamPath("", "key", "Key (hex-encoded)").
		AddResponse(http.StatusOK, "Result", []byte("value"), nil)
}

func handleCallView(c echo.Context) error {
	chainID, err := coretypes.ChainIDFromBase58(c.Param("chainID"))
	if err != nil {
		return httperrors.BadRequest(fmt.Sprintf("Invalid chain ID: %+v", c.Param("chainID")))
	}
	contractHname, err := coretypes.HnameFromString(c.Param("contractHname"))
	if err != nil {
		return httperrors.BadRequest(fmt.Sprintf("Invalid contract ID: %+v", c.Param("contractHname")))
	}

	fname := c.Param("fname")

	var params dict.Dict
	if c.Request().Body != nil {
		if err := json.NewDecoder(c.Request().Body).Decode(&params); err != nil {
			return httperrors.BadRequest("Invalid request body")
		}
	}
	theChain := chains.AllChains().Get(chainID)
	if theChain == nil {
		return httperrors.NotFound(fmt.Sprintf("Chain not found: %s", chainID))
	}
	ret, err := webapiutil.CallView(theChain, contractHname, coretypes.Hn(fname), params)
	if err != nil {
		return httperrors.BadRequest(fmt.Sprintf("View call failed: %v", err))
	}

	return c.JSON(http.StatusOK, ret)
}

const retryOnStateInvalidatedRetry = 100 * time.Millisecond //nolint:gofumpt
const retryOnStateInvalidatedTimeout = 5 * time.Minute

func handleStateGet(c echo.Context) error {
	chainID, err := coretypes.ChainIDFromBase58(c.Param("chainID"))
	if err != nil {
		return httperrors.BadRequest(fmt.Sprintf("Invalid chain ID: %+v", c.Param("chainID")))
	}

	key, err := hex.DecodeString(c.Param("key"))
	if err != nil {
		return httperrors.BadRequest(fmt.Sprintf("cannot parse hex-encoded key: %+v", c.Param("key")))
	}

	theChain := chains.AllChains().Get(chainID)
	if theChain == nil {
		return httperrors.NotFound(fmt.Sprintf("Chain not found: %s", chainID))
	}

	var ret []byte
	err = optimism.RetryOnStateInvalidated(func() error {
		var err error
		ret, err = theChain.GetStateReader().KVStoreReader().Get(kv.Key(key))
		return err
	}, retryOnStateInvalidatedRetry, time.Now().Add(retryOnStateInvalidatedTimeout))
	if err != nil {
		reason := fmt.Sprintf("View call failed: %v", err)
		if errors.Is(err, optimism.ErrStateHasBeenInvalidated) {
			return httperrors.Conflict(reason)
		}
		return httperrors.BadRequest(reason)
	}

	return c.JSON(http.StatusOK, ret)
}
