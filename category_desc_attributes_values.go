package ozonapi

import "net/http"

// https://docs.ozon.ru/api/seller/?abt_att=1&origin_referer=docs.ozon.ru#operation/DescriptionCategoryAPI_GetAttributeValues

type RequestParams_category_desc_attributes_values struct {
	// Идентификатор характеристики
	Attribute_id int64 `json:"attribute_id"`

	// Идентификатор категории
	Category_id int64 `json:"description_category_id"`

	// Default: "DEFAULT"
	// Enum: "DEFAULT" "RU" "EN" "TR" "ZH_HANS".
	// Язык в ответе:
	// EN — английский,
	// RU — русский,
	// TR — турецкий,
	// ZH_HANS — китайский.
	// По умолчанию используется русский язык
	Language string `json:"language"`

	// Идентификатор типа товара
	Type_id int64 `json:"type_id"`

	// Идентификатор справочника, с которого нужно начать ответ. Если last_value_id — 10, то в ответе будут справочники, начиная с одиннадцатого
	Last_values_count int64 `json:"last_value_id"`

	// Количество значений в ответе: максимум — 2000, минимум — 1
	Limit int64 `json:"limit"`
}

type Response_category_desc_attributes_values struct {
	baseResponse
	Result   []ResponseItem_category_desc_attributes_values `json:"result"`
	Has_next bool                                           `json:"has_next"`
}
type ResponseItem_category_desc_attributes_values struct {
	// Идентификатор значения характеристики
	Id int64 `json:"id"`

	// Дополнительное описание
	Info string `json:"info"`

	// Ссылка на изображение
	Picture string `json:"picture"`

	// Значение характеристики товара
	Value string `json:"value"`
}

// limit - 2000 max
func NewRequestParams_category_desc_attributes_values(attribute_id, category_id, type_id int64, limit, last_values_count int) *RequestParams_category_desc_attributes_values {
	return &RequestParams_category_desc_attributes_values{
		Attribute_id:      attribute_id,
		Category_id:       category_id,
		Language:          "RU",
		Type_id:           type_id,
		Limit:             int64(limit),
		Last_values_count: int64(last_values_count),
	}
}

func (cl *OzonClient) GetCategory_desc_attributes_values(params *RequestParams_category_desc_attributes_values) (*Response_category_desc_attributes_values, error) {
	req, err := cl.newRequest(http.MethodPost, url_category_desc_attributes_values, params)
	if err != nil {
		return nil, err
	}
	resp := &Response_category_desc_attributes_values{}

	response, err := cl.doRequest(req, resp)
	if err != nil {
		return nil, err
	}
	resp.baseResponse = response.base
	//response.pasteBase(&resp.baseResponse)

	return resp, nil
}
