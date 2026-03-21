package ozonapi

import (
	"encoding/json"
	"errors"
	"strings"
	"time"
)

var ErrClosedBatch = errors.New("closed batched client")

const url_base = "https://api-seller.ozon.ru"
const time_format = "2006-01-02T15:04:05Z"

const (
	url_prices_list   = "/v5/product/info/prices" // done
	url_prices_update = "/v1/product/import/prices"
	url_stocks_update = "/v2/products/stocks"

	url_products_description     = "/v1/product/info/description"
	url_products_list            = "/v3/product/list"      // done
	url_products_list_details    = "/v3/product/info/list" // done
	url_products_list_attributes = "/v4/product/info/attributes"
	url_products_upsert          = "/v3/product/import"

	url_posting_fbo_list = "/v2/posting/fbo/list"

	url_posting_fbs_list             = "/v3/posting/fbs/list"
	url_posting_fbs_pass_to_shipping = "/v2/posting/fbs/awaiting-delivery"

	url_shipment_act_fbs_create = "/v2/posting/fbs/act/create" //????

	url_category_desc_tree              = "/v1/description-category/tree"
	url_category_desc_attributes        = "/v1/description-category/attribute"
	url_category_desc_attributes_values = "/v1/description-category/attribute/values"

	url_transactions_list = "/v3/finance/transaction/list"
)

const (
	request_ratelimit_stocks_update = 70
)

const (
	RequestItems_cap_prices_update            = 1000
	RequestItems_cap_prices_list              = 1000
	RequestItems_cap_stocks_update            = 100
	RequestItems_cap_products_list            = 1000
	RequestItems_cap_products_list_details    = 1000
	RequestItems_cap_products_list_attributes = 1000
	RequestItems_cap_products_upsert          = 100
	RequestItems_cap_transactions_list        = 1000
)

//const cap_refresh_time = time.Minute

type SwitchingAttribute string

const (
	// enable
	SwitchingAttributeEnabled SwitchingAttribute = "ENABLED"
	// disable
	SwitchingAttributeDisabled SwitchingAttribute = "DISABLED"
	// don't change anything. default value
	SwitchingAttributeUnknown SwitchingAttribute = "UNKNOWN"
)

type CurrencyCode string

const (
	CurrencyCode_USD CurrencyCode = "USD"
	CurrencyCode_EUR CurrencyCode = "EUR"
	CurrencyCode_TRY CurrencyCode = "TRY"
	CurrencyCode_CNY CurrencyCode = "CNY"
	CurrencyCode_RUB CurrencyCode = "RUB"
	CurrencyCode_GBP CurrencyCode = "GBP"
)

type DimensionUnit string

const (
	DimensionUnit_MM DimensionUnit = "mm"
	DimensionUnit_CM DimensionUnit = "cm"
	DimensionUnit_IN DimensionUnit = "in"
)

type WeightUnit string

const (
	WeightUnit_G  WeightUnit = "g"
	WeightUnit_KG WeightUnit = "kg"
	WeightUnit_lb WeightUnit = "lb"
)

type VisibilityFilter string

const (
	ALL                      VisibilityFilter = "ALL"                      //все товары, кроме архивных.
	VISIBLE                  VisibilityFilter = "VISIBLE"                  //товары, которые видны покупателям.
	INVISIBLE                VisibilityFilter = "INVISIBLE"                //товары, которые не видны покупателям.
	EMPTY_STOCK              VisibilityFilter = "EMPTY_STOCK"              //товары, у которых не указано наличие.
	NOT_MODERATED            VisibilityFilter = "NOT_MODERATED"            //товары, которые не прошли модерацию.
	MODERATED                VisibilityFilter = "MODERATED"                //товары, которые прошли модерацию.
	DISABLED                 VisibilityFilter = "DISABLED"                 //товары, которые видны покупателям, но недоступны к покупке.
	STATE_FAILED             VisibilityFilter = "STATE_FAILED"             //товары, создание которых завершилось ошибкой.
	READY_TO_SUPPLY          VisibilityFilter = "READY_TO_SUPPLY"          //товары, готовые к поставке.
	VALIDATION_STATE_PENDING VisibilityFilter = "VALIDATION_STATE_PENDING" //товары, которые проходят проверку валидатором на премодерации.
	VALIDATION_STATE_FAIL    VisibilityFilter = "VALIDATION_STATE_FAIL"    //товары, которые не прошли проверку валидатором на премодерации.
	VALIDATION_STATE_SUCCESS VisibilityFilter = "VALIDATION_STATE_SUCCESS" //товары, которые прошли проверку валидатором на премодерации.
	TO_SUPPLY                VisibilityFilter = "TO_SUPPLY"                //товары, готовые к продаже.
	IN_SALE                  VisibilityFilter = "IN_SALE"                  //товары в продаже.
	REMOVED_FROM_SALE        VisibilityFilter = "REMOVED_FROM_SALE"        //товары, скрытые от покупателей.
	OVERPRICED               VisibilityFilter = "OVERPRICED"               //товары с завышенной ценой.
	CRITICALLY_OVERPRICED    VisibilityFilter = "CRITICALLY_OVERPRICED"    //товары со слишком завышенной ценой.
	EMPTY_BARCODE            VisibilityFilter = "EMPTY_BARCODE"            //товары без штрихкода.
	BARCODE_EXISTS           VisibilityFilter = "BARCODE_EXISTS"           //товары со штрихкодом.
	QUARANTINE               VisibilityFilter = "QUARANTINE"               //товары на карантине после изменения цены более чем на 50%.
	ARCHIVED                 VisibilityFilter = "ARCHIVED"                 //товары в архиве.
	OVERPRICED_WITH_STOCK    VisibilityFilter = "OVERPRICED_WITH_STOCK"    //товары в продаже со стоимостью выше, чем у конкурентов.
	PARTIAL_APPROVED         VisibilityFilter = "PARTIAL_APPROVED"         //товары в продаже с пустым или неполным описанием.
)

type SortingDirection string

const (
	ASC  SortingDirection = "ASC"
	DESC SortingDirection = "DESC"
)

type ShipmentStatus string

const (
	// acceptance is in progress
	AcceptanceInProgress ShipmentStatus = "acceptance_in_progress"
	// arbitration
	Arbitration ShipmentStatus = "arbitration"
	// client arbitration
	ClientArbitration ShipmentStatus = "client_arbitration"
	// awaiting confirmation
	AwaitingApprove ShipmentStatus = "awaiting_approve"
	// awaiting shipping
	AwaitingDeliver ShipmentStatus = "awaiting_deliver"
	// awaiting packaging
	AwaitingPackaging ShipmentStatus = "awaiting_packaging"
	// created
	AwaitingVerification ShipmentStatus = "awaiting_verification"
	// cancelled
	CancelledSubstatus ShipmentStatus = "cancelled"
	// delivered
	Delivered ShipmentStatus = "delivered"
	// delivery is in progress
	Delivering ShipmentStatus = "delivering"
	// picked up by driver
	DriverPickup ShipmentStatus = "driver_pickup"
	// not accepted at the sorting center
	NotAccepted ShipmentStatus = "not_accepted"
	// sent by the seller
	SentBySeller ShipmentStatus = "sent_by_seller"
)

type PRROptionStatus string

const (
	// carrying the bulky product using the elevator
	PRROptionLift PRROptionStatus = "lift"

	// carrying the bulky product upstairs
	PRROptionStairs PRROptionStatus = "stairs"

	// the customer canceled the service,
	// you don't need to lift the shipment
	PRROptionNone PRROptionStatus = "none"

	// delivery is included in the price.
	// According to the offer you need to
	// deliver products to the floor
	PRROptionDeliveryDefault PRROptionStatus = "delivery_default"
)

type SaleSchema string

const (
	SaleSchemaFBO  SaleSchema = "FBO"
	SaleSchemaFBS  SaleSchema = "FBS"
	SaleSchemaRFBS SaleSchema = "RFBS"
	SaleSchemaFBP  SaleSchema = "FBP"
)

type OzonTime_1 struct {
	time.Time
}

const ozontime_1_layout = "2006-01-02 15:04:05"

func (ot OzonTime_1) MarshalJSON() ([]byte, error) {
	if ot.Time.IsZero() {
		return json.Marshal(nil)
	}

	return json.Marshal(ot.Time.Format(ozontime_1_layout))
}
func (ot *OzonTime_1) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), `"`)
	if s == "null" || s == "" {
		return
	}
	ot.Time, err = time.Parse(ozontime_1_layout, s)
	return
}
func (ot *OzonTime_1) ParseFrom(str string) (err error) {
	ot.Time, err = time.Parse(ozontime_1_layout, str)
	return
}
func (ot OzonTime_1) In(loc *time.Location) OzonTime_1 {
	ot.Time = ot.Time.In(loc)
	return ot
}
func (ot OzonTime_1) String() string {
	return ot.Time.Format(ozontime_1_layout)
}
