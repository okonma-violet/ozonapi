package ozonapi

import (
	"context"
	"net/http"
)

type RequestParams_stocks_update struct {
	// Stock details
	Stocks []RequestItem_stocks_update `json:"stocks"`
}

type RequestItem_stocks_update struct {
	// Product identifier in the seller's system
	OfferId string `json:"offer_id"`

	// Product identifier
	ProductId int64 `json:"product_id"`

	// Quantity
	Stock int64 `json:"stock"`

	// Warehouse identifier derived from the /v1/warehouse/list method
	WarehouseId int64 `json:"warehouse_id"`
}

type Response_stocks_update struct {
	baseResponse
	Result []ResponseItem_stocks_update `json:"result"`
}

type ResponseItem_stocks_update struct {
	// An array of errors that occurred while processing the request
	Errors []ResponseItem_stocks_update_error `json:"errors"`

	// Product identifier in the seller's system
	Offerid string `json:"offer_id"`

	// Product identifier
	ProductId int64 `json:"product_id"`

	// If the request was completed successfully and the stocks are updated — true
	Updated bool `json:"updated"`

	// Warehouse identifier derived from the /v1/warehouse/list method
	WarehouseId int64 `json:"warehouse_id"`
}

type ResponseItem_stocks_update_error struct {
	Code    ErrCode_stocks_update `json:"code"`
	Message string                `json:"message"`
}

type ErrCode_stocks_update string

const ErrCode_nochanges ErrCode_stocks_update = "SKU STOCK NOT CHANGED"

func (cl *OzonClient) ApplyRateLimit_stocks_update(ctx context.Context) *OzonClient {
	cl.applyRateLimit(ctx, request_ratelimit_stocks_update, default_burst(request_ratelimit_stocks_update))
	return cl
}

func NewRequestParams_stocks_update() *RequestParams_stocks_update {
	return &RequestParams_stocks_update{Stocks: make([]RequestItem_stocks_update, 0, RequestItems_cap_stocks_update)}
}

// return false if requestitems cap reached (and shit did not add)
func (rp *RequestParams_stocks_update) Add(offerid string, productid, storageid, stock int64) bool {
	if len(rp.Stocks) < RequestItems_cap_stocks_update {
		rp.Stocks = append(rp.Stocks, RequestItem_stocks_update{OfferId: offerid, ProductId: productid, Stock: stock, WarehouseId: storageid})
		return true
	}
	return false
}

func (cl *OzonClient) UpdateStocks(params *RequestParams_stocks_update) (*Response_stocks_update, error) {
	req, err := cl.newRequest(http.MethodPost, url_stocks_update, params)
	if err != nil {
		return nil, err
	}
	resp := &Response_stocks_update{}

	response, err := cl.doRequest(req, resp)
	if err != nil {
		return nil, err
	}
	resp.baseResponse = response.base
	//response.pasteBase(&resp.baseResponse)

	return resp, nil
}

func (cl *OzonClient) UpdateStocksDebug(params *RequestParams_stocks_update) (*Response_stocks_update, *http.Request, *http.Response, []byte, error) {
	req, err := cl.newRequest(http.MethodPost, url_stocks_update, params)
	if err != nil {
		return nil, nil, nil, nil, err
	}
	reqq, _ := cl.newRequest(http.MethodPost, url_stocks_update, params)

	resp := &Response_stocks_update{}

	response, respp, resppbody, err := cl.doRequestDebug(req, resp)
	if err != nil {
		return nil, reqq, nil, nil, err
	}
	resp.baseResponse = response.base
	//response.pasteBase(&resp.baseResponse)

	return resp, reqq, respp, resppbody, nil
}
