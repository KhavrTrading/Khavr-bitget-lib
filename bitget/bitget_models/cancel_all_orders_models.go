package bitget_models

// CancelAllOrdersRequest represents the payload for canceling all orders.
type CancelAllOrdersRequest struct {
	// Required: Product type (e.g., "USDT-FUTURES", "COIN-FUTURES", etc.)
	ProductType string `json:"productType"`
	// Optional: Margin coin, must be capitalized (e.g., "USDT").
	MarginCoin string `json:"marginCoin,omitempty"`
	// Optional: Request time as a Unix millisecond timestamp.
	RequestTime string `json:"requestTime,omitempty"`
	// Optional: Valid window period as a Unix millisecond timestamp.
	ReceiveWindow string `json:"receiveWindow,omitempty"`
}

// CancelAllOrdersSuccess represents a successfully canceled order.
type CancelAllOrdersSuccess struct {
	OrderId   string `json:"orderId"`   // Order ID.
	ClientOid string `json:"clientOid"` // Customized order ID.
}

// CancelAllOrdersFailure represents an order that failed to cancel.
type CancelAllOrdersFailure struct {
	OrderId   string `json:"orderId"`   // Order ID.
	ClientOid string `json:"clientOid"` // Customized order ID.
	ErrorMsg  string `json:"errorMsg"`  // Failure reason.
	ErrorCode string `json:"errorCode"` // Error code.
}

// CancelAllOrdersData holds the response data.
type CancelAllOrdersData struct {
	SuccessList []CancelAllOrdersSuccess `json:"successList"` // Successfully canceled orders.
	FailureList []CancelAllOrdersFailure `json:"failureList"` // Orders that failed to cancel.
}

// CancelAllOrdersResponse represents the API response.
type CancelAllOrdersResponse struct {
	Code        string              `json:"code"`        // "00000" indicates success.
	Msg         string              `json:"msg"`         // Response message.
	RequestTime int64               `json:"requestTime"` // Timestamp of the request.
	Data        CancelAllOrdersData `json:"data"`        // Payload data.
}
