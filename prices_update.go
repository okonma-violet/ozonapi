package ozonapi

import (
	"net/http"
)

type RequestParams_prices_update struct {
	// Product prices details
	Prices []RequestItem_prices_update `json:"prices"`
}

type RequestItem_prices_update struct {
	// Attribute for enabling and disabling promos auto-application
	AutoActionEnabled SwitchingAttribute `json:"auto_action_enabled"`

	AutoAddToOzonActionsListEnabled SwitchingAttribute `json:"auto_add_to_ozon_actions_list_enabled"`

	// Currency of your prices. The passed value must be the same as the one set in the personal account settings.
	// By default, the passed value is RUB, Russian ruble
	CurrencyCode CurrencyCode `json:"currency_code"`

	// Minimum product price with all promotions applied
	MinPrice string `json:"min_price"`

	// Product identifier in the seller's system
	OfferId string `json:"offer_id"`

	// Price before discounts. Displayed strikethrough on the product description page.
	// Specified in rubles.
	// The fractional part is separated by decimal point,
	// up to two digits after the decimal point.
	//
	// If there are no discounts on the product, pass 0 to this field and specify the correct price in the price field
	OldPrice string `json:"old_price"`

	// Product price including discounts. This value is displayed on the product description page.
	//
	// If the old_price parameter value is greater than 0,
	// there should be a certain difference between price and old_price.
	// It depends on the price value
	//
	// < 400 - min diff. 20 rubles
	//
	// 400-10,000 - min diff. 5%
	//
	// > 10,000 - min diff. 500 rubles
	Price string `json:"price"`

	// Attribute for enabling and disabling pricing strategies auto-application
	//
	// If you've previously enabled automatic application of pricing strategies and don't want to disable it, pass UNKNOWN in the next requests.
	//
	// If you pass `ENABLED` in this parameter, pass `strategy_id` in the `/v1/pricing-strategy/products/add` method request.
	//
	// If you pass `DISABLED` in this parameter, the product is removed from the strategy
	PriceStrategyEnabled SwitchingAttribute `json:"price_strategy_enabled"`

	// Product identifier
	ProductId int64 `json:"product_id"`
}

type Response_prices_update struct {
	baseResponse
	Result []ResponseItem_prices_update `json:"result"`
}
type ResponseItem_prices_update struct {
	// An array of errors that occurred while processing the request
	Errors []ResponseItem_prices_update_error `json:"errors"`
	// Product identifier in the seller's system
	OfferId string `json:"offer_id"`
	// Product ID
	ProductId int64 `json:"product_id"`
	// If the product details have been successfully updated — true
	Updated bool `json:"updated"`
}

type ResponseItem_prices_update_error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func NewRequestParams_prices_update() *RequestParams_prices_update {
	return &RequestParams_prices_update{Prices: make([]RequestItem_prices_update, 0, RequestItems_cap_prices_update)}
}

func (rp *RequestParams_prices_update) CapReached() bool {
	return len(rp.Prices) >= RequestItems_cap_prices_update
}

// return false if request items cap reached (and shit did not add)
func (rp *RequestParams_prices_update) Add(offerid string, productid int64, nosaleprice, price string, disableactions bool) bool {
	var actswitch SwitchingAttribute
	if disableactions {
		actswitch = SwitchingAttributeDisabled
	} else {
		actswitch = SwitchingAttributeUnknown
	}
	if len(rp.Prices) < RequestItems_cap_prices_update {
		rp.Prices = append(rp.Prices, RequestItem_prices_update{OfferId: offerid, ProductId: productid, OldPrice: nosaleprice, Price: price, AutoActionEnabled: actswitch, AutoAddToOzonActionsListEnabled: actswitch, CurrencyCode: CurrencyCode_RUB, PriceStrategyEnabled: SwitchingAttributeUnknown})
		return true
	}
	return false
}
func (cl *OzonClient) UpdatePrices(params *RequestParams_prices_update) (*Response_prices_update, error) {
	req, err := cl.newRequest(http.MethodPost, url_prices_update, params)
	if err != nil {
		return nil, err
	}
	resp := &Response_prices_update{}

	response, err := cl.doRequest(req, resp)
	if err != nil {
		return nil, err
	}
	resp.baseResponse = response.base
	//response.pasteBase(&resp.baseResponse)

	return resp, nil
}

// TODO: UpdatePricesDebug->UpdatePricesVerbose, вынести json-энкодинг из newRequest, делать reqClone и в его body пихать io.nopcloser, либо просто отдельно reqbody возвращать
func (cl *OzonClient) UpdatePricesDebug(params *RequestParams_prices_update) (*Response_prices_update, *http.Request, *http.Response, []byte, error) {
	req, err := cl.newRequest(http.MethodPost, url_prices_update, params)
	if err != nil {
		return nil, nil, nil, nil, err
	}
	reqq, _ := cl.newRequest(http.MethodPost, url_prices_update, params)

	resp := &Response_prices_update{}

	response, respp, resppbody, err := cl.doRequestDebug(req, resp)
	if err != nil {
		return nil, reqq, nil, nil, err
	}
	resp.baseResponse = response.base
	//response.pasteBase(&resp.baseResponse)

	return resp, reqq, respp, resppbody, nil
}
