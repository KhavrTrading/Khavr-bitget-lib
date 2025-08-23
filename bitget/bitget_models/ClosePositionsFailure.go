package bitget_models

type ClosePositionsFailure struct {
	OrderId   string `json:"orderId"`   // Order ID.
	ClientOid string `json:"clientOid"` // Customized order ID.
	Symbol    string `json:"symbol"`    // Trading pair.
	ErrorMsg  string `json:"errorMsg"`  // Failure reason.
	ErrorCode string `json:"errorCode"` // Failure code.
}
