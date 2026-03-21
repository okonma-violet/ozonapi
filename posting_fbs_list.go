package ozonapi

import (
	"net/http"
	"time"
)

type RequestParams_posting_fbs_list struct {
	// Sorting direction
	Direction SortingDirection `json:"dir"`

	//Filter
	Filter RequestItem_posting_fbs_list_filter `json:"filter"`

	// Number of shipments in the response:
	//   - maximum is 1000,
	//   - minimum is 1.
	Limit int64 `json:"limit"`

	// Number of elements that will be skipped in the response. For example, if offset=10, the response will start with the 11th element found
	Offset int64 `json:"offset"`

	// Additional fields that should be added to the response
	With RequestItem_posting_fbs_list_with `json:"with"`
}

type RequestItem_posting_fbs_list_filter struct {
	// Delivery method identifier
	DeliveryMethodId []int64 `json:"delivery_method_id,omitempty"`

	// Optional, but "Since" and "To" are still needed
	LastChangedStatusDate RequestItem_posting_fbs_list_filter_changed_status

	// // Filter for shipments delivered from partner warehouse (FBP). You can pass one of the following values:
	// //
	// // Default value is all.
	// //
	// // The FBP scheme is available only for sellers from China
	// FBPFilter FBPFilter `json:"fbpFilter"`

	// Order identifier
	OrderId int64 `json:"order_id,omitempty"`

	// Delivery service identifier
	ProviderId []int64 `json:"provider_id,omitempty"`

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

	// Shipment status
	Status string `json:"status,omitempty"`

	// Warehouse identifier
	WarehouseId []int64 `json:"warehouse_id,omitempty"`
}

type RequestItem_posting_fbs_list_filter_changed_status struct {
	From string `json:"from"`
	To   string `json:"to"`
}

type RequestItem_posting_fbs_list_with struct {
	// Add analytics data to the response
	AnalyticsData bool `json:"analytics_data"`

	// Add the shipment barcodes to the response
	Barcodes bool `json:"barcodes"`

	// Add financial data to the response
	FinancialData bool `json:"financial_data"`

	// Transliterate the return values
	Translit bool `json:"translit"`
}

type Response_posting_fbs_list struct {
	baseResponse
	Result ResponseResult_posting_fbs_list `json:"result"`
}
type ResponseResult_posting_fbs_list struct {
	// Indicates that the response returned not the entire array of shipments:
	//   - true — it is necessary to make a new request with a different offset value to get information on the remaining shipments;
	//   - false — the entire array of shipments for the filter specified in the request was returned in the response
	HasNext bool `json:"has_next"`

	// Shipment details
	Postings []ResponseItem_posting_fbs_list `json:"postings"`
}
type ResponseItem_posting_fbs_list struct {
	// Recipient details
	Addressee ResponseItem_posting_fbs_list_addressee `json:"addressee"`

	// Analytics data
	AnalyticsData ResponseItem_posting_fbs_list_analyticsdata `json:"analytics_data"`

	// Shipment barcodes
	Barcodes ResponseItem_posting_fbs_list_barcode `json:"barcodes"`

	// Cancellation details
	Cancellation ResponseItem_posting_fbs_list_cancellation `json:"cancellation"`

	// Customer details
	Customer ResponseItem_posting_fbs_list_customer `json:"customer"`

	// Date when the shipment was transferred for delivery
	DeliveringDate time.Time `json:"delivering_date"`

	// Delivery method
	DeliveryMethod ResponseItem_posting_fbs_list_deliverymethod `json:"delivery_method"`

	// Data on the product cost, discount amount, payout and commission
	FinancialData ResponseItem_posting_fbs_list_financialdata `json:"financial_data"`

	// Start date and time of shipment processing
	InProccessAt time.Time `json:"in_process_at"`

	// If Ozon Express fast delivery was used — `true`
	IsExpress bool `json:"is_express"`

	// Indication that there is a multi-box product in the shipment
	// and you need to pass the number of boxes for it
	IsMultibox bool `json:"is_multibox"`

	// Number of boxes in which the product is packed
	MultiBoxQuantity int32 `json:"multi_box_qty"`

	// Identifier of the order to which the shipment belongs
	OrderId int64 `json:"order_id"`

	// Number of the order to which the shipment belongs
	OrderNumber string `json:"order_number"`

	// Number of the parent shipment which split resulted in the current shipment
	ParentPostingNumber string `json:"parent_posting_number"`

	// Shipment number
	PostingNumber string `json:"posting_number"`

	// List of products in the shipment
	Products []ResponseItem_posting_fbs_list_postingproduct `json:"products"`

	// The parameter is only relevant for bulky products
	// with a delivery by a third-party or integrated service
	PRROption PRROptionStatus `json:"prr_option"`

	// Array of Ozon Product IDs (SKU) for which you need to pass the
	// customs cargo declaration (CCD) number, the manufacturing country,
	// product batch registration number, or "Chestny ZNAK" labeling to
	// change the shipment status to the next one
	Requirements ResponseItem_posting_fbs_list_requirements `json:"requirements"`

	// Дата и время, до которой необходимо собрать отправление. Показываем рекомендованное время отгрузки.
	// По истечении этого времени начнёт применяться новый тариф, информацию о нём уточняйте в поле tariffication
	ShipmentDate time.Time `json:"shipment_date"`

	// Дата и время отгрузки без просрочки.
	ShipmentDateWithoutDelay time.Time `json:"shipment_date_without_delay"`

	// Shipment status
	Status string `json:"status"`

	// Shipment substatus
	Substatus string `json:"substatus"`

	// Type of integration with the delivery service
	TPLIntegrationType string `json:"tpl_integration_type"`

	// Shipment tracking number
	TrackingNumber string `json:"tracking_number"`
}
type ResponseItem_posting_fbs_list_addressee struct {
	// Recipient name
	Name string `json:"name"`

	// Recipient phone number.
	//
	// Returns an empty string
	Phone string `json:"phone"`
}
type ResponseItem_posting_fbs_list_analyticsdata struct {
	// Delivery city
	City string `json:"city"`

	// Delivery start date and time
	DeliveryDateBegin time.Time `json:"delivery_date_begin"`

	// Delivery end date and time
	DeliveryDateEnd time.Time `json:"delivery_date_end"`

	// Delivery method
	DeliveryType string `json:"delivery_type"`

	// Indication that the recipient is a legal entity:
	//
	// true — a legal entity,
	// false — a natural person
	IsLegal bool `json:"is_legal"`

	// Premium subscription availability
	IsPremium bool `json:"is_premium"`

	// Payment method
	PaymentTypeGroupName string `json:"payment_type_group_name"`

	// Delivery region
	Region string `json:"region"`

	// Delivery service
	TPLProvider string `json:"tpl_provider"`

	// Delivery service identifier
	TPLProviderId int64 `json:"tpl_provider_id"`

	// Order shipping warehouse name
	Warehouse string `json:"warehouse"`

	// Warehouse identifier
	WarehouseId int64 `json:"warehouse_id"`
}

type ResponseItem_posting_fbs_list_barcode struct {
	// Lower barcode on the shipment label
	LowerBarcode string `json:"lower_barcode"`

	// Upper barcode on the shipment label
	UpperBarcode string `json:"upper_barcode"`
}

type ResponseItem_posting_fbs_list_cancellation struct {
	// If the cancellation affects the seller's rating—true
	AffectCancellationRating bool `json:"affect_cancellation_rating"`

	// Cancellation reason
	CancelReason string `json:"cancel_reason"`

	// Cancellation reason identifier
	CancelReasonId int64 `json:"cancel_reason_id"`

	// Shipment cancellation initiator:
	//
	// Клиент—client,
	// Ozon,
	// Продавец—seller
	CancellationInitiator string `json:"cancellation_initiator"`

	// Cancellation type:
	//
	// client—canceled by client.
	// ozon—canceled by Ozon.
	// seller—canceled by seller
	CancellationType string `json:"cancellation_type"`

	// If the cancellation occurred after the shipment had been packaged—true
	CancelledAfterShip bool `json:"cancelled_after_ship"`
}

type ResponseItem_posting_fbs_list_customer struct {
	// Delivery address details
	Address ResponseItem_posting_fbs_list_customeraddress `json:"address"`

	// Customer e-mail
	CustomerEmail string `json:"customer_email"`

	// Customer identifier
	CustomerId int64 `json:"customer_id"`

	// Customer name
	Name string `json:"name"`

	// Customer phone number.
	//
	// Returns an empty string
	Phone string `json:"phone"`
}

type ResponseItem_posting_fbs_list_customeraddress struct {
	// Address in text format
	AddressTail string `json:"address_tail"`

	// Delivery city
	City string `json:"city"`

	// Comment on the order
	Comment string `json:"comment"`

	// Delivery country
	Country string `json:"country"`

	// Delivery area
	District string `json:"district"`

	// Latitude
	Latitude float64 `json:"latitude"`

	// Longitude
	Longitude float64 `json:"longitude"`

	// 3PL pick-up point code
	ProviderPVZCode string `json:"provider_pvz_code"`

	// Pick-up point code
	PVZCode int64 `json:"pvz_code"`

	// Delivery region
	Region string `json:"region"`

	// Recipient postal code
	ZIPCode string `json:"zip_code"`
}

type ResponseItem_posting_fbs_list_deliverymethod struct {
	// Delivery method identifier
	Id int64 `json:"id"`

	// Delivery method name
	Name string `json:"name"`

	// Delivery service
	TPLProvider string `json:"tpl_provider"`

	// Delivery service identifier
	TPLProviderId int64 `json:"tpl_provider_id"`

	// Warehouse name
	Warehouse string `json:"warehouse"`

	// Warehouse identifier
	WarehouseId int64 `json:"warehouse_id"`
}

type ResponseItem_posting_fbs_list_financialdata struct {
	// Identifier of the cluster, where the shipment is sent from
	ClusterFrom string `json:"cluster_from"`

	// Identifier of the cluster, where the shipment is delivered
	ClusterTo string `json:"cluster_to"`

	// Services
	PostingServices ResponseItem_posting_fbs_list_marketplaceservices `json:"posting_services"`

	// List of products in the shipment
	Products []ResponseItem_posting_fbs_list_financialdataproduct `json:"products"`
}

type ResponseItem_posting_fbs_list_marketplaceservices struct {
	// Last mile
	DeliveryToCustomer float64 `json:"marketplace_service_item_deliv_to_customer"`

	// Pipeline
	DirectFlowTrans float64 `json:"marketplace_service_item_direct_flow_trans"`

	// Shipment processing in the fulfilment warehouse (FF)
	DropoffFF float64 `json:"marketplace_service_item_item_dropoff_ff"`

	// Shipment processing at the pick up point
	DropoffPVZ float64 `json:"marketplace_service_item_dropoff_pvz"`

	// Shipment processing at the sorting center
	DropoffSC float64 `json:"marketplace_service_item_dropoff_sc"`

	// Order packaging
	Fulfillment float64 `json:"marketplace_service_item_fulfillment"`

	// Transport arrival to the seller's address for shipments pick-up (Pick-up)
	Pickup float64 `json:"marketplace_service_item_pickup"`

	// Return processing
	ReturnAfterDeliveryToCustomer float64 `json:"marketplace_service_item_return_after_deliv_to_customer"`

	// Reverse pipeline
	ReturnFlowTrans float64 `json:"marketplace_service_item_return_flow_trans"`

	// Cancellations processing
	ReturnNotDeliveryToCustomer float64 `json:"marketplace_service_item_return_not_deliv_to_customer"`

	// Non-purchase processing
	ReturnPartGoodsCustomer float64 `json:"marketplace_service_item_return_part_goods_customer"`
}

type ResponseItem_posting_fbs_list_requirements struct {
	// Array of Ozon Product IDs (SKU) for which you need to pass the customs cargo declaration (CCD) numbers.
	//
	// To pack the shipment, pass the CCD number for all listed SKUs.
	// If you do not have a CCD number, pass the value `is_gtd_absent` = true
	// via the `/v3/posting/fbs/ship/package` or `/v3/posting/fbs/ship` method
	ProductsRequiringGTD []int64 `json:"products_requiring_gtd"`

	// Array of Ozon Product IDs (SKU) for which
	// you need to pass the manufacturing country.
	//
	// To pack the shipment, pass the manufacturing
	// country value for all listed SKUs via the `/v2/posting/fbs/product/country/set` method
	ProductsRequiringCountry []int64 `json:"products_requiring_country"`

	// Array of Ozon Product IDs (SKU) for which you need to pass the "Chestny ZNAK" labeling
	ProductsRequiringMandatoryMark []int64 `json:"products_requiring_mandatory_mark"`

	// Array of Ozon Product IDs (SKU) for which you need to pass a product batch registration number
	ProductsRequiringRNPT []int64 `json:"products_requiring_rnpt"`
}

type ResponseItem_posting_fbs_list_postingproduct struct {
	// Mandatory product labeling
	MandatoryMark []string `json:"mandatory_mark"`

	// Product name
	Name string `json:"name"`

	// Product identifier in the seller's system
	OfferId string `json:"offer_id"`

	// Currency of your prices. It matches the one set in the personal account settings
	CurrencyCode string `json:"currency_code"`

	// Product price
	Price string `json:"price"`

	// Product quantity in the shipment
	Quantity int32 `json:"quantity"`

	// Product identifier in the Ozon system, SKU
	SKU int64 `json:"sku"`
}

type ResponseItem_posting_fbs_list_financialdataproduct struct {
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

	// Services
	ItemServices ResponseItem_posting_fbs_list_marketplaceservices `json:"item_services"`

	// Currency of your prices. It matches the currency set in the personal account settings
	CurrencyCode string `json:"currency_code"`

	// Price before discounts. Displayed strikethrough on the product description page
	OldPrice float64 `json:"old_price"`

	// Payment to the seller
	Payout float64 `json:"payout"`

	// Delivery details.
	//
	// Returns `null`
	Picking ResponseItem_posting_fbs_list_financialdataproductpicking `json:"picking"`

	// Product price including discounts. This value is shown on the product description page
	Price float64 `json:"price"`

	// Product identifier
	ProductId int64 `json:"product_id"`

	// Product quantity in the shipment
	Quantity int64 `json:"quantity"`

	// Discount percentage
	TotalDiscountPercent float64 `json:"total_discount_percent"`

	// Discount amount
	TotalDiscountValue float64 `json:"total_discount_value"`
}

type ResponseItem_posting_fbs_list_financialdataproductpicking struct {
	// Delivery cost
	Amount float64 `json:"amount"`

	// Delivery date and time
	Moment time.Time `json:"moment"`

	// Bulky products or not
	Tag string `json:"tag"`
}

func NewRequestParams_posting_fbs_list(since, to, sincestat, tostat time.Time, offset, limit int64, sortingdir SortingDirection) *RequestParams_posting_fbs_list {
	req := &RequestParams_posting_fbs_list{
		Direction: sortingdir,
		Limit:     limit,
		Offset:    offset,
		With: RequestItem_posting_fbs_list_with{
			AnalyticsData: true,
			Barcodes:      true,
			FinancialData: true,
			Translit:      true,
		},
		Filter: RequestItem_posting_fbs_list_filter{
			Since: since.UTC().Format(time_format),
			To:    to.UTC().Format(time_format),
		},
	}
	if !sincestat.IsZero() {
		req.Filter.LastChangedStatusDate.From = sincestat.UTC().Format(time_format)
	}
	if !tostat.IsZero() {
		req.Filter.LastChangedStatusDate.To = tostat.UTC().Format(time_format)
	}
	return req
}

func (cl *OzonClient) GetPostingFbsList(params *RequestParams_posting_fbs_list) (*Response_posting_fbs_list, error) {
	req, err := cl.newRequest(http.MethodPost, url_posting_fbs_list, params)
	if err != nil {
		return nil, err
	}
	resp := &Response_posting_fbs_list{}

	response, err := cl.doRequest(req, resp)
	if err != nil {
		return nil, err
	}
	resp.baseResponse = response.base
	//response.pasteBase(&resp.baseResponse)

	return resp, nil
}

func (cl *OzonClient) GetPostingFbsListDebug(params *RequestParams_posting_fbs_list) (*Response_posting_fbs_list, *http.Request, *http.Response, []byte, error) {
	req, err := cl.newRequest(http.MethodPost, url_posting_fbs_list, params)
	if err != nil {
		return nil, nil, nil, nil, err
	}
	reqq, _ := cl.newRequest(http.MethodPost, url_posting_fbs_list, params)

	resp := &Response_posting_fbs_list{}

	response, respp, resppbody, err := cl.doRequestDebug(req, resp)
	if err != nil {
		return nil, reqq, nil, nil, err
	}
	resp.baseResponse = response.base

	return resp, reqq, respp, resppbody, nil
}
