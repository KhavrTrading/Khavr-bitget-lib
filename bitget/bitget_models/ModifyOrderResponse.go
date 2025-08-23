package bitget_models

// ModifyOrderResponse represents the response returned after modifying an order.
type ModifyOrderResponse struct {
	Code        string          `json:"code"`        // "00000" indicates success.
	Msg         string          `json:"msg"`         // Response message.
	RequestTime int64           `json:"requestTime"` // Timestamp of the request.
	Data        ModifyOrderData `json:"data"`        // Order details.
}
