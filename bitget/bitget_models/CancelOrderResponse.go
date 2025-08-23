package bitget_models

// CancelOrderResponse represents the response from a cancel order request.
type CancelOrderResponse struct {
	Code        string          `json:"code"`        // "00000" indicates success.
	Msg         string          `json:"msg"`         // Response message.
	RequestTime int64           `json:"requestTime"` // Timestamp of the request.
	Data        CancelOrderData `json:"data"`        // Data containing order details.
}
