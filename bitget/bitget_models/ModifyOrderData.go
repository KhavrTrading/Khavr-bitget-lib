package bitget_models

// ModifyOrderData contains the order details returned by the modify order request.
type ModifyOrderData struct {
	OrderId   string `json:"orderId"`   // Order ID of the modified order.
	ClientOid string `json:"clientOid"` // Customized order ID.
}
