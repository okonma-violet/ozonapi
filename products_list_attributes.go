package ozonapi

import (
	"errors"
	"net/http"
	"strconv"
)

// https://docs.ozon.ru/api/seller/?abt_att=1&origin_referer=docs.ozon.ru#operation/ProductAPI_GetProductAttributesV4

type RequestParams_products_list_attributes struct {
	// Filter by product
	Filter RequestFilter_products_list_attributes `json:"filter"`
	// Identifier of the last value on the page. Leave this field blank in the first request.
	// To get the next values, specify last_id from the response of the previous request
	LastId string `json:"last_id"`
	// Number of values per page. Minimum is 1, maximum is 1000
	Limit int64 `json:"limit"`
	// Параметр, по которому товары будут отсортированы:
	// sku — сортировка по идентификатору товара в системе Ozon;
	// offer_id — сортировка по артикулу товара;
	// id — сортировка по идентификатору товара;
	// title — сортировка по названию товара.
	SortBy string `json:"sort_by"`
	// Направление сортировки:
	// asc — по возрастанию,
	// desc — по убыванию.
	SortDir string `json:"sort_dir"`
}
type RequestFilter_products_list_attributes struct {
	// Filter by the offer_id parameter. You can pass a list of values in this parameter
	OfferId []string `json:"offer_id"`

	// Filter by the product_id parameter. You can pass a list of values in this parameter
	ProductId []string `json:"product_id"`

	// Filter by product visibility
	Visibility VisibilityFilter `json:"visibility"`
}

type Response_products_list_attributes struct {
	baseResponse
	ResultItems []ResponseItem_products_list_attributes `json:"result"`
	// Identifier of the last value on the page.
	// To get the next values, specify the recieved value in the next request in the last_id parameter
	LastId string `json:"last_id"`
	// Total number of products
	Total int32 `json:"total"`
}

type ResponseItem_products_list_attributes struct {
	ProductId int64  `json:"id"`
	OfferId   string `json:"offer_id"`
	SKU       int64  `json:"sku"`

	// Список характеристик товара
	Attributes []ResponseItem_products_list_attributes_attrib `json:"attributes"`
	// Список идентификаторов характеристик со значением по умолчанию
	AttributesWithDefaults []int64 `json:"attributes_with_defaults"`
	// Массив вложенных характеристик. dictionaryValueId ??????
	ComplexAttributes []ResponseItem_products_list_attributes_attrib `json:"complex_attributes"`

	// Ссылка на главное изображение товара
	PrimaryImage string `json:"primary_image"`
	// Массив ссылок на изображения товара. Порядок изображений аналогичен порядку в карточке товаров
	Images []string `json:"images"`

	// Информация о модели
	ModelInfo ResponseItem_products_list_attributes_modelinfo `json:"model_info"`
	// Массив PDF-файлов
	PDFList []ResponseItem_products_list_attributes_pdf `json:"pdf_list"`

	// Идентификатор категории
	CategoryId int64 `json:"description_category_id"`
	// Идентификатор типа товара
	TypeId int64 `json:"type_id,omitempty"`

	// Штрихкод
	Barcode string `json:"barcode"`
	// Все штрихкоды товара
	Barcodes []string `json:"barcodes"`
	// Маркетинговый цвет
	Color string `json:"color_image"`
	// Название товара. <= 500 characters
	Name string `json:"name"`

	// Единица измерения габаритов
	DimensionUnit DimensionUnit `json:"dimension_unit"`
	// Глубина упаковки
	Depth int64 `json:"depth"`
	// Высота упаковки
	Height int64 `json:"height"`
	// Ширина упаковки
	Width int64 `json:"width"`
	// Единица измерения веса
	WeightUnit WeightUnit `json:"weight_unit"`
	// Вес товара в упаковке. Предельное значение — 1000 килограммов или конвертированная величина в других единицах измерения
	Weight int32 `json:"weight"`
}

type ResponseItem_products_list_attributes_attrib struct {
	// Идентификатор характеристики, которая поддерживает вложенные свойства. Например, у характеристики «Процессор» есть вложенные характеристики «Производитель», «L2 Cache» и другие. У каждой из вложенных характеристик может быть несколько вариантов значений.
	ComplexId int64 `json:"complex_id"`
	// Идентификатор характеристики
	Id int64 `json:"id"`
	// Массив вложенных значений характеристики
	Values []ResponseItem_products_list_attributes_attrib_value `json:"values"`
}
type ResponseItem_products_list_attributes_attrib_value struct {
	// Идентификатор справочника
	DictionaryValueId int64 `json:"dictionary_value_id"`
	// Значение из справочника
	Value string `json:"value"`
}
type ResponseItem_products_list_attributes_modelinfo struct {
	// Идентификатор модели
	ModelId int64 `json:"model_id"`
	// Количество объединённых товаров модели
	Count int64 `json:"count"`
}
type ResponseItem_products_list_attributes_pdf struct {
	// Путь к PDF-файлу
	Filename string `json:"file_name"`
	// Название файла
	Name string `json:"name"`
}

func NewRequestParams_products_list_attributes(visibility VisibilityFilter, offerids []string, productids []string, lastid string, limit int) *RequestParams_products_list_attributes {
	if limit <= 0 || limit > RequestItems_cap_products_list_attributes {
		panic("incorrect limit (<0 or >ozon's limit) : " + strconv.Itoa(int(limit)))
	}

	return &RequestParams_products_list_attributes{
		Filter: RequestFilter_products_list_attributes{
			Visibility: visibility,
			ProductId:  productids,
			OfferId:    offerids,
		},
		LastId: lastid,
		Limit:  int64(limit),
	}
}

func (rp *RequestParams_products_list_attributes) AddOfferId(offerid string) bool {
	if rp.Filter.OfferId == nil {
		rp.Filter.OfferId = make([]string, 0, RequestItems_cap_products_list_attributes)
	}
	if len(rp.Filter.OfferId) < RequestItems_cap_products_list_attributes {
		rp.Filter.OfferId = append(rp.Filter.OfferId, offerid)
		return true
	}
	return false
}
func (rp *RequestParams_products_list_attributes) AddOfferIdR(offerids []string) bool {
	if len(rp.Filter.OfferId)+len(offerids) < RequestItems_cap_products_list_attributes {
		if rp.Filter.OfferId == nil {
			rp.Filter.OfferId = offerids
		} else {
			rp.Filter.OfferId = append(rp.Filter.OfferId, offerids...)
		}
		return true
	}
	return false
}
func (rp *RequestParams_products_list_attributes) AddProductId(productid int64) bool {
	if rp.Filter.ProductId == nil {
		rp.Filter.ProductId = make([]string, 0, RequestItems_cap_products_list_attributes)
	}
	if len(rp.Filter.ProductId) < RequestItems_cap_products_list_attributes {
		rp.Filter.ProductId = append(rp.Filter.ProductId, strconv.FormatInt(productid, 10))
		return true
	}
	return false
}
func (rp *RequestParams_products_list_attributes) AddProductIdR(productids []int64) bool {
	if len(rp.Filter.ProductId)+len(productids) < RequestItems_cap_products_list_attributes {
		if rp.Filter.ProductId == nil {
			rp.Filter.ProductId = make([]string, 0, RequestItems_cap_products_list_attributes)
		}
		for i := range productids {
			rp.Filter.ProductId = append(rp.Filter.ProductId, strconv.FormatInt(productids[i], 10))
		}
		return true
	}
	return false
}

func (cl *OzonClient) GetProductsListAttributes(params *RequestParams_products_list_attributes) (*Response_products_list_attributes, error) {
	if params.Filter.OfferId != nil && params.Filter.ProductId != nil {
		return nil, errors.New("cannot mix offerids/productids/skus in this request")
	}
	req, err := cl.newRequest(http.MethodPost, url_products_list_attributes, params)
	if err != nil {
		return nil, err
	}
	resp := &Response_products_list_attributes{}

	response, err := cl.doRequest(req, resp)
	if err != nil {
		return nil, err
	}
	resp.baseResponse = response.base
	//response.pasteBase(&resp.baseResponse)

	return resp, nil
}
