package bitget_models

type OrderCancelFailure struct {
	OrderId   string `json:"orderId"`   // Order ID that failed to cancel.
	ClientOid string `json:"clientOid"` // Client customized order ID.
	ErrorMsg  string `json:"errorMsg"`  // Failure reason.
	ErrorCode string `json:"errorCode"` // Failure code.
}
