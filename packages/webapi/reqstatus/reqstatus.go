package reqstatus

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/iotaledger/wasp/packages/chain"
	"github.com/iotaledger/wasp/packages/chains"
	"github.com/iotaledger/wasp/packages/isc"
	"github.com/iotaledger/wasp/packages/isc/coreutil"
	"github.com/iotaledger/wasp/packages/kv/optimism"
	"github.com/iotaledger/wasp/packages/util/panicutil"
	"github.com/iotaledger/wasp/packages/webapi/httperrors"
	"github.com/iotaledger/wasp/packages/webapi/model"
	"github.com/iotaledger/wasp/packages/webapi/routes"
	"github.com/labstack/echo/v4"
	"github.com/pangpanglabs/echoswagger/v2"
	"golang.org/x/xerrors"
)

type reqstatusWebAPI struct {
	getChain func(chainID *isc.ChainID) chain.ChainRequests
}

// TODO  add examples for receipt json
func AddEndpoints(server echoswagger.ApiRouter, getChain chains.ChainProvider) {
	r := &reqstatusWebAPI{func(chainID *isc.ChainID) chain.ChainRequests {
		return getChain(chainID)
	}}

	server.GET(routes.RequestReceipt(":chainID", ":reqID"), r.handleRequestReceipt).
		SetSummary("Get the processing status of a given request in the node").
		AddParamPath("", "chainID", "ChainID (bech32)").
		AddParamPath("", "reqID", "Request ID").
		AddResponse(http.StatusOK, "Request Receipt", model.RequestReceiptResponse{}, nil)

	server.GET(routes.WaitRequestProcessed(":chainID", ":reqID"), r.handleWaitRequestProcessed).
		SetSummary("Wait until the given request has been processed by the node").
		AddParamPath("", "chainID", "ChainID (bech32)").
		AddParamPath("", "reqID", "Request ID").
		AddParamBody(model.WaitRequestProcessedParams{}, "Params", "Optional parameters", false).
		AddResponse(http.StatusOK, "Request Receipt", model.RequestReceiptResponse{}, nil)
}

func (r *reqstatusWebAPI) handleRequestReceipt(c echo.Context) error {
	ch, reqID, err := r.parseParams(c)
	if err != nil {
		return err
	}

	receiptResponse, err := getISCReceipt(ch, reqID)
	if err != nil {
		return httperrors.ServerError(err.Error())
	}
	return c.JSON(http.StatusOK, receiptResponse)
}

const WaitRequestProcessedDefaultTimeout = 30 * time.Second

func (r *reqstatusWebAPI) handleWaitRequestProcessed(c echo.Context) error {
	ch, reqID, err := r.parseParams(c)
	if err != nil {
		return err
	}

	req := model.WaitRequestProcessedParams{
		Timeout: WaitRequestProcessedDefaultTimeout,
	}
	if c.Request().Header.Get("Content-Type") == "application/json" {
		if err := c.Bind(&req); err != nil {
			return httperrors.BadRequest("Invalid request body")
		}
	}

	tryGetReceipt := func() (bool, error) {
		receiptResponse, err := getISCReceipt(ch, reqID)
		if err != nil {
			return receiptResponse != nil, httperrors.ServerError(err.Error())
		}
		if receiptResponse != nil {
			return true, c.JSON(http.StatusOK, receiptResponse)
		}
		return false, nil
	}

	found, ret := tryGetReceipt()
	if found {
		return ret
	}

	// subscribe to event
	requestProcessed := make(chan bool)
	attachID := ch.AttachToRequestProcessed(func(rid isc.RequestID) {
		if rid == reqID {
			requestProcessed <- true
		}
	})
	defer ch.DetachFromRequestProcessed(attachID)

	select {
	case <-requestProcessed:
		found, ret := tryGetReceipt()
		if found {
			return ret
		}
		return httperrors.ServerError("Unexpected error, receipt not found after request was processed")
	case <-time.After(req.Timeout):
		// check again, in case event was triggered just before we subscribed
		found, ret := tryGetReceipt()
		if found {
			return ret
		}
		return httperrors.Timeout("Timeout while waiting for request to be processed")
	}
}

func (r *reqstatusWebAPI) parseParams(c echo.Context) (chain.ChainRequests, isc.RequestID, error) {
	chainID, err := isc.ChainIDFromString(c.Param("chainID"))
	if err != nil {
		return nil, isc.RequestID{}, httperrors.BadRequest(fmt.Sprintf("Invalid Chain ID %+v: %s", c.Param("chainID"), err.Error()))
	}
	theChain := r.getChain(chainID)
	if theChain == nil {
		return nil, isc.RequestID{}, httperrors.NotFound(fmt.Sprintf("Chain not found: %s", chainID.String()))
	}
	reqID, err := isc.RequestIDFromString(c.Param("reqID"))
	if err != nil {
		return nil, isc.RequestID{}, httperrors.BadRequest(fmt.Sprintf("Invalid request id %+v: %s", c.Param("reqID"), err.Error()))
	}
	return theChain, reqID, nil
}

func doGetISCReceipt(ch chain.ChainRequests, reqID isc.RequestID) (*model.RequestReceiptResponse, error) {
	receipt, err := ch.GetRequestReceipt(reqID)
	if err != nil {
		return nil, xerrors.Errorf("error getting request receipt: %s", err)
	}
	if receipt == nil {
		return nil, nil
	}

	resolvedError, err := ch.ResolveError(receipt.Error)
	if err != nil {
		return nil, xerrors.Errorf("error resolving the receipt error: %s", err)
	}
	iscReceipt := receipt.ToISCReceipt(resolvedError)

	receiptJSON, err := json.Marshal(iscReceipt)
	if err != nil {
		return nil, xerrors.Errorf("error marshaling receipt into JSON: %s", err)
	}
	return &model.RequestReceiptResponse{
		Receipt: string(receiptJSON),
	}, nil
}

func getISCReceipt(ch chain.ChainRequests, reqID isc.RequestID) (ret *model.RequestReceiptResponse, err error) {
	err = optimism.RetryOnStateInvalidated(func() (err error) {
		panicCatchErr := panicutil.CatchPanicReturnError(func() {
			ret, err = doGetISCReceipt(ch, reqID)
		}, coreutil.ErrorStateInvalidated)
		if err != nil {
			return err
		}
		return panicCatchErr
	})
	return ret, err
}
