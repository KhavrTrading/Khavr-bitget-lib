package bitget_models

type BatchCancelOrderData struct {
	SuccessList []OrderCancelSuccess `json:"successList"` // Successfully canceled orders.
	FailureList []OrderCancelFailure `json:"failureList"` // Orders that failed to cancel.
}
