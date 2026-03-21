package ozonapi

import (
	"net/http"
	"time"
)

// https://docs.ozon.ru/api/seller/?abt_att=1&origin_referer=docs.ozon.ru#operation/FinanceAPI_FinanceTransactionListV3

type RequestParams_transactions_list struct {
	// Filter by product
	Filter RequestFilter_transactions_list `json:"filter"`
	// Номер страницы, возвращаемой в запросе.
	Page int64 `json:"page"`
	// Количество элементов на странице, <= 1000
	PageSize int64 `json:"page_size"`
}

type RequestFilter_transactions_list struct {
	Date RequestFilter_transactions_list_date `json:"date"`
	// (в документации перечислены не все)
	// ClientReturnAgentOperation — получение возврата, отмены, невыкупа от покупателя;
	// MarketplaceMarketingActionCostOperation — услуги продвижения товаров;
	// MarketplaceSaleReviewsOperation — приобретение отзывов на платформе;
	// MarketplaceSellerCompensationOperation — прочие компенсации;
	// OperationAgentDeliveredToCustomer — доставка покупателю;
	// OperationAgentDeliveredToCustomerCanceled — доставка покупателю — исправленное начисление;
	// OperationAgentStornoDeliveredToCustomer — доставка покупателю — отмена начисления;
	// OperationClaim — начисление по претензии;
	// OperationCorrectionSeller — инвентаризация взаиморасчетов;
	// OperationDefectiveWriteOff — компенсация за повреждённый на складе товар;
	// OperationItemReturn — доставка и обработка возврата, отмены, невыкупа;
	// OperationLackWriteOff — компенсация за утерянный на складе товар;
	// OperationMarketplaceCrossDockServiceWriteOff — доставка товаров на склад Ozon (кросс-докинг);
	// OperationMarketplaceServiceStorage — услуга размещения товаров на складе;
	// OperationSetOff — взаимозачёт с другими договорами контрагента;
	// MarketplaceSellerReexposureDeliveryReturnOperation — перечисление за доставку от покупателя;
	// OperationReturnGoodsFBSofRMS — доставка и обработка возврата, отмены, невыкупа;
	// ReturnAgentOperationRFBS — возврат перечисления за доставку покупателю;
	// ItemAgentServiceStarsMembership — вознаграждение за услугу «Звёздные товары»;
	// MarketplaceSellerShippingCompensationReturnOperation — компенсация перечисления за доставку;
	// OperationMarketplaceServicePremiumCashback — услуга продвижения Premium;
	// MarketplaceServicePremiumPromotion — услуга продвижения Premium, фиксированная комиссия;
	// MarketplaceRedistributionOfAcquiringOperation — оплата эквайринга;
	// MarketplaceReturnStorageServiceAtThePickupPointFbsItem — краткосрочное размещение возврата FBS;
	// MarketplaceReturnStorageServiceInTheWarehouseFbsItem — долгосрочное размещение возврата FBS;
	// MarketplaceServiceItemDeliveryKGT — доставка КГТ;
	// MarketplaceServiceItemDirectFlowLogistic — логистика;
	// MarketplaceServiceItemReturnFlowLogistic — обратная логистика;
	// MarketplaceServicePremiumCashbackIndividualPoints — услуга продвижения «Бонусы продавца»;
	// OperationMarketplaceWithHoldingForUndeliverableGoods — удержание за недовложение товара;
	// MarketplaceServiceItemDirectFlowLogisticVDC — логистика вРЦ;
	// MarketplaceServiceItemDropoffPPZ — услуга drop-off в пункте приёма заказов;
	// MarketplaceServicePremiumCashback — услуга продвижения Premium;
	// MarketplaceServiceItemRedistributionReturnsPVZ — перевыставление возвратов на пункте выдачи;
	// OperationElectronicServiceStencil — услуга «Трафареты»;
	// OperationElectronicServicesPromotionInSearch — услуга «Продвижение в поиске»;
	// OperationMarketplaceServiceItemElectronicServicesBrandShelf — услуга «Брендовая полка»;
	// OperationSubscriptionPremium — подписка Premium.
	OperationType []string `json:"operation_type"`
	PostingNum    string   `json:"posting_number"`
	// Тип начисления:
	// all — все,
	// orders — заказы,
	// returns — возвраты и отмены,
	// services — сервисные сборы,
	// compensation — компенсация,
	// transferDelivery — стоимость доставки,
	// other — прочее.
	// Некоторые операции могут быть разделены во времени. Например, при приёме возврата от покупателя списывается стоимость товара и возвращается комиссия, а когда товар возвращается на склад, взимается стоимость услуга по обработке возврата
	TransactionType string `json:"transaction_type"`
}

type RequestFilter_transactions_list_date struct {
	// Формат: YYYY-MM-DDTHH:mm:ss.sssZ
	// Пример: 2019-11-25T10:43:06.51
	From time.Time `json:"from"`
	To   time.Time `json:"to"`
}

type Response_transactions_list struct {
	baseResponse
	Result ResponseResult_transactions_list `json:"result"`
}
type ResponseResult_transactions_list struct {
	// Информация об операциях
	Operations []ResponseItem_transactions_list `json:"operations"`
	// Количество страниц. Если 0, страниц больше нет
	PageCount int64 `json:"page_count"`
	// Количество транзакций на всех страницах. Если 0, транзакций больше нет
	RowCount int32 `json:"row_count"`
}
type ResponseItem_transactions_list struct {
	// Идентификатор операции
	OperationId int64 `json:"operation_id"`
	// Тип операции, см коммент из реквеста
	OperationType string `json:"operation_type"`
	// Название типа операции
	OperationTypeName string `json:"operation_type_name"`
	// Дата операции
	OperationDate OzonTime_1 `json:"operation_date"`
	// Тип начисления:
	// all — все,
	// orders — заказы,
	// returns — возвраты и отмены,
	// services — сервисные сборы,
	// compensation — компенсация,
	// transferDelivery — стоимость доставки,
	// other — прочее
	Type string `json:"type"`

	// Итоговая сумма операции
	Amount float64 `json:"amount"`
	// Общая стоимость товаров и возвратов в заданный период
	AccrualsForSale float64 `json:"accruals_for_sale"`
	// Комиссия за продажу или возврат комиссии за продажу
	SaleCommission float64 `json:"sale_commission"`

	// Стоимость доставки для начислений по тарифам, которые действовали до 1 февраля 2021 года, а также начислений для крупногабаритных товаров
	DeliveryCharge float64 `json:"delivery_charge"`
	// Плата за возвраты и отмены для начислений по тарифам, которые действовали до 1 февраля 2021 года, а также начислений для крупногабаритных товаров
	ReturnDeliveryCharge float64 `json:"return_delivery_charge"`

	// Информация о товаре
	Services []ResponseItem_transactions_list_service `json:"services"`

	// Информация о товаре
	Posting ResponseItem_transactions_list_posting `json:"posting"`
	// Информация о товаре
	Items []ResponseItem_transactions_list_item `json:"items"`
}

type ResponseItem_transactions_list_item struct {
	Name string `json:"name"`
	SKU  int64  `json:"sku"`
}

type ResponseItem_transactions_list_posting struct {
	// FBO — доставка со склада Ozon,
	// FBS — доставка со своего склада,
	// CROSSBORDER — доставка из-за рубежа,
	// RFBS — доставка по выбору продавца,
	// FBP — доставка с партнёрских складов Ozon,
	// FBOECONOMY — доставка эконом-товаров со склада Ozon,
	// FBSECONOMY — доставка эконом-товаров со своего склада
	DeliverySchema string `json:"delivery_schema"`
	// Дата принятия отправления в обработку
	OrderDate OzonTime_1 `json:"order_date"`
	// Номер отправления
	PostingNumber string `json:"posting_number"`
	// Идентификатор склада
	WarehouseId int64 `json:"warehouse_id"`
}

type ResponseItem_transactions_list_service struct {
	// Название услуги:
	// MarketplaceNotDeliveredCostItem — возврат невостребованного товара от покупателя на склад.
	// MarketplaceReturnAfterDeliveryCostItem — возврат от покупателя на склад после доставки.
	// MarketplaceDeliveryCostItem — доставка товара до покупателя.
	// MarketplaceSaleReviewsItem — приобретение отзывов на платформе.
	// ItemAdvertisementForSupplierLogistic — доставка товаров на склад Ozon — кросс-докинг.
	// OperationMarketplaceServiceStorage — размещения товаров.
	// MarketplaceMarketingActionCostItem — продвижение товаров.
	// MarketplaceServiceItemInstallment — продвижениe и продажа в рассрочку.
	// MarketplaceServiceItemMarkingItems — обязательная маркировка товаров.
	// MarketplaceServiceItemFlexiblePaymentSchedule — гибкий график выплат.
	// MarketplaceServiceItemReturnFromStock — комплектация товаров для вывоза продавцом.
	// ItemAdvertisementForSupplierLogisticSeller — транспортно-экспедиционные услуги.
	// ItemAgentServiceStarsMembership — вознаграждение за услугу «Звёздные товары».
	// MarketplaceServiceItemDelivToCustomer — последняя миля.
	// MarketplaceServiceItemDirectFlowTrans — магистраль.
	// MarketplaceServiceItemDropoffFF — обработка отправления.
	// MarketplaceServiceItemDropoffPVZ — обработка отправления.
	// MarketplaceServiceItemDropoffSC — обработка отправления.
	// MarketplaceServiceItemFulfillment — сборка заказа.
	// MarketplaceServiceItemPickup — выезд транспортного средства по адресу продавца для забора отправлений — Pick-up.
	// MarketplaceServiceItemReturnAfterDelivToCustomer — обработка возврата.
	// MarketplaceServiceItemReturnFlowTrans — обратная магистраль.
	// MarketplaceServiceItemReturnNotDelivToCustomer — обработка отмен.
	// MarketplaceServiceItemReturnPartGoodsCustomer — обработка невыкупа.
	// MarketplaceRedistributionOfAcquiringOperation — оплата эквайринга.
	// MarketplaceReturnStorageServiceAtThePickupPointFbsItem — краткосрочное размещение возврата FBS.
	// MarketplaceReturnStorageServiceInTheWarehouseFbsItem — долгосрочное размещение возврата FBS.
	// MarketplaceServiceItemDeliveryKGT — доставка крупногабаритного товара (КГТ).
	// MarketplaceServiceItemDirectFlowLogistic — логистика.
	// MarketplaceServiceItemReturnFlowLogistic — обратная логистика.
	// MarketplaceServicePremiumCashbackIndividualPoints — услуга продвижения «Бонусы продавца».
	// MarketplaceServicePremiumPromotion — услуга продвижение Premium, фиксированная комиссия.
	// OperationMarketplaceWithHoldingForUndeliverableGoods — удержание за недовложение товара.
	// MarketplaceServiceItemDropoffPPZ — услуга drop-off в пункте приёма заказов.
	// MarketplaceServiceItemRedistributionReturnsPVZ — перевыставление возвратов на ПВЗ.
	// OperationMarketplaceAgencyFeeAggregator3PLGlobal — тарификация агентской услуги Agregator 3PL Global.
	// MarketplaceServiceItemDirectFlowLogisticVDC — логистика вРЦ
	Name string `json:"name"`
	// Цена
	Price float64 `json:"price"`
}

func NewRequestParams_transactions_list(from, to time.Time, postingnum string, page, limit int64) *RequestParams_transactions_list {
	if limit <= 0 || limit > RequestItems_cap_transactions_list {
		panic("improper limit")
	}
	if page < 1 {
		panic("improper page")
	}

	rp := &RequestParams_transactions_list{
		Filter: RequestFilter_transactions_list{
			PostingNum:      postingnum,
			TransactionType: "all",
		},
		Page:     page,
		PageSize: limit,
	}

	if !from.IsZero() && !to.IsZero() {
		if d := to.Sub(from).Hours() / 24; d > 28 || d < 0 {
			panic("improper from/to (to-from >= 28 days or to before from)")
		}
		rp.Filter.Date.From = from.UTC()
		rp.Filter.Date.To = to.UTC()
	}
	return rp
}

func (cl *OzonClient) GetTransactionsList(params *RequestParams_transactions_list) (*Response_transactions_list, error) {
	req, err := cl.newRequest(http.MethodPost, url_transactions_list, params)
	if err != nil {
		return nil, err
	}
	resp := &Response_transactions_list{}

	response, err := cl.doRequest(req, resp)
	if err != nil {
		return nil, err
	}
	resp.baseResponse = response.base
	//response.pasteBase(&resp.baseResponse)

	return resp, nil
}
