package bitget_models

type PlaceOrderResponse struct {
	Code        string    `json:"code"`        // "00000" indicates success.
	Msg         string    `json:"msg"`         // Message, usually "success" on success.
	RequestTime int64     `json:"requestTime"` // Timestamp of the request.
	Data        OrderData `json:"data"`        // Data payload containing order details.
}
