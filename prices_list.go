package ozonapi

import (
	"net/http"
	"strconv"
	"time"
)

type RequestParams_prices_list struct {
	// Filter by product
	Filter RequestFilter_prices_list `json:"filter"`

	// Указатель для выборки следующих данных
	Cursor string `json:"cursor"`

	// Number of values per page. Minimum is 1, maximum is 1000
	Limit int32 `json:"limit"`
}
type RequestFilter_prices_list struct {
	// Filter by the offer_id parameter. You can pass a list of values in this parameter
	OfferId []string `json:"offer_id"`

	// Filter by the product_id parameter. You can pass a list of values in this parameter
	ProductId []string `json:"product_id"`

	// Filter by product visibility
	Visibility VisibilityFilter `json:"visibility"`
}

type Response_prices_list struct {
	baseResponse

	// Products list
	Items []ResponseItem_prices_list `json:"items"`

	// Указатель для выборки следующих данных
	Cursor string `json:"cursor"`

	// Total number of products
	Total int32 `json:"total"`
}

type ResponseItem_prices_list struct {
	// Maximum acquiring fee
	Acquiring float64 `json:"acquiring"`

	// Commissions information
	Commissions ResponseItem_prices_list_commissions `json:"commissions"`

	// Promotions information
	MarketingActions ResponseItem_prices_list_marketingactions `json:"marketing_actions"`

	// Seller product identifier
	OfferId string `json:"offer_id"`

	// Product price
	Price ResponseItem_prices_list_price `json:"price"`

	// Product price indexes
	PriceIndexes ResponseItem_prices_list_priceindexes `json:"prices_indexes"`

	// Product identifier
	ProductId int64 `json:"product_id"`

	// Product volume weight
	VolumeWeight float64 `json:"volume_weight"`
}

type ResponseItem_prices_list_commissions struct {
	// Last mile (FBO)
	FBOLastMile float64 `json:"fbo_deliv_to_customer_amount"`

	// Pipeline to (FBO)
	FBOPipelineTo float64 `json:"fbo_direct_flow_trans_max_amount"`

	// Pipeline from (FBO)
	FBOPipelineFrom float64 `json:"fbo_direct_flow_trans_min_amount"`

	// Return and cancellation fees (FBO)
	FBOReturnCancellationFee float64 `json:"fbo_return_flow_amount"`

	// Last mile (FBS)
	FBSLastMile float64 `json:"fbs_deliv_to_customer_amount"`

	// Pipeline to (FBS)
	FBSPipelineTo float64 `json:"fbs_direct_flow_trans_max_amount"`

	// Pipeline from (FBS)
	FBSPipelineFrom float64 `json:"fbs_direct_flow_trans_min_amount"`

	// First mile min (FBS)
	FBSFirstMileMin float64 `json:"fbs_first_mile_min_amount"`

	// First mile max (FBS)
	FBSFirstMileMax float64 `json:"fbs_first_mile_max_amount"`

	// Return and cancellation fees, shipment processing (FBS)
	FBSReturnCancellationProcessingFee float64 `json:"fbs_return_flow_amount"`

	// Sales commission percentage (FBO)
	SalesCommissionFBORate float64 `json:"sales_percent_fbo"`

	// Sales commission percentage (FBS)
	SalesCommissionFBSRate float64 `json:"sales_percent_fbs"`
}

type ResponseItem_prices_list_marketingactions struct {
	// Seller's promotions. The parameters date_from, date_to, discount_value and title are specified for each seller's promotion
	Actions []ResponseItem_prices_list_marketingactionsitem `json:"actions"`

	// Current period start date and time for all current promotions
	CurrentPeriodFrom time.Time `json:"current_period_from"`

	// Current period end date and time for all current promotions
	CurrentPeriodTo time.Time `json:"current_period_to"`

	// If a promotion can be applied to the product at the expense of Ozon, this field is set to true
	OzonActionsExist bool `json:"ozon_actions_exist"`
}
type ResponseItem_prices_list_marketingactionsitem struct {
	// Date and time when the seller's promotion starts
	DateFrom time.Time `json:"date_from"`

	// Date and time when the seller's promotion ends
	DateTo time.Time `json:"date_to"`

	// Discount on the seller's promotion
	Value float64 `json:"value"`

	// Promotion name
	Title string `json:"title"`
}
type ResponseItem_prices_list_price struct {
	// If promos auto-application is enabled, the value is true
	AutoActionEnabled bool `json:"auto_action_enabled"`

	// Currency of your prices. It matches the currency set in the personal account settings
	CurrencyCode string `json:"currency_code"`

	// Product price including all promotion discounts. This value will be indicated on the Ozon storefront
	MarketingPrice float64 `json:"marketing_price"`

	// Product price with seller's promotions applied
	MarketingSellerPrice float64 `json:"marketing_seller_price"`

	// Minimum product price with all promotions applied
	MinPrice float64 `json:"min_price"`

	// Price before discounts. Displayed strikethrough on the product description page
	OldPrice float64 `json:"old_price"`

	// Product price including discounts. This value is shown on the product description page
	Price float64 `json:"price"`

	// Retailer price
	RetailPrice float64 `json:"retail_price"`

	// Product VAT rate
	VAT float64 `json:"vat"`
}
type ResponseItem_prices_list_priceindexes struct {
	// Default: "WITHOUT_INDEX"
	// Итоговый индекс цены товара:
	// WITHOUT_INDEX — нет индекса,
	// GREEN — выгодный,
	// YELLOW — умеренный,
	// RED — невыгодный.
	ColorIndex string `json:"color_index"`

	// Competitors' product price on other marketplaces
	ExternalIndexData ResponseItem_prices_list_priceindext `json:"external_index_data"`

	// Competitors' product price on Ozon
	OzonIndexData ResponseItem_prices_list_priceindext `json:"ozon_index_data"`

	// Price of your product on other marketplaces
	SelfMarketplaceIndexData ResponseItem_prices_list_priceindext `json:"self_marketplaces_index_data"`
}
type ResponseItem_prices_list_priceindext struct {
	// Минимальная цена товара у конкурентов на другой площадке/на Ozon или вашего товара на других площадках
	MinimalPrice string `json:"minimal_price"`

	// Валюта цены
	MinimalPriceCurrency string `json:"minimal_price_currency"`

	// Значение индекса цены
	PriceIndexValue float64 `json:"price_index_value"`
}

func NewRequestParams_prices_list(visibility VisibilityFilter, offerids []string, productids []string, cursor string, limit int32) *RequestParams_prices_list {
	if limit <= 0 || limit > RequestItems_cap_prices_list {
		panic("incorrect limit (<0 or >ozon's limit) : " + strconv.Itoa(int(limit)))
	}

	return &RequestParams_prices_list{
		Filter: RequestFilter_prices_list{
			Visibility: visibility,
			ProductId:  productids,
			OfferId:    offerids,
		},
		Cursor: cursor,
		Limit:  limit,
	}
}

// return false if requestitems cap reached (and shit did not add)
func (rp *RequestParams_prices_list) Add(offerid string, productid string) bool {
	if len(rp.Filter.OfferId)+len(rp.Filter.ProductId) < RequestItems_cap_prices_list {
		if offerid != "" {
			rp.Filter.OfferId = append(rp.Filter.OfferId, offerid)
		}
		if productid != "" {
			rp.Filter.ProductId = append(rp.Filter.ProductId, productid)
		}
		return true
	}
	return false
}

func (cl *OzonClient) GetPricesList(params *RequestParams_prices_list) (*Response_prices_list, error) {
	req, err := cl.newRequest(http.MethodPost, url_prices_list, params)
	if err != nil {
		return nil, err
	}
	resp := &Response_prices_list{}

	response, err := cl.doRequest(req, resp)
	if err != nil {
		return nil, err
	}
	resp.baseResponse = response.base
	//response.pasteBase(&resp.baseResponse)

	return resp, nil
}
