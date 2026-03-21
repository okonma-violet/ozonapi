package ozonapi

import (
	"net/http"
	"time"
)

// https://docs.ozon.ru/api/seller/?abt_att=1&origin_referer=docs.ozon.ru#operation/PostingAPI_GetFboPostingList

type RequestParams_posting_fbo_list struct {
	// Sorting direction
	Direction SortingDirection `json:"dir"`
	//Filter
	Filter RequestItem_posting_fbo_list_filter `json:"filter"`
	// Number of shipments in the response:
	//   - maximum is 1000,
	//   - minimum is 1.
	Limit int64 `json:"limit"`
	// Number of elements that will be skipped in the response. For example, if offset=10, the response will start with the 11th element found
	Offset int64 `json:"offset"`
	// Transliterate the return values
	Translit bool `json:"translit"`
	// Additional fields that should be added to the response
	With RequestItem_posting_fbo_list_with `json:"with"`
}

type RequestItem_posting_fbo_list_filter struct {
	// Start date of the period for which a list of shipments should be generated.
	//
	// Format: YYYYY-MM-DDTHH:MM:SSZ.
	//
	// Example: 2019-08-24T14:15:22Z
	Since string `json:"since"`
	// End date of the period for which a list of shipments should be generated.
	//
	// Format: YYYYY-MM-DDTHH:MM:SSZ.
	//
	// Example: 2019-08-24T14:15:22Z.
	To string `json:"to"`
	// awaiting_packaging — ожидает упаковки,
	// awaiting_deliver — ожидает отгрузки,
	// delivering — доставляется,
	// delivered — доставлено,
	// cancelled — отменено.
	Status string `json:"status,omitempty"`
}

type RequestItem_posting_fbo_list_with struct {
	// Аналитические данные
	AnalyticsData bool `json:"analytics_data"`
	// Финансовые данные
	FinancialData bool `json:"financial_data"`
	// Юридические данные
	LegalInfo bool `json:"legal_info"`
}

type Response_posting_fbo_list struct {
	baseResponse
	Result []ResponseItem_posting_fbo_list `json:"result"`
}

type ResponseItem_posting_fbo_list struct {
	// idk
	AdditionalData []interface{} `json:"-"` //additional_data

	// Analytics data
	AnalyticsData ResponseItem_posting_fbo_list_analyticsdata `json:"analytics_data"`

	// Cancellation reason identifier
	CancelReasonId int64 `json:"cancel_reason_id"`

	// Дата и время создания отправления
	CreatedAt time.Time `json:"created_at"`

	// Data on the product cost, discount amount, payout and commission
	FinancialData ResponseItem_posting_fbo_list_financialdata `json:"financial_data"`

	// Start date and time of shipment processing
	InProccessAt time.Time `json:"in_process_at"`

	// Юридическая информация о покупателе
	LegalInfo ResponseItem_posting_fbo_list_legalinfo `json:"legal_info"`

	// Identifier of the order to which the shipment belongs
	OrderId int64 `json:"order_id"`

	// Number of the order to which the shipment belongs
	OrderNumber string `json:"order_number"`

	// Shipment number
	PostingNumber string `json:"posting_number"`

	// List of products in the shipment
	Products []ResponseItem_posting_fbo_list_postingproduct `json:"products"`

	// Shipment status
	Status string `json:"status"`

	// Shipment substatus
	Substatus string `json:"substatus"`
}

type ResponseItem_posting_fbo_list_analyticsdata struct {
	// Delivery city
	City string `json:"city"`
	// Delivery method
	DeliveryType string `json:"delivery_type"`
	// Indication that the recipient is a legal entity:
	//
	// true — a legal entity,
	// false — a natural person
	IsLegal bool `json:"is_legal"`
	// Premium subscription availability
	IsPremium bool `json:"is_premium"`
	// картой онлайн,
	// Ozon Карта,
	// автосписание с Ozon Карты при выдаче,
	// сохранённой картой при получении,
	// Система Быстрых Платежей,
	// Ozon Рассрочка,
	// оплата на расчётный счёт,
	// SberPay,
	// предоплата на стороне внешнего продавца
	PaymentTypeGroupName string `json:"payment_type_group_name"`
	// Order shipping warehouse name
	Warehouse string `json:"warehouse_name"`
	// Warehouse identifier
	WarehouseId int64 `json:"warehouse_id"`
	// Delivery start date and time
	ClientDeliveryDateBegin time.Time `json:"client_delivery_date_begin"`
	// Delivery end date and time
	ClientDeliveryDateEnd time.Time `json:"client_delivery_date_end"`
}

type ResponseItem_posting_fbo_list_legalinfo struct {
	// Название компании
	CompanyName string `json:"company_name"`
	// инн
	INN string `json:"inn"`
	// кпп
	KPP string `json:"kpp"`
}

type ResponseItem_posting_fbo_list_postingproduct struct {
	// Коды активации для услуг и цифровых товаров
	DigitalCodes []string `json:"digital_codes"`
	// Product name
	Name string `json:"name"`
	// Product identifier in the seller's system
	OfferId string `json:"offer_id"`
	// Currency of your prices. It matches the one set in the personal account settings
	CurrencyCode string `json:"currency_code"`
	// Product price
	Price string `json:"price"`
	// Признак выкупа товара в ЕАЭС и другие страны
	IsMarketplaceBuyout bool `json:"is_marketplace_buyout"`
	// Product quantity in the shipment
	Quantity int32 `json:"quantity"`
	// Product identifier in the Ozon system, SKU
	SKU int64 `json:"sku"`
}
type ResponseItem_posting_fbo_list_financialdata struct {
	// Identifier of the cluster, where the shipment is sent from
	ClusterFrom string `json:"cluster_from"`
	// Identifier of the cluster, where the shipment is delivered
	ClusterTo string `json:"cluster_to"`
	// List of products in the shipment
	Products []ResponseItem_posting_fbo_list_financialdataproduct `json:"products"`
}

type ResponseItem_posting_fbo_list_financialdataproduct struct {
	// Actions
	Actions []string `json:"actions"`
	// Customer price
	ClientPrice string `json:"client_price"`
	// Commission amount for the product
	CommissionAmount float64 `json:"commission_amount"`
	// Commission percentage
	CommissionPercent int64 `json:"commission_percent"`
	// Code of the currency used to calculate the commissions
	CommissionsCurrencyCode string `json:"commissions_currency_code"`
	// Currency of your prices. It matches the currency set in the personal account settings
	CurrencyCode string `json:"currency_code"`
	// Price before discounts. Displayed strikethrough on the product description page
	OldPrice float64 `json:"old_price"`
	// Payment to the seller
	Payout float64 `json:"payout"`
	// Product price including discounts. This value is shown on the product description page
	Price float64 `json:"price"`
	// SKU
	ProductId int64 `json:"product_id"`
	// Product quantity in the shipment
	Quantity int64 `json:"quantity"`
	// Discount percentage
	TotalDiscountPercent float64 `json:"total_discount_percent"`
	// Discount amount
	TotalDiscountValue float64 `json:"total_discount_value"`
}

func NewRequestParams_posting_fbo_list(since, to time.Time, offset, limit int64, sortingdir SortingDirection) *RequestParams_posting_fbo_list {
	req := &RequestParams_posting_fbo_list{
		Direction: sortingdir,
		Limit:     limit,
		Offset:    offset,
		With: RequestItem_posting_fbo_list_with{
			AnalyticsData: true,
			FinancialData: true,
			LegalInfo:     true,
		},
		Filter: RequestItem_posting_fbo_list_filter{
			Since: since.UTC().Format(time_format),
			To:    to.UTC().Format(time_format),
		},
	}

	return req
}

func (cl *OzonClient) GetPostingFboList(params *RequestParams_posting_fbo_list) (*Response_posting_fbo_list, error) {
	req, err := cl.newRequest(http.MethodPost, url_posting_fbo_list, params)
	if err != nil {
		return nil, err
	}
	resp := &Response_posting_fbo_list{}

	response, err := cl.doRequest(req, resp)
	if err != nil {
		return nil, err
	}
	resp.baseResponse = response.base
	return resp, nil
}

func (cl *OzonClient) GetPostingFboListDebug(params *RequestParams_posting_fbo_list) (*Response_posting_fbo_list, *http.Request, *http.Response, []byte, error) {
	req, err := cl.newRequest(http.MethodPost, url_posting_fbo_list, params)
	if err != nil {
		return nil, nil, nil, nil, err
	}
	reqq, _ := cl.newRequest(http.MethodPost, url_posting_fbo_list, params)

	resp := &Response_posting_fbo_list{}

	response, respp, resppbody, err := cl.doRequestDebug(req, resp)
	if err != nil {
		return nil, reqq, nil, nil, err
	}
	resp.baseResponse = response.base

	return resp, reqq, respp, resppbody, nil
}
