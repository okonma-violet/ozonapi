package ozonapi

import "net/http"

type RequestParams_products_description struct {
	// Product identifier in the seller's system
	OfferId string `json:"offer_id"`

	// Product identifier
	ProductId int64 `json:"product_id"`
}

type Response_products_description struct {
	baseResponse
	Result ResponseResult_products_description `json:"result"`
}
type ResponseResult_products_description struct {
	// Description
	Description string `json:"description"`

	// Identifier
	Id int64 `json:"id"`

	// Name
	Name string `json:"name"`

	// Product identifier in the seller's system
	OfferId string `json:"offer_id"`
}

func NewRequestParams_products_description(visibility VisibilityFilter, productid int64, offerid string) *RequestParams_products_description {
	return &RequestParams_products_description{
		OfferId:   offerid,
		ProductId: productid,
	}
}

func (cl *OzonClient) GetProductsDescription(params *RequestParams_products_description) (*Response_products_description, error) {
	req, err := cl.newRequest(http.MethodPost, url_products_description, params)
	if err != nil {
		return nil, err
	}
	resp := &Response_products_description{}

	response, err := cl.doRequest(req, resp)
	if err != nil {
		return nil, err
	}
	resp.baseResponse = response.base
	//response.pasteBase(&resp.baseResponse)

	return resp, nil
}
