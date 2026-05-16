package ozonapi

import (
	"net/http"
	"time"
)

//https://docs.ozon.ru/api/seller/?abt_att=2&origin_referer=docs.ozon.ru&__rr=3#operation/FinanceAPI_FinanceCashFlowStatementList

type RequestParams_cash_flow_list struct {
	Date        RequestItem_cash_flow_list_date `json:"date"`
	WithDetails bool                            `json:"with_details"`
	Page        int                             `json:"page"`
	PageSize    int                             `json:"page_size"`
}

type RequestItem_cash_flow_list_date struct {
	From time.Time `json:"from"`
	To   time.Time `json:"to"`
}

type Response_cash_flow_list struct {
	baseResponse
	Result ResponseResult_cash_flow_list `json:"result"`
}

type ResponseResult_cash_flow_list struct {
	CashFlows []ResponseItem_cash_flow_list_cashflow `json:"cash_flows"`
	Details   []ResponseItem_cash_flow_list_detail   `json:"details"`
	PageCount int                                    `json:"page_count"`
}

type ResponseItem_cash_flow_list_cashflow struct {
	Period                      ResponseItem_cash_flow_list_period `json:"period"`
	OrdersAmount                float64                            `json:"orders_amount"`
	ReturnsAmount               float64                            `json:"returns_amount"`
	CommissionAmount            float64                            `json:"commission_amount"`
	ServicesAmount              float64                            `json:"services_amount"`
	ItemDeliveryAndReturnAmount float64                            `json:"item_delivery_and_return_amount"`
	CurrencyCode                string                             `json:"currency_code"`
}

type ResponseItem_cash_flow_list_detail struct {
	Period             ResponseItem_cash_flow_list_period    `json:"period"`
	BeginBalanceAmount float64                               `json:"begin_balance_amount"`
	Payments           []ResponseItem_cash_flow_list_payment `json:"payments"`
	Delivery           ResponseItem_cash_flow_list_delivery  `json:"delivery"`
	Return             ResponseItem_cash_flow_list_return    `json:"return"`
	Loan               float64                               `json:"loan"`
	InvoiceTransfer    float64                               `json:"invoice_transfer"`
	Rfbs               ResponseItem_cash_flow_list_rfbs      `json:"rfbs"`
	Services           ResponseItem_cash_flow_list_services  `json:"services"`
	Others             ResponseItem_cash_flow_list_others    `json:"others"`
	EndBalanceAmount   float64                               `json:"end_balance_amount"`
}

type ResponseItem_cash_flow_list_period struct {
	Id    int64     `json:"id"`
	Begin time.Time `json:"begin"`
	End   time.Time `json:"end"`
}

type ResponseItem_cash_flow_list_payment struct {
	Payment      float64 `json:"payment"`
	CurrencyCode string  `json:"currency_code"`
}

type ResponseItem_cash_flow_list_delivery struct {
	Total            float64                                      `json:"total"`
	Amount           float64                                      `json:"amount"`
	DeliveryServices ResponseItem_cash_flow_list_deliveryservices `json:"delivery_services"`
}

type ResponseItem_cash_flow_list_deliveryservices struct {
	Total float64                                   `json:"total"`
	Items []ResponseItem_cash_flow_list_serviceitem `json:"items"`
}

type ResponseItem_cash_flow_list_return struct {
	Total          float64                                    `json:"total"`
	Amount         float64                                    `json:"amount"`
	ReturnServices ResponseItem_cash_flow_list_returnservices `json:"return_services"`
}

type ResponseItem_cash_flow_list_returnservices struct {
	Total float64                                   `json:"total"`
	Items []ResponseItem_cash_flow_list_serviceitem `json:"items"`
}

type ResponseItem_cash_flow_list_rfbs struct {
	Total                      float64 `json:"total"`
	TransferDelivery           float64 `json:"transfer_delivery"`
	TransferDeliveryReturn     float64 `json:"transfer_delivery_return"`
	CompensationDeliveryReturn float64 `json:"compensation_delivery_return"`
	PartialCompensation        float64 `json:"partial_compensation"`
	PartialCompensationReturn  float64 `json:"partial_compensation_return"`
}

type ResponseItem_cash_flow_list_services struct {
	Total float64                                   `json:"total"`
	Items []ResponseItem_cash_flow_list_serviceitem `json:"items"`
}

type ResponseItem_cash_flow_list_others struct {
	Total float64                                   `json:"total"`
	Items []ResponseItem_cash_flow_list_serviceitem `json:"items"`
}

type ResponseItem_cash_flow_list_serviceitem struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func NewRequestParams_cash_flow_list(since, to time.Time, page, pagesize int) *RequestParams_cash_flow_list {
	req := &RequestParams_cash_flow_list{
		Date: RequestItem_cash_flow_list_date{
			From: since.UTC(),
			To:   to.UTC(),
		},
		WithDetails: true,
		Page:        page,
		PageSize:    pagesize,
	}

	return req
}

func (cl *OzonClient) Get_cash_flow_list(params *RequestParams_cash_flow_list) (*Response_cash_flow_list, error) {
	req, err := cl.newRequest(http.MethodPost, url_cash_flow_list, params)
	if err != nil {
		return nil, err
	}
	resp := &Response_cash_flow_list{}

	response, err := cl.doRequest(req, resp)
	if err != nil {
		return nil, err
	}
	resp.baseResponse = response.base
	return resp, nil
}
