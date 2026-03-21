package ozonapi

import "net/http"

// https://docs.ozon.ru/api/seller/?abt_att=1&origin_referer=docs.ozon.ru#operation/DescriptionCategoryAPI_GetAttributes

type RequestParams_category_desc_attributes struct {
	// Идентификатор категории
	Category_id int64 `json:"description_category_id"`

	// Default: "DEFAULT"
	// Enum: "DEFAULT" "RU" "EN" "TR" "ZH_HANS".
	// Язык в ответе:
	// EN — английский,
	// RU — русский,
	// TR — турецкий,
	// ZH_HANS — китайский.
	// По умолчанию используется русский язык.
	Language string `json:"language"`

	// Идентификатор типа товара
	Type_id int64 `json:"type_id"`
}

type Response_category_desc_attributes struct {
	baseResponse
	Result []ResponseItem_category_desc_attributes `json:"result"`
}
type ResponseItem_category_desc_attributes struct {
	// Идентификатор характеристики
	Id int64 `json:"id"`

	// Идентификатор комплексного атрибута
	// Как я понял - атрибуты под одним комплексайди ходят вместе, например видео (у которого атрибут ссылка, атрибут название и атрибут товары на видео)
	Attribute_complex_id int64 `json:"attribute_complex_id"`

	// Название
	Name string `json:"name"`

	// Описание характеристики
	Description string `json:"description"`

	// Тип характеристики
	Type string `json:"type"`

	// true, если характеристика — набор значений.
	// false, если характеристика — одно значение
	Is_collection bool `json:"is_collection"`

	// Признак обязательной характеристики:
	// true — обязательная характеристика,
	// false — характеристику можно не указывать
	Is_required bool `json:"is_required"`

	// Признак аспектного атрибута. Аспектный атрибут — характеристика, по которой отличаются товары одной модели.
	// Например, у одежды и обуви одной модели могут быть разные расцветки и размеры. То есть цвет и размер — это аспектные атрибуты.
	// Значения поля:
	// true — атрибут аспектный и его нельзя изменить после поставки товара на склад или продажи со своего склада.
	// false — атрибут не аспектный, можно изменить в любое время
	Is_aspect bool `json:"is_aspect"`

	// Максимальное количество значений для атрибута
	Max_value_count int64 `json:"max_value_count"`

	// Идентификатор группы характеристик
	Group_id int64 `json:"group_id"`

	// Название группы характеристик
	Group_name string `json:"group_name"`

	// Идентификатор справочника
	// Если у dictionary_id значение 0, у атрибута нет вложенных справочников. Если значение другое, то справочники есть
	Dictionary_id int64 `json:"dictionary_id"`

	// Признак, что значения словарного атрибута зависят от категории:
	// true — у атрибута разные значения для каждой категории.
	// false — у атрибута одинаковые значения для всех категорий
	Category_dependent bool `json:"category_dependent"`

	// Признак, что комплексная характеристика — набор значений:
	// true, если комплексная характеристика — набор значений,
	// false, если комплексная характеристика — одно значение
	Complex_is_collection bool `json:"complex_is_collection"`
}

func NewRequestParams_category_desc_attributes(category_id, type_id int64) *RequestParams_category_desc_attributes {
	return &RequestParams_category_desc_attributes{
		Category_id: category_id,
		Language:    "RU",
		Type_id:     type_id,
	}
}

func (cl *OzonClient) GetCategory_desc_attributes(params *RequestParams_category_desc_attributes) (*Response_category_desc_attributes, error) {
	req, err := cl.newRequest(http.MethodPost, url_category_desc_attributes, params)
	if err != nil {
		return nil, err
	}
	resp := &Response_category_desc_attributes{}

	response, err := cl.doRequest(req, resp)
	if err != nil {
		return nil, err
	}
	resp.baseResponse = response.base
	//response.pasteBase(&resp.baseResponse)

	return resp, nil
}
