package bitget_models

type ClosePositionsData struct {
	SuccessList []ClosePositionsSuccess `json:"successList"` // Successfully closed orders.
	FailureList []ClosePositionsFailure `json:"failureList"` // Orders that failed to close.
}
