package ozonapi

//shipment_act_fbs_create

import (
	"net/http"
	"time"
)

type RequestParams_shipment_act_fbs_create struct {
	// Number of package units.
	//
	// Use this parameter if you have trusted acceptance enabled and ship orders by package units.
	// If you do not have trusted acceptance enabled, skip it
	ContainersCount int `json:"containers_count"`

	// Delivery method identifier
	// id нашего склада в отправлении (??)
	DeliveryMethodId int64 `json:"delivery_method_id"`

	// Shipping date.
	//
	// To make documents printing available before the shipping day,
	// enable Printing the acceptance certificate in advance in your personal account under the method settings.
	// The time for packaging (packaging SLA) should be more than 13 hours
	DepartureDate time.Time `json:"departure_date"`
}

type Response_shipment_act_fbs_create struct {
	baseResponse
	Result ResponseItem_shipment_act_fbs_create `json:"result"`
}
type ResponseItem_shipment_act_fbs_create struct {
	// Document generation task number
	Id int64 `json:"id"`
}

func NewRequestParams_shipment_act_fbs_create(container_count int, delivery_method_id int64, departuredate time.Time) *RequestParams_shipment_act_fbs_create {
	return &RequestParams_shipment_act_fbs_create{
		ContainersCount:  container_count,
		DeliveryMethodId: delivery_method_id,
		DepartureDate:    departuredate,
	}
}

func (cl *OzonClient) CreateShipmentActFbs(params *RequestParams_shipment_act_fbs_create) (*Response_shipment_act_fbs_create, error) {
	req, err := cl.newRequest(http.MethodPost, url_shipment_act_fbs_create, params)
	if err != nil {
		return nil, err
	}
	resp := &Response_shipment_act_fbs_create{}

	response, err := cl.doRequest(req, resp)
	if err != nil {
		return nil, err
	}
	resp.baseResponse = response.base
	//response.pasteBase(&resp.baseResponse)

	return resp, nil
}
