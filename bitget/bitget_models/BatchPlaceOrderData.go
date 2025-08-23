package bitget_models

type BatchPlaceOrderData struct {
	SuccessList []BatchOrderSuccess `json:"successList"` // List of successful orders.
	FailureList []BatchOrderFailure `json:"failureList"` // List of failed orders.
}
