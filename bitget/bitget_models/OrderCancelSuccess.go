package bitget_models

type OrderCancelSuccess struct {
	OrderId   string `json:"orderId"`   // Successfully canceled order ID.
	ClientOid string `json:"clientOid"` // Client customized order ID.
}
