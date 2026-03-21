package ozonapi

import (
	"net/http"
	"strconv"
	"time"
)

// https://docs.ozon.ru/api/seller/?abt_att=1&origin_referer=docs.ozon.ru#operation/ProductAPI_GetProductInfoList

type RequestParams_products_list_details struct {
	// Filter by the offer_id parameter. You can pass a list of values in this parameter
	OfferId []string `json:"offer_id"`

	// Filter by the product_id parameter. You can pass a list of values in this parameter
	ProductId []string `json:"product_id"`

	// Filter by the sku parameter. You can pass a list of values in this parameter
	SKU []string `json:"sku"`
}

type Response_products_list_details struct {
	baseResponse
	Items []ResponseItem_products_list_details `json:"items"`
}

type ResponseItem_products_list_details struct {
	// All product barcodes
	Barcodes []string `json:"barcodes"`

	// Изображение цвета товара
	ColorImage []string `json:"color_image"`

	// Commission fees details
	// !Contains array of comms for EACH sale schema, identity required by "SaleSchema" field
	Commissions []ResponseItem_products_list_details_commissions `json:"commissions"`

	// Date and time when the product was created
	CreatedAt time.Time `json:"created_at"`

	// Currency of your prices. It matches the currency set in the personal account settings
	CurrencyCode string `json:"currency_code"`

	// Category identifier
	DescriptionCategoryId int64 `json:"description_category_id"`

	// Остатки уценённого товара на складе Ozon
	DiscountedFboStocks int64 `json:"discounted_fbo_stocks"`

	// Информация об ошибках при создании или валидации товара
	Errors []ResponseItem_products_list_details_error `json:"errors"`

	// true if the product has markdown equivalents at the Ozon warehouse
	HasDiscountedFboItem bool `json:"has_discounted_fbo_item"`

	// Product id
	Id int64 `json:"id"`

	// An array of links to images. The images in the array are arranged in the order of their arrangement on the site. If the `primary_image` parameter is not specified, the first image in the list is the main one for the product
	Images []string `json:"images"`

	// Array of 360 images
	Images360 []string `json:"images360"`

	// true, если товар архивирован вручную
	IsArchived bool `json:"is_archived"`

	// true, если товар архивирован автоматически
	IsAutoArchived bool `json:"is_autoarchived"`

	// Признак, является ли товар уценённым:
	// Если товар создавался продавцом как уценённый — true.
	// Если товар не уценённый или был уценён Ozon — false
	IsDiscounted bool `json:"is_discounted"`

	// Indication of a bulky product
	IsKGT bool `json:"is_kgt"`

	// If prepayment is possible, the value is true
	IsPrepaymentAllowed bool `json:"is_prepayment_allowed"`

	// Признак супер-товара
	IsSuper bool `json:"is_super"`

	// The price of the product including all promotion discounts. This value will be shown on the Ozon storefront,
	// stringed float
	MarketingPrice string `json:"marketing_price"`

	// Минимальная цена товара после применения акций
	// Приходит пустым, отключен?
	MinPrice string `json:"min_price"`

	// Информация о модели товара
	ModelInfo ResponseItem_products_list_details_modelinfo `json:"model_info"`

	// Name
	Name string `json:"name"`

	// Product identifier in the seller's system
	OfferId string `json:"offer_id"`

	// Price before discounts. Displayed strikethrough on the product description page,
	// stringed float
	OldPrice string `json:"old_price"`

	// Product price including discounts. This value is shown on the product description page,
	// stringed float
	Price string `json:"price"`

	// Product price indexes
	PriceIndexes ResponseItem_products_list_details_priceindexes `json:"price_indexes"`

	// Main product image
	PrimaryImage []string `json:"primary_image"`

	// Details about the sources of similar offers. Learn more in Help Сenter
	Sources []ResponseItem_products_list_details_source `json:"sources"`

	// Product state description
	Statuses ResponseItem_products_list_details_statuses `json:"statuses"`

	// Details about product stocks
	Stocks ResponseItem_products_list_details_stocks `json:"stocks"`

	// Идентификатор типа товара
	TypeId int64 `json:"type_id"`

	// Date of the last product update
	UpdatedAt time.Time `json:"updated_at"`

	// Ставка НДС для товара
	VAT string `json:"vat"`

	// Product visibility settings
	VisibilityDetails ResponseItem_products_list_details_visibility `json:"visibility_details"`

	// Объёмный вес товара
	VolumeWeight float64 `json:"volume_weights"`
}
type ResponseItem_products_list_details_commissions struct {
	// Delivery cost
	DeliveryAmount float64 `json:"delivery_amount"`

	// Commission percentage
	Percent float64 `json:"percent"`

	// Return cost
	ReturnAmount float64 `json:"return_amount"`

	// Sale scheme
	SaleSchema SaleSchema `json:"sale_schema"`

	// Commission fee amount
	Value float64 `json:"value"`
}

type ResponseItem_products_list_details_priceindexes struct {
	// Default: "COLOR_INDEX_UNSPECIFIED"
	// Виды индекса цен:
	// COLOR_INDEX_UNSPECIFIED — не определён,
	// COLOR_INDEX_WITHOUT_INDEX — отсутствует,
	// COLOR_INDEX_GREEN — выгодный,
	// COLOR_INDEX_YELLOW — умеренный,
	// COLOR_INDEX_RED — невыгодный
	ColorIndex string `json:"color_index"`

	// Цена товара у конкурентов на других площадках
	ExternalIndexData ResponseItem_products_list_details_priceindexext `json:"external_index_data"`

	// Цена товара у конкурентов на Ozon
	OzonIndexData ResponseItem_products_list_details_priceindexext `json:"ozon_index_data"`

	// Price of your product on other marketplaces
	SelfMarketplaceIndexData ResponseItem_products_list_details_priceindexext `json:"self_marketplaces_index_data"`
}
type ResponseItem_products_list_details_priceindexext struct {
	// Минимальная цена товара у конкурентов на другой площадке/на Ozon или вашего товара на других площадках
	MinimalPrice string `json:"minimal_price"`

	// Валюта цены
	MinimalPriceCurrency string `json:"minimal_price_currency"`

	// Значение индекса цены
	PriceIndexValue float64 `json:"price_index_value"`
}
type ResponseItem_products_list_details_statuses struct {
	// Indiction that the product was created
	IsCreated bool `json:"is_created"`

	// Moderation status
	ModerateStatus string `json:"moderate_status"`

	// Product status
	Status string `json:"status"`

	// Product status description
	StatusDescription string `json:"status_description"`

	// Product status in which an error occurred
	StatusFailed string `json:"status_failed"`

	// Product status name
	StatusName string `json:"status_name"`

	// Tooltips for the current product status
	StatusTooltip string `json:"status_tooltip"`

	// The last time product status changed
	StatusUpdatedAt time.Time `json:"status_updated_at"`

	// Validation status
	ValidationStatus string `json:"validation_status"`
}
type ResponseItem_products_list_details_error struct {
	// Error attribute identifier
	AttributeId int64 `json:"attribute_id"`

	// Error code
	Code string `json:"code"`

	// Error field
	Field string `json:"field"`

	// Error level:
	// ERROR_LEVEL_UNSPECIFIED — не определено;
	// ERROR_LEVEL_ERROR — некритичная ошибка, товар можно продавать;
	// ERROR_LEVEL_WARNING — критичная ошибка, товар можно продавать;
	// ERROR_LEVEL_INTERNAL — критичная ошибка, товар нельзя продавать.
	Level string `json:"level"`

	// Product state in which an error occurred
	State string `json:"state"`

	// Error description
	Texts ResponseItem_products_list_details_error_texts `json:"texts"`
}

type ResponseItem_products_list_details_error_texts struct {
	// Attribute name
	AttributeName string `json:"attribute_name"`
	// Error description
	Description string `json:"description"`

	// Код ошибки в системе Ozon
	HintCode string `json:"hint_code"`

	// Текст ошибки
	Message string `json:"message"`

	// Краткое описание ошибки
	ShortDescription string `json:"short_description"`

	// Additional fields for error description
	Params []ResponseItem_products_list_details_error_params `json:"params"`
}

type ResponseItem_products_list_details_error_params struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type ResponseItem_products_list_details_modelinfo struct {
	// Количество товаров в ответе
	Count int64 `json:"count"`
	// Идентификатор модели товара
	ModelId int64 `json:"model_id"`
}

type ResponseItem_products_list_details_source struct {
	// Date and time when the product was created
	CreatedAt time.Time `json:"created_at"`

	// Список квантов с товарами
	QuantCode string `json:"quant_code"`

	// Default: "SHIPMENT_TYPE_UNSPECIFIED"
	// Тип упаковки:
	// SHIPMENT_TYPE_UNSPECIFIED — не указано;
	// SHIPMENT_TYPE_GENERAL — обычный товар;
	// SHIPMENT_TYPE_BOX — коробка;
	// SHIPMENT_TYPE_PALLET — палета
	ShipmentType string `json:"shipment_type"`

	// Product identifier in the Ozon system, SKU
	SKU int64 `json:"sku"`

	// Схема продажи, строчными
	// "SDS — это признак того, что товар переведен на единую SKU. Это означает, что товар может продаваться по любой схеме работы под одним и тем же SKU (например FBO и FBS)"
	Source string `json:"source"`
}
type ResponseItem_products_list_details_stocks struct {
	// true, если есть остаток на складах
	HasStock bool `json:"has_stock"`

	// Статус остатков товара
	Stocks []ResponseItem_products_list_details_stocksext `json:"reserved"`
}

type ResponseItem_products_list_details_stocksext struct {
	// Currently at the warehouse
	Present int32 `json:"present"`

	// Reserved
	Reserved int32 `json:"reserved"`

	// Product identifier in the Ozon system, SKU
	SKU int64 `json:"sku"`

	// Схема продажи
	Source string `json:"source"`
}

type ResponseItem_products_list_details_visibility struct {
	// If the price is set, the value is true
	HasPrice bool `json:"has_price"`

	// If there is stock at the warehouses, the value is true
	HasStock bool `json:"has_stock"`
}

// func NewRequestParams_products_list_details() *RequestParams_products_list_details {
// 	return &RequestParams_products_list_details{}
// }

// func (rp *RequestParams_products_list_details) AddOfferId(offerid string) bool {
// 	if rp.OfferId == nil {
// 		rp.OfferId = make([]string, 0, RequestItems_cap_products_list_details)
// 	}
// 	if len(rp.OfferId) < RequestItems_cap_products_list_details {
// 		rp.OfferId = append(rp.OfferId, offerid)
// 		return true
// 	}
// 	return false
// }

// func (rp *RequestParams_products_list_details) AddProductId(productid int64) bool {
// 	if rp.ProductId == nil {
// 		rp.ProductId = make([]string, 0, RequestItems_cap_products_list_details)
// 	}
// 	if len(rp.ProductId) < RequestItems_cap_products_list_details {
// 		rp.ProductId = append(rp.ProductId, strconv.FormatInt(productid, 10))
// 		return true
// 	}
// 	return false
// }

// use only one type of parameters;
func NewRequestParams_products_list_details(offerids []string, productids []string, skus []string) *RequestParams_products_list_details {
	return &RequestParams_products_list_details{
		OfferId:   offerids,
		ProductId: productids,
		SKU:       skus,
	}
}

func (rp *RequestParams_products_list_details) CapReached() bool {
	if len(rp.OfferId) >= RequestItems_cap_products_list_details {
		return true
	}
	if len(rp.ProductId) >= RequestItems_cap_products_list_details {
		return true
	}
	if len(rp.SKU) >= RequestItems_cap_products_list_details {
		return true
	}
	return false
}

// use only one type of parameters;
// return false if requestitems cap reached (and shit did not add)
func (rp *RequestParams_products_list_details) Add(offerid string, productid string, sku string) bool {
	if len(rp.OfferId)+len(rp.ProductId)+len(rp.SKU) < RequestItems_cap_products_list_details {
		if offerid != "" {
			rp.OfferId = append(rp.OfferId, offerid)
		}
		if productid != "" {
			rp.ProductId = append(rp.ProductId, productid)
		}
		if sku != "" {
			rp.SKU = append(rp.SKU, sku)
		}
		return true
	}
	return false
}
func (rp *RequestParams_products_list_details) AddOfferId(offerid string) bool {
	if len(rp.OfferId) < RequestItems_cap_products_list_details {
		if rp.OfferId == nil {
			rp.OfferId = make([]string, 0, RequestItems_cap_products_list_details)
		}
		rp.OfferId = append(rp.OfferId, offerid)
		return true
	}
	return false
}
func (rp *RequestParams_products_list_details) AddOfferIdR(offerids []string) bool {
	if len(rp.OfferId)+len(offerids) < RequestItems_cap_products_list_details {
		if rp.OfferId == nil {
			rp.OfferId = offerids
		} else {
			rp.OfferId = append(rp.OfferId, offerids...)
		}
		return true
	}
	return false
}
func (rp *RequestParams_products_list_details) AddProductIdR(productids []int64) bool {
	if len(rp.ProductId)+len(productids) < RequestItems_cap_products_list_details {
		if rp.ProductId == nil {
			rp.ProductId = make([]string, 0, RequestItems_cap_products_list_details)
		}
		for i := range productids {
			rp.ProductId = append(rp.ProductId, strconv.FormatInt(productids[i], 10))
		}
		return true
	}
	return false
}
func (cl *OzonClient) GetProductsListDetails(params *RequestParams_products_list_details) (*Response_products_list_details, error) {
	req, err := cl.newRequest(http.MethodPost, url_products_list_details, params)
	if err != nil {
		return nil, err
	}
	resp := &Response_products_list_details{}

	response, err := cl.doRequest(req, resp)
	if err != nil {
		return nil, err
	}
	resp.baseResponse = response.base
	//response.pasteBase(&resp.baseResponse)

	return resp, nil
}
