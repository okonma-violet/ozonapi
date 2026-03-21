package ozonapi

import "net/http"

// https://docs.ozon.ru/api/seller/?abt_att=1&origin_referer=docs.ozon.ru#operation/DescriptionCategoryAPI_GetTree

type RequestParams_category_desc_tree struct {
	// Default: "DEFAULT"
	// Enum: "DEFAULT" "RU" "EN" "TR" "ZH_HANS".
	// Язык в ответе:
	// EN — английский,
	// RU — русский,
	// TR — турецкий,
	// ZH_HANS — китайский.
	// По умолчанию используется русский язык.
	Language string `json:"language"`
}

type Response_category_desc_tree struct {
	baseResponse
	Result []ResponseItem_category_desc_tree `json:"result"`
}
type ResponseItem_category_desc_tree struct {
	// Идентификатор категории
	Category_id int64 `json:"description_category_id"`

	// Название категории
	Category_name string `json:"category_name"`

	// Дерево подкатегорий
	Children []ResponseItem_category_desc_tree `json:"children"`

	// true, если в категории нельзя создавать товары. false, если можно
	Disabled bool `json:"disabled"`

	// Идентификатор типа товара
	Type_id int64 `json:"type_id"`

	// Название типа товара
	Type_name string `json:"type_name"`
}

// func NewRequestParams_category_desc_tree() *RequestParams_category_desc_tree {
// 	// return &RequestParams_category_desc_tree{
// 	// 	Language: "RU",
// 	// }
// }

func (cl *OzonClient) GetCategory_desc_tree() (*Response_category_desc_tree, error) {
	req, err := cl.newRequest(http.MethodPost, url_category_desc_tree, &RequestParams_category_desc_tree{
		Language: "RU",
	})
	if err != nil {
		return nil, err
	}
	resp := &Response_category_desc_tree{}

	response, err := cl.doRequest(req, resp)
	if err != nil {
		return nil, err
	}
	resp.baseResponse = response.base
	//response.pasteBase(&resp.baseResponse)

	return resp, nil
}
