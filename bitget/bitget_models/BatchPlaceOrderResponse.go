package bitget_models

type BatchPlaceOrderResponse struct {
	Code        string              `json:"code"`        // "00000" indicates success.
	Data        BatchPlaceOrderData `json:"data"`        // Contains successList and failureList.
	Msg         string              `json:"msg"`         // Response message.
	RequestTime int64               `json:"requestTime"` // Timestamp of the request.
}
