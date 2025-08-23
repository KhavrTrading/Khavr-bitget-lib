package bitget_models

type BatchOrderFailure struct {
	OrderId   string `json:"orderId"`   // Order ID (if available).
	ClientOid string `json:"clientOid"` // Custom order ID.
	ErrorMsg  string `json:"errorMsg"`  // Failure reason.
	ErrorCode string `json:"errorCode"` // Failure code.
}
