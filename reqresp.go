package ozonapi

import (
	"bytes"
	"encoding/json"
	"strconv"
	"strings"

	"io"
	"net/http"
)

type Response struct {
	base baseResponse
	data interface{}
}

type baseResponse struct {
	StatusCode int
	Code       int                  `json:"code"`
	Details    []baseResponseDetail `json:"details"`
	Message    string               `json:"message"`
}
type baseResponseDetail struct {
	TypeUrl string `json:"typeUrl"`
	Value   string `json:"value"`
}

// strings only base response data, use only for not ok statuscode
func (br baseResponse) String() string {
	rd := make([]string, len(br.Details))
	for i := 0; i < len(br.Details); i++ {
		rd[i] = br.Details[i].TypeUrl + "-" + br.Details[i].Value
	}
	return "statuscode: " + strconv.Itoa(br.StatusCode) + "; code: " + strconv.Itoa(br.Code) + "; details: [" + strings.Join(rd, ";") + "]; message: " + br.Message
}

func (c *OzonClient) newRequest(method string, url string, bodyParams interface{}) (*http.Request, error) {
	bodyJson, err := json.Marshal(bodyParams)
	if err != nil {
		return nil, err
	}
	// fmt.Println("-----req", url, string(bodyJson))
	url = url_base + url
	req, err := http.NewRequest(method, url, bytes.NewBuffer(bodyJson))
	if err != nil {
		return nil, err
	}

	for k, v := range c.options {
		req.Header.Set(k, v)
	}

	return req, nil
}

// func (c *OzonClient) newRequests(method string, url string, bodyParams []interface{}) ([]*http.Request, error) {
// 	reqs := make([]*http.Request, len(bodyParams))
// 	var err error
// 	for i := 0; i < len(bodyParams); i++ {
// 		if reqs[i], err = c.newRequest(method, url, bodyParams[i]); err != nil {
// 			return nil, err
// 		}
// 	}
// 	return reqs, nil
// }

func (c *OzonClient) doRequest(req *http.Request, respbody interface{}) (*Response, error) {
	if c.ratelimiter != nil {
		if !c.ratelimiter.wait() {
			return nil, errLimiterClosed
		}
	}

	httpresp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer httpresp.Body.Close()

	body, err := io.ReadAll(httpresp.Body)
	if err != nil {
		return nil, err
	}
	// fmt.Println("+++++resp", string(body))
	response := &Response{data: respbody}
	response.base.StatusCode = httpresp.StatusCode
	if httpresp.StatusCode == http.StatusOK {

		//mm := make(map[string]interface{})
		//err = json.Unmarshal(body, &mm)
		//fmt.Println(mm)
		//fmt.Println(response.data)
		err = json.Unmarshal(body, &response.data)
	} else {
		err = json.Unmarshal(body, &response.base)
	}
	if err != nil {
		return nil, err
	}

	return response, nil
}

// TODO: сделать эту функцию базовой, единственной
func (c *OzonClient) doRequestDebug(req *http.Request, respbody interface{}) (*Response, *http.Response, []byte, error) {
	if c.ratelimiter != nil {
		if !c.ratelimiter.wait() {
			return nil, nil, nil, errLimiterClosed
		}
	}

	httpresp, err := c.client.Do(req)
	if err != nil {
		return nil, nil, nil, err
	}
	defer httpresp.Body.Close()

	body, err := io.ReadAll(httpresp.Body)
	if err != nil {
		return nil, nil, nil, err
	}
	// fmt.Println(string(body))
	response := &Response{data: respbody}
	response.base.StatusCode = httpresp.StatusCode
	if httpresp.StatusCode == http.StatusOK {
		//mm := make(map[string]interface{})
		//err = json.Unmarshal(body, &mm)
		//fmt.Println(mm)
		//fmt.Println(response.data)
		err = json.Unmarshal(body, &response.data)
	} else {
		err = json.Unmarshal(body, &response.base)
	}
	if err != nil {
		return nil, httpresp, body, err
	}

	return response, httpresp, body, nil
}

// func resliceByItemCap[T any](maxitems int, items []T) [][]T {
// 	fullparts := len(items) / maxitems
// 	var res [][]T

// 	if len(items)%maxitems > 0 {
// 		res = make([][]T, fullparts+1)
// 	} else {
// 		res = make([][]T, fullparts)
// 	}
// 	l, r := 0, maxitems
// 	for i := 0; i < fullparts; i++ {
// 		res[i] = items[l:r]
// 		l += maxitems
// 		r += maxitems
// 	}
// 	if fullparts != len(res) {
// 		res[len(res)-1] = items[l:]
// 	}
// 	return res
// }

// func (r *Response) pasteBase(br *baseResponse) {
// 	br.Code = r.base.Code
// 	br.Details = r.base.Details
// 	br.StatusCode = r.base.StatusCode
// 	br.Message = r.base.Message
// }
