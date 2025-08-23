package bitget_models

type ClosePositionsSuccess struct {
	OrderId   string `json:"orderId"`   // Order ID.
	ClientOid string `json:"clientOid"` // Customized order ID.
	Symbol    string `json:"symbol"`    // Trading pair.
}
