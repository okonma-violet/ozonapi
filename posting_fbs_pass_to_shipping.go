package ozonapi

//shipment_act_fbs_create

import (
	"net/http"
)

type RequestParams_posting_fbs_pass_to_shipping struct {
	PostingNumber []string `json:"posting_number"`
}

type Response_posting_fbs_pass_to_shipping struct {
	baseResponse
	Result bool `json:"result"`
}

func NewRequestParams_posting_fbs_pass_to_shipping() *RequestParams_posting_fbs_pass_to_shipping {
	return &RequestParams_posting_fbs_pass_to_shipping{
		PostingNumber: make([]string, 0, 1),
	}
}

func (rp *RequestParams_posting_fbs_pass_to_shipping) Add(postingnum string) *RequestParams_posting_fbs_pass_to_shipping {
	rp.PostingNumber = append(rp.PostingNumber, postingnum)
	return rp
}

// статус отправления изменится на awaiting_deliver (в озоне)
func (cl *OzonClient) PassPostingFbsToShipping(params *RequestParams_posting_fbs_pass_to_shipping) (*Response_posting_fbs_pass_to_shipping, error) {
	req, err := cl.newRequest(http.MethodPost, url_posting_fbs_pass_to_shipping, params)
	if err != nil {
		return nil, err
	}
	resp := &Response_posting_fbs_pass_to_shipping{}

	response, err := cl.doRequest(req, resp)
	if err != nil {
		return nil, err
	}
	resp.baseResponse = response.base
	//response.pasteBase(&resp.baseResponse)

	return resp, nil
}
