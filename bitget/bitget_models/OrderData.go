package bitget_models

type OrderData struct {
	OrderId   string `json:"orderId"`   // The exchange-assigned order ID.
	ClientOid string `json:"clientOid"` // The client order ID sent in the request.
}
