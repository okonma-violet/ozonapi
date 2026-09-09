package ozonapi

import (
	"net/http"
	"time"
)

// https://docs.ozon.ru/api/seller/?abt_att=1&origin_referer=docs.ozon.ru#operation/FinanceAPI_FinanceTransactionListV3

type RequestParams_returns_list struct {
	// Filter by product
	Filter RequestFilter_returns_list `json:"filter"`
	// Количество подгружаемых возвратов. Максимальное значение — 500.
	Limit int `json:"limit"`
	// Идентификатор последнего подгруженного возврата
	LastId int64 `json:"last_id"`
}

// Filter — основной фильтр запроса.
// Используйте только один из: logistic_return_date, storage_tariffication_start_date
// или visual_status_change_moment — иначе вернётся ошибка.
type RequestFilter_returns_list struct {
	LogisticReturnDate            *RequestFilter_date_range `json:"logistic_return_date,omitempty"`
	StorageTarifficationStartDate *RequestFilter_date_range `json:"storage_tariffication_start_date,omitempty"`
	VisualStatusChangeMoment      *RequestFilter_date_range `json:"visual_status_change_moment,omitempty"`

	OrderID              int64    `json:"order_id,omitempty"`
	PostingNumbers       []string `json:"posting_numbers,omitempty"` // не больше 50
	ProductName          string   `json:"product_name,omitempty"`
	OfferID              string   `json:"offer_id,omitempty"`
	VisualStatusName     string   `json:"visual_status_name,omitempty"` // см. VisualStatusName константы
	WarehouseID          int64    `json:"warehouse_id,omitempty"`
	Barcode              string   `json:"barcode,omitempty"`
	ReturnSchema         string   `json:"return_schema,omitempty"`          // "FBS" | "FBO"
	CompensationStatusID int32    `json:"compensation_status_id,omitempty"` // 1–4, см. CompensationStatus константы
}

// DateRange — общий тип для периодов фильтрации по дате.
type RequestFilter_date_range struct {
	TimeFrom time.Time `json:"time_from"`
	TimeTo   time.Time `json:"time_to"`
}

// VisualStatusName — возможные значения статуса возврата.
const (
	VisualStatusDisputeOpened                     = "DisputeOpened"                                 // открыт спор с покупателем
	VisualStatusOnSellerApproval                  = "OnSellerApproval"                              // на согласовании у продавца
	VisualStatusArrivedAtReturnPlace              = "ArrivedAtReturnPlace"                          // в пункте выдачи
	VisualStatusOnSellerClarification             = "OnSellerClarification"                         // на уточнении у продавца
	VisualStatusOnSellerClarificationAfterPartial = "OnSellerClarificationAfterPartialCompensation" // на уточнении после частичной компенсации
	VisualStatusOfferedPartialCompensation        = "OfferedPartialCompensation"                    // предложена частичная компенсация
	VisualStatusReturnMoneyApproved               = "ReturnMoneyApproved"                           // одобрен возврат денег
	VisualStatusPartialCompensationReturned       = "PartialCompensationReturned"                   // вернули часть денег
	VisualStatusCancelledDisputeNotOpen           = "CancelledDisputeNotOpen"                       // возврат отклонён, спор не открыт
	VisualStatusRejected                          = "Rejected"                                      // заявка отклонена
	VisualStatusCrmRejected                       = "CrmRejected"                                   // заявка отклонена Ozon
	VisualStatusCancelled                         = "Cancelled"                                     // заявка отменена
	VisualStatusApproved                          = "Approved"                                      // заявка одобрена продавцом
	VisualStatusApprovedByOzon                    = "ApprovedByOzon"                                // заявка одобрена Ozon
	VisualStatusReceivedBySeller                  = "ReceivedBySeller"                              // продавец получил возврат
	VisualStatusMovingToSeller                    = "MovingToSeller"                                // возврат на пути к продавцу
	VisualStatusReturningToSellerByCourier        = "ReturningToSellerByCourier"                    // курьер везёт возврат продавцу
	VisualStatusUtilizing                         = "Utilizing"                                     // на утилизации
	VisualStatusUtilized                          = "Utilized"                                      // утилизирован
	VisualStatusMoneyReturned                     = "MoneyReturned"                                 // покупателю вернули всю сумму
	VisualStatusPartialCompensationInProcess      = "PartialCompensationInProcess"                  // одобрен частичный возврат денег
	VisualStatusDisputeYouOpened                  = "DisputeYouOpened"                              // продавец открыл спор
	VisualStatusCompensationRejected              = "CompensationRejected"                          // отказано в компенсации
	VisualStatusDisputeOpening                    = "DisputeOpening"                                // обращение в поддержку отправлено
	VisualStatusCompensationOffered               = "CompensationOffered"                           // ожидает решения по компенсации
	VisualStatusWaitingCompensation               = "WaitingCompensation"                           // ожидает компенсации
	VisualStatusSendingError                      = "SendingError"                                  // ошибка при отправке обращения
	VisualStatusCompensationRejectedBySla         = "CompensationRejectedBySla"                     // истёк срок решения
	VisualStatusCompensationRejectedBySeller      = "CompensationRejectedBySeller"                  // продавец отказался от компенсации
	VisualStatusMovingToOzon                      = "MovingToOzon"                                  // едет на склад Ozon
	VisualStatusReturnedToOzon                    = "ReturnedToOzon"                                // на складе Ozon
	VisualStatusMoneyReturnedBySystem             = "MoneyReturnedBySystem"                         // быстрый возврат
	VisualStatusWaitingShipment                   = "WaitingShipment"                               // ожидает отправки
)

// CompensationStatus — возможные значения статуса компенсации.
const (
	CompensationStatusSent          int32 = 1 // отправлена
	CompensationStatusReceived      int32 = 2 // получена
	CompensationStatusCancelled     int32 = 3 // отменена
	CompensationStatusDecompensated int32 = 4 // проведена декомпенсация
)

// ReturnSchema — схемы доставки.
const (
	ReturnSchemaFBS = "FBS"
	ReturnSchemaFBO = "FBO"
)

type Response_returns_list struct {
	baseResponse
	Result ResponseResult_returns_list `json:"result"`
}
type ResponseResult_returns_list struct {
	Items   []ResponseItem_return `json:"returns"`
	HasNext bool                  `json:"has_next"`
}

// ReturnItem — информация об одном возврате.
type ResponseItem_return struct {
	Exemplars          []ResponseItem_exemplar           `json:"exemplars"`
	ID                 int64                             `json:"id"`
	CompanyID          int64                             `json:"company_id"`
	ReturnReasonName   string                            `json:"return_reason_name"`
	Type               string                            `json:"type"`   // см. ReturnType константы
	Schema             string                            `json:"schema"` // "FBS" | "FBO"
	OrderID            int64                             `json:"order_id"`
	OrderNumber        string                            `json:"order_number"`
	Place              *ResponseItem_warehouse           `json:"place"`
	TargetPlace        *ResponseItem_warehouse           `json:"target_place"`
	Storage            *ResponseItem_storage             `json:"storage"`
	Product            *ResponseItem_product             `json:"product"`
	Logistic           *ResponseItem_logistic            `json:"logistic"`
	Visual             *ResponseItem_visual              `json:"visual"`
	AdditionalInfo     *ResponseItem_additional_info     `json:"additional_info"`
	PostingNumber      string                            `json:"posting_number"`
	ClearingID         int64                             `json:"clearing_id"`
	ReturnClearingID   int64                             `json:"return_clearing_id"`
	CompensationStatus *ResponseItem_compensation_status `json:"compensation_status"`
}

// Exemplar — информация об экземпляре.
type ResponseItem_exemplar struct {
	ID int64 `json:"id"`
}

// Warehouse — склад, где находится или куда едет возврат.
type ResponseItem_warehouse struct {
	ID      int64  `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
}

// Storage — информация о хранении.
type ResponseItem_storage struct {
	Sum                     *ResponseItem_money `json:"sum"`
	TarifficationFirstDate  *time.Time          `json:"tariffication_first_date"`
	TarifficationStartDate  *time.Time          `json:"tariffication_start_date"`
	ArrivedMoment           *time.Time          `json:"arrived_moment"`
	Days                    int64               `json:"days"`
	UtilizationSum          *ResponseItem_money `json:"utilization_sum"`
	UtilizationForecastDate string              `json:"utilization_forecast_date"`
}

// Money — стоимость в валюте.
type ResponseItem_money struct {
	CurrencyCode string  `json:"currency_code"`
	Price        float64 `json:"price"`
}

// Product — информация о товаре.
type ResponseItem_product struct {
	SKU                    int64                                  `json:"sku"`
	OfferID                string                                 `json:"offer_id"`
	Name                   string                                 `json:"name"`
	Price                  *ResponseItem_money                    `json:"price"`
	PriceWithoutCommission *ResponseItem_price_without_commission `json:"price_without_commission"`
	Commission             *ResponseItem_commission               `json:"commission"`
}

// PriceWithoutCommission — стоимость товара без комиссии.
type ResponseItem_price_without_commission struct {
	CurrencyCode      string  `json:"currency_code"`
	Price             float64 `json:"price"`
	CommissionPercent float64 `json:"commission_percent"`
}

// Commission — информация о комиссии.
type ResponseItem_commission struct {
	Quantity int32 `json:"quantity"`
}

// Logistic — информация о возврате.
type ResponseItem_logistic struct {
	TechnicalReturnMoment           *time.Time `json:"technical_return_moment"`
	FinalMoment                     *time.Time `json:"final_moment"`
	CancelledWithCompensationMoment *time.Time `json:"cancelled_with_compensation_moment"`
	ReturnDate                      *time.Time `json:"return_date"`
	Barcode                         string     `json:"barcode"`
}

// Visual — информация о статусе возврата.
type ResponseItem_visual struct {
	Status       *ResponseItem_visual_status `json:"status"`
	ChangeMoment *time.Time                  `json:"change_moment"`
}

// VisualStatus — статус возврата.
type ResponseItem_visual_status struct {
	ID          int32  `json:"id"`
	DisplayName string `json:"display_name"`
	SysName     string `json:"sys_name"`
}

// AdditionalInfo — дополнительная информация.
type ResponseItem_additional_info struct {
	IsOpened      bool  `json:"is_opened"`
	IsSuperEconom bool  `json:"is_super_econom"`
	SourceID      int64 `json:"source_id"`
}

// CompensationStatus — информация о статусе компенсации.
type ResponseItem_compensation_status struct {
	Status       *ResponseItem_compensation_status_value `json:"status"`
	ChangeMoment *time.Time                              `json:"change_moment"`
}

// CompensationStatusValue — статус компенсации.
type ResponseItem_compensation_status_value struct {
	ID          int32  `json:"id"`
	DisplayName string `json:"display_name"`
	SysName     string `json:"sys_name"`
}

const (
	ReturnTypeCancellation  = "Cancellation"  // отмена (до вручения) // отдается с постингом
	ReturnTypeFullReturn    = "FullReturn"    // полный отказ при вручении // отдается с постингом
	ReturnTypePartialReturn = "PartialReturn" // частичный отказ при вручении
	ReturnTypeClientReturn  = "ClientReturn"  // клиентский возврат (после вручения)
	ReturnTypeUnknown       = "Unknown"       // технический возврат
)

// CompensationSysName — системные названия статуса компенсации.
const (
	CompensationSysSent               = "Sent"               // отправлена
	CompensationSysReceived           = "Received"           // получена
	CompensationSysCanceled           = "Canceled"           // отменена
	CompensationSysDecompensationSent = "DecompensationSent" // проведена декомпенсация
)

func NewRequestParams_returns_list(from, to time.Time, postingnums []string, limit int, lastid int64) *RequestParams_returns_list {
	if limit <= 0 || limit > RequestItems_cap_returns_list {
		panic("improper limit")
	}
	if lastid < 1 {
		panic("improper lastid")
	}

	rp := &RequestParams_returns_list{
		LastId: lastid,
		Limit:  limit,
	}

	if len(postingnums) > 0 {
		rp.Filter.PostingNumbers = postingnums
	}

	if !from.IsZero() && !to.IsZero() {
		rp.Filter.LogisticReturnDate = &RequestFilter_date_range{
			TimeFrom: from.UTC(),
			TimeTo:   to.UTC()}
	}
	return rp
}

func (cl *OzonClient) GetReturnsList(params *RequestParams_returns_list) (*Response_returns_list, error) {
	req, err := cl.newRequest(http.MethodPost, url_returns_list, params)
	if err != nil {
		return nil, err
	}
	resp := &Response_returns_list{}

	response, err := cl.doRequest(req, resp)
	if err != nil {
		return nil, err
	}
	resp.baseResponse = response.base
	//response.pasteBase(&resp.baseResponse)

	return resp, nil
}
