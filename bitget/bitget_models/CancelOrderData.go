package bitget_models

// CancelOrderData holds the details of the canceled order.
type CancelOrderData struct {
	OrderId   string `json:"orderId"`   // Order ID.
	ClientOid string `json:"clientOid"` // Client customized order ID.
}
