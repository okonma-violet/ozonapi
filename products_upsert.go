package ozonapi

import (
	"net/http"
	"strconv"
)

// https://docs.ozon.ru/api/seller/?abt_att=1&origin_referer=docs.ozon.ru#operation/ProductAPI_ImportProductsV3

type RequestParams_products_upsert struct {
	// Product prices details
	Items []RequestItem_products_upsert `json:"items"`
}

type RequestItem_products_upsert struct {
	// Product identifier in the seller's system
	OfferId string `json:"offer_id"`

	// Массив с характеристиками товара. Характеристики отличаются для разных категорий
	Attributes []RequestItem_products_upsert_attribute `json:"attributes,omitempty"`
	// Массив вложенных характеристик
	ComplexAttributes []RequestItem_products_upsert_attribute `json:"complex_attributes"`

	// Ссылка на главное изображение товара
	PrimaryImage string `json:"primary_image"`
	// Массив изображений. До 30 штук. Изображения показываются на сайте в таком же порядке, как в массиве.
	// Если не передать значение primary_image, первое изображение в массиве будет главным для товара.
	// Если вы передали значение primary_image, передайте до 29 изображений. Если параметр primary_image пустой, передайте до 30 изображений.
	// Формат: адрес ссылки на изображение в общедоступном облачном хранилище. Формат изображения по ссылке — JPG или PNG
	Images []string `json:"images"`
	// Массив изображений 360. До 70 штук.
	// Формат: адрес ссылки на изображение в общедоступном облачном хранилище. Формат изображения по ссылке — JPG
	Images360 []string `json:"images360"`

	// Список PDF-файлов
	PDFList []RequestItem_products_upsert_pdf `json:"pdf_list"`
	// Акции
	Promotions []RequestItem_products_upsert_promotion `json:"promotions"`

	// // Идентификатор категории
	CategoryId int64 `json:"description_category_id,omitempty"`
	// Новый идентификатор категории. Укажите его, если нужно изменить текущую категорию товара.
	NewCategoryId int64 `json:"new_description_category_id,omitempty"`
	// Идентификатор типа товара.
	// Значения можно получить из такого же параметра type_id в ответе метода /v1/description-category/tree.
	// При заполнении этого параметра можно не указывать в attibutes атрибут с параметром id:8229, type_id будет использоваться в приоритете
	TypeId int64 `json:"type_id,omitempty"`

	// Название товара. До 500 символов
	Name string `json:"name"`

	// Валюта ваших цен. Переданное значение должно совпадать с валютой, которая установлена в настройках личного кабинета. По умолчанию передаётся RUB — российский рубль
	CurrencyCode CurrencyCode `json:"currency_code"`
	// Цена товара с учётом скидок, отображается на карточке товара
	Price string `json:"price"`

	// Геоограничения — при необходимости заполните параметр в личном кабинете при создании или редактировании товара. Необязательный параметр
	GeoNames string `json:"geo_names,omitempty"`
	// Единица измерения габаритов
	DimensionUnit DimensionUnit `json:"dimension_unit"`
	// Глубина упаковки
	Depth int32 `json:"depth"`
	// Высота упаковки
	Height int32 `json:"height"`
	// Ширина упаковки
	Width int32 `json:"width"`

	// Единица измерения веса
	WeightUnit WeightUnit `json:"weight_unit"`
	// Вес товара в упаковке. Предельное значение — 1000 килограммов или конвертированная величина в других единицах измерения
	Weight int32 `json:"weight"`

	// хз че это
	// Default: "IS_CODE_SERVICE"
	// Enum: "IS_CODE_SERVICE" "IS_NO_CODE_SERVICE"
	ServiceType string `json:"service_type"`
	// Ставка НДС для товара:
	// 0 — не облагается НДС,
	// 0.05 — 5%,
	// 0.07 — 7%,
	// 0.1 — 10%,
	// 0.2 — 20%.
	VAT string `json:"vat"`
}

type RequestItem_products_upsert_attribute struct {
	// Идентификатор характеристики, которая поддерживает вложенные свойства. Например, у характеристики «Процессор» есть вложенные характеристики «Производитель», «L2 Cache» и другие. У каждой из вложенных характеристик может быть несколько вариантов значений.
	ComplexId int64 `json:"complex_id"`
	// Идентификатор характеристики
	Id int64 `json:"id"`
	// Массив вложенных значений характеристики
	Values []RequestItem_products_upsert_attribute_value `json:"values"`
}

type RequestItem_products_upsert_attribute_value struct {
	// Идентификатор справочника
	DictionaryValueId int64 `json:"dictionary_value_id"`
	// Значение из справочника
	Value string `json:"value"`
}

type RequestItem_products_upsert_pdf struct {
	// Индекс документа в хранилище, который задаёт порядок
	Index int64 `json:"index"`
	// Адрес файла
	SourceURL string `json:"src_url"`
	// Название файла
	Name string `json:"name"`
}

type RequestItem_products_upsert_promotion struct {
	// Default: "UNKNOWN"
	// Атрибут для действий с акцией:
	// ENABLE — включить,
	// DISABLE — выключить,
	// UNKNOWN — ничего не менять, передаётся по умолчанию.
	Operation string `json:"operation"`
	// Default: "REVIEWS_PROMO"
	// Value: "REVIEWS_PROMO"
	// Тип акции:
	// REVIEWS_PROMO — акция «Баллы за отзывы».
	Type string `json:"type"`
}

type Response_products_upsert struct {
	baseResponse
	Result ResponseResult_products_upsert `json:"result"`
}
type ResponseResult_products_upsert struct {
	TaskId int64 `json:"task_id"`
}

func NewRequestParams_products_upsert() *RequestParams_products_upsert {
	return &RequestParams_products_upsert{Items: make([]RequestItem_products_upsert, 0, RequestItems_cap_products_upsert)}
}

func (rp *RequestParams_products_upsert) CapReached() bool {
	return len(rp.Items) >= RequestItems_cap_products_upsert
}
func (rp *RequestParams_products_upsert) Len() int {
	return len(rp.Items)
}

// return false if request items cap reached (and shit did not add)
func (rp *RequestParams_products_upsert) Add(offerid string, price int, newcategoryid int64, typeid int64, depth, height, width, weight int, dimensionunit DimensionUnit, weightunit WeightUnit, attributes []RequestItem_products_upsert_attribute) bool {
	if len(rp.Items) < RequestItems_cap_prices_update {
		item := RequestItem_products_upsert{
			OfferId:       offerid,
			Price:         strconv.Itoa(price),
			Depth:         int32(depth),
			DimensionUnit: dimensionunit,
			WeightUnit:    weightunit,
			Height:        int32(height),
			Weight:        int32(weight),
			Width:         int32(width),
			NewCategoryId: newcategoryid,
			TypeId:        typeid,
		}
		if len(attributes) != 0 {
			item.Attributes = attributes
		}
		rp.Items = append(rp.Items, item)
		return true
	}
	return false
}

// return false if request items cap reached (and shit did not add)
func (rp *RequestParams_products_upsert) AddR(item RequestItem_products_upsert) bool {
	if len(rp.Items) < RequestItems_cap_prices_update {
		rp.Items = append(rp.Items, item)
		return true
	}
	return false
}

func (cl *OzonClient) UpsertProducts(params *RequestParams_products_upsert) (*Response_products_upsert, error) {
	req, err := cl.newRequest(http.MethodPost, url_products_upsert, params)
	if err != nil {
		return nil, err
	}

	resp := &Response_products_upsert{}

	response, err := cl.doRequest(req, resp)
	if err != nil {
		return nil, err
	}
	resp.baseResponse = response.base
	//response.pasteBase(&resp.baseResponse)

	return resp, nil
}
