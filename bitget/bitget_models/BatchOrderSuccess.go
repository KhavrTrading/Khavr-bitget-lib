package bitget_models

type BatchOrderSuccess struct {
	OrderId   string `json:"orderId"`   // Exchange-assigned order ID.
	ClientOid string `json:"clientOid"` // Custom order ID from the request.
}
