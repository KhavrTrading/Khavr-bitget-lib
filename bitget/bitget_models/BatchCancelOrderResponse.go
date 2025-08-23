package bitget_models

type BatchCancelOrderResponse struct {
	Code        string               `json:"code"`        // "00000" indicates success.
	Msg         string               `json:"msg"`         // Response message.
	RequestTime int64                `json:"requestTime"` // Timestamp of the request.
	Data        BatchCancelOrderData `json:"data"`        // Contains lists of successes and failures.
}
