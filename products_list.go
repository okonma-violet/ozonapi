package ozonapi

import "net/http"

type RequestParams_products_list struct {
	// Filter by product
	Filter RequestFilter_products_list `json:"filter"`

	// Identifier of the last value on the page. Leave this field blank in the first request.
	// To get the next values, specify last_id from the response of the previous request
	LastId string `json:"last_id"`

	// Number of values per page. Minimum is 1, maximum is 1000
	Limit int64 `json:"limit"`
}
type RequestFilter_products_list struct {
	// Filter by the offer_id parameter. You can pass a list of values in this parameter
	OfferId []string `json:"offer_id"`

	// Filter by the product_id parameter. You can pass a list of values in this parameter
	ProductId []string `json:"product_id"`

	// Filter by product visibility
	Visibility VisibilityFilter `json:"visibility"`
}

type Response_products_list struct {
	baseResponse
	Result ResponseResult_products_list `json:"result"`
}
type ResponseResult_products_list struct {
	// Products list
	Items []ResponseItem_products_list `json:"items"`

	// Identifier of the last value on the page.
	// To get the next values, specify the recieved value in the next request in the last_id parameter
	LastId string `json:"last_id"`

	// Total number of products
	Total int32 `json:"total"`
}
type ResponseItem_products_list struct {
	// Product identifier in the seller's system
	OfferId string `json:"offer_id"`

	// Product ID
	ProductId int64 `json:"product_id"`

	HasFboStocks bool `json:"has_fbo_stocks"`
	HasFbsStocks bool `json:"has_fbs_stocks"`
	Archieved    bool `json:"archieved"`
	IsDiscounted bool `json:"is_discounted"`

	// хз шо это
	Quants []ResponseItem_products_list_quant `json:"quants"`
}

type ResponseItem_products_list_quant struct {
	// Идентификатор эконом-товара
	QuantCode string `json:"quant_code"`

	// Размер кванта
	QuantSize int64 `json:"quant_size"`
}

func NewRequestParams_products_list(visibility VisibilityFilter, lastid string, limit int64) *RequestParams_products_list {
	if limit <= 0 || limit > RequestItems_cap_products_list {
		limit = RequestItems_cap_products_list
	}
	return &RequestParams_products_list{
		Filter: RequestFilter_products_list{
			Visibility: visibility,
		},
		LastId: lastid,
		Limit:  limit,
	}
}

func (cl *OzonClient) GetProductsList(params *RequestParams_products_list) (*Response_products_list, error) {
	req, err := cl.newRequest(http.MethodPost, url_products_list, params)
	if err != nil {
		return nil, err
	}
	resp := &Response_products_list{}

	response, err := cl.doRequest(req, resp)
	if err != nil {
		return nil, err
	}
	resp.baseResponse = response.base
	//response.pasteBase(&resp.baseResponse)

	return resp, nil
}
