package ozonapi

import (
	"net/http"
	"time"
)

// https://api-seller.ozon.ru/v4/posting/fbs/list

type RequestParams_posting_fbs_list struct {
	// Sorting direction
	Direction SortingDirection `json:"sort_dir"`

	//Filter
	Filter RequestItem_posting_fbs_list_filter `json:"filter"`

	// Number of shipments in the response:
	//   - maximum is 1000,
	//   - minimum is 1.
	Limit int64 `json:"limit"`

	// Указатель для выборки следующих данных
	Cursor string `json:"cursor"`

	// Additional fields that should be added to the response
	With RequestItem_posting_fbs_list_with `json:"with"`
}

type RequestItem_posting_fbs_list_filter struct {
	// Delivery method identifier
	DeliveryMethodId []int64 `json:"delivery_method_ids,omitempty"`

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
	ProviderIds []int64 `json:"provider_ids,omitempty"`

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
	Statuses []string `json:"statuses,omitempty"`

	// Warehouse identifier
	WarehouseIds []int64 `json:"warehouse_ids,omitempty"`
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

	// Юр
	LegalInfo bool `json:"legal_info"`

	// Transliterate the return values
	Translit bool `json:"translit"`
}

type Response_posting_fbs_list struct {
	baseResponse

	Cursor string `json:"cursor"`
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

	AvailableActions []string `json:"available_actions"`

	// Shipment barcodes
	Barcodes ResponseItem_posting_fbs_list_barcode `json:"barcodes"`

	// Cancellation details
	Cancellation ResponseItem_posting_fbs_list_cancellation `json:"cancellation"`

	// Customer details
	Customer ResponseItem_posting_fbs_list_customer `json:"customer"`

	// Информация о грузоместе
	Container any `json:"container"`

	// Тип сортировки грузоместа:
	// SORT — сортируемый;
	// NON-SORT — несортируемый.
	ContainerSortType any `json:"container_sort_type"`

	// Дата передачи отправления в доставку
	DeliveringDate time.Time `json:"delivering_date"`

	// Delivery method
	DeliveryMethod ResponseItem_posting_fbs_list_deliverymethod `json:"delivery_method"`

	// Схема доставки:
	// SDS — идентификатор единого SKU;
	// FBO — идентификатор товара, который продаётся со склада Ozon;
	// FBS — идентификатор товара, который продаётся со склада FBS;
	// Crossborder — идентификатор товара, который продаётся из-за границы.
	DeliveringScheme string `json:"delivery_schema"`

	// Идентификатор места назначения
	DestPlaceId int64 `json:"destination_place_id"`

	// Название места назначения
	DestPlaceName string `json:"destination_place_name"`

	// Информация о заказе с внешней платформы
	ExternalOrder any `json:"external_order"`

	// Data on the product cost, discount amount, payout and commission
	FinancialData ResponseItem_posting_fbs_list_financialdata `json:"financial_data"`

	// Start date and time of shipment processing
	InProccessAt time.Time `json:"in_process_at"`

	// если отправление доставляется методом «Самовывоз из магазина»
	IsClickCollect bool `json:"is_click_and_collect"`

	// If Ozon Express fast delivery was used — `true`
	IsExpress bool `json:"is_express"`

	// Indication that there is a multi-box product in the shipment
	// and you need to pass the number of boxes for it
	IsMultibox bool `json:"is_multibox"`

	// Number of boxes in which the product is packed
	MultiBoxQuantity int32 `json:"multi_box_qty"`

	// если товар — пересорт
	IsPeresortable bool `json:"is_presortable"`

	// Юридическая информация о покупателе
	LegalInfo any `json:"legal_info"`

	// Список товаров с дополнительными характеристиками
	Optional any `json:"optional"`

	// Identifier of the order to which the shipment belongs
	OrderId int64 `json:"order_id"`

	// Number of the order to which the shipment belongs
	OrderNumber string `json:"order_number"`

	// Number of the parent shipment which split resulted in the current shipment
	ParentPostingNumber string `json:"parent_posting_number"`

	// Shipment number
	PostingNumber string `json:"posting_number"`

	// Дата и время успешной валидации кода курьера. Проверьте код курьера методом /v1/posting/fbs/pick-up-code/verify
	PickupCodeVerifiedAt time.Time `json:"pickup_code_verified_at"`

	// List of products in the shipment
	Products []ResponseItem_posting_fbs_list_postingproduct `json:"products"`

	// The parameter is only relevant for bulky products
	// with a delivery by a third-party or integrated service
	PRROption PRROptionStatus `json:"prr_option"`

	// Идентификатор эконом-товара???
	QuantunId int64 `json:"quantum_id"`

	// если нужно заполнить атрибуты отслеживаемости
	RequireBLRTraceableAttrs bool `json:"require_blr_traceable_attrs"`

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

	// Статус отправления:
	// acceptance_in_progress — идёт приёмка;
	// arbitration — арбитраж;
	// awaiting_approve — ожидает подтверждения;
	// awaiting_deliver — ожидает отгрузки;
	// awaiting_packaging — ожидает упаковки;
	// awaiting_registration — ожидает регистрации;
	// awaiting_verification — создано;
	// cancelled — отменено;
	// cancelled_from_split_pending — отменено из-за разделения отправления;
	// client_arbitration — клиентский арбитраж доставки;
	// delivering — доставляется;
	// driver_pickup — у водителя;
	// not_accepted — не принято на сортировочном центре
	Status string `json:"status"`

	// Подстатус отправления:
	// posting_acceptance_in_progress— идёт приёмка;
	// posting_in_arbitration — арбитраж;
	// posting_created — создано;
	// posting_in_carriage — в перевозке;
	// posting_not_in_carriage — не добавлено в перевозку;
	// posting_registered — зарегистрировано;
	// posting_transferring_to_delivery, если status=awaiting_deliver — передаётся в доставку;
	// posting_awaiting_passport_data — ожидает паспортных данных;
	// posting_created — создано;
	// posting_awaiting_registration — ожидает регистрации;
	// posting_registration_error — ошибка регистрации;
	// posting_transferring_to_delivery, если status=awaiting_registration — передаётся курьеру;
	// posting_split_pending — создано;
	// posting_canceled — отменено;
	// posting_in_client_arbitration — клиентский арбитраж доставки;
	// posting_delivered — доставлено;
	// posting_received — получено;
	// posting_conditionally_delivered — условно доставлено;
	// posting_in_courier_service — курьер в пути;
	// posting_in_pickup_point — в пункте выдачи;
	// posting_on_way_to_city — в пути в ваш город;
	// posting_on_way_to_pickup_point — в пути в пункт выдачи;
	// posting_returned_to_warehouse — возвращено на склад;
	// posting_transferred_to_courier_service — передаётся в службу доставки;
	// posting_driver_pick_up — у водителя;
	// posting_not_in_sort_center — не принято на сортировочном центре;
	// ship_failed — сборка не удалась
	Substatus string `json:"substatus"`

	// Информация по тарификации отгрузки
	Tariffication any `json:"tariffication"`

	// Информация по тарификации отгрузки
	TarifficationSteps []any `json:"tariffication_steps"`

	// Type of integration with the delivery service
	TPLIntegrationType string `json:"tpl_integration_type"`

	// Shipment tracking number
	TrackingNumber string `json:"tracking_number"`

	// Объёмный вес товара
	VolWeight string `json:"volume_weight"`
}
type ResponseItem_posting_fbs_list_addressee struct {
	// Recipient name
	Name string `json:"name"`
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

	// Способ оплаты:
	// картой онлайн;
	// карта Ozon Банка;
	// автосписание с карты Ozon Банка при выдаче;
	// сохранённой картой при получении;
	// Система Быстрых Платежей;
	// Ozon Рассрочка;
	// оплата на расчётный счёт;
	// SberPay;
	// предоплата на стороне внешнего продавца.
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

	// // Services
	// PostingServices ResponseItem_posting_fbs_list_marketplaceservices `json:"posting_services"`

	// List of products in the shipment
	Products []ResponseItem_posting_fbs_list_financialdataproduct `json:"products"`
}

// type ResponseItem_posting_fbs_list_marketplaceservices struct {
// 	// Last mile
// 	DeliveryToCustomer float64 `json:"marketplace_service_item_deliv_to_customer"`

// 	// Pipeline
// 	DirectFlowTrans float64 `json:"marketplace_service_item_direct_flow_trans"`

// 	// Shipment processing in the fulfilment warehouse (FF)
// 	DropoffFF float64 `json:"marketplace_service_item_item_dropoff_ff"`

// 	// Shipment processing at the pick up point
// 	DropoffPVZ float64 `json:"marketplace_service_item_dropoff_pvz"`

// 	// Shipment processing at the sorting center
// 	DropoffSC float64 `json:"marketplace_service_item_dropoff_sc"`

// 	// Order packaging
// 	Fulfillment float64 `json:"marketplace_service_item_fulfillment"`

// 	// Transport arrival to the seller's address for shipments pick-up (Pick-up)
// 	Pickup float64 `json:"marketplace_service_item_pickup"`

// 	// Return processing
// 	ReturnAfterDeliveryToCustomer float64 `json:"marketplace_service_item_return_after_deliv_to_customer"`

// 	// Reverse pipeline
// 	ReturnFlowTrans float64 `json:"marketplace_service_item_return_flow_trans"`

// 	// Cancellations processing
// 	ReturnNotDeliveryToCustomer float64 `json:"marketplace_service_item_return_not_deliv_to_customer"`

// 	// Non-purchase processing
// 	ReturnPartGoodsCustomer float64 `json:"marketplace_service_item_return_part_goods_customer"`
// }

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

	// ЕСТЬ ЕЩЕ НО ПОХУЙ
}

type ResponseItem_posting_fbs_list_postingproduct struct {
	// Список IMEI мобильных устройств
	Imeis []string `json:"imei"`

	// если товар отслеживаемый
	IsBLRTraceable bool `json:"is_blr_traceable"`

	// если Ozon выкупил товар для доставки в ЕАЭС и другие страны
	IsMarketplaceBuyout bool `json:"is_marketplace_buyout"`

	// Product name
	Name string `json:"name"`

	// Product identifier in the seller's system
	OfferId string `json:"offer_id"`

	// Product price
	Price ResponseItem_posting_fbs_list_financialdataproductcustomerprice `json:"price"`

	// Product quantity in the shipment
	Quantity int32 `json:"quantity"`

	// Product identifier in the Ozon system, SKU
	SKU int64 `json:"sku"`

	ProductColor int64 `json:"product_color"`

	Weight int64 `json:"weight"`
}

type ResponseItem_posting_fbs_list_financialdataproductcommission struct {
	// Commission amount for the product
	CommissionAmount float64 `json:"amount"`

	// Commission percentage
	CommissionPercent int64 `json:"percent"`

	// Code of the currency used to calculate the commissions
	CommissionsCurrencyCode string `json:"currency"`
}

type ResponseItem_posting_fbs_list_financialdataproductcustomerprice struct {
	// Customer price
	ClientPrice string `json:"amount"`

	// Currency of your prices. It matches the currency set in the personal account settings
	CurrencyCode string `json:"currency"`
}

type ResponseItem_posting_fbs_list_financialdataproduct struct {
	// Actions
	Actions []string `json:"actions"`

	// Customer price
	CustomerPrice ResponseItem_posting_fbs_list_financialdataproductcustomerprice `json:"customer_price"`

	// Commission amount for the product
	Commission ResponseItem_posting_fbs_list_financialdataproductcommission `json:"commission"`

	// // Services
	// ItemServices ResponseItem_posting_fbs_list_marketplaceservices `json:"item_services"`

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

func NewRequestParams_posting_fbs_list(since, to, sincestat, tostat time.Time, cursor string, limit int64, sortingdir SortingDirection) *RequestParams_posting_fbs_list {
	req := &RequestParams_posting_fbs_list{
		Direction: sortingdir,
		Limit:     limit,
		Cursor:    cursor,
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
