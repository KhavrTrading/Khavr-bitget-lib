package bitget_models

// OrderCancelID represents an individual order identifier.
type OrderCancelID struct {
	// Order ID. If both orderId and clientOid are provided, orderId takes precedence.
	OrderId string `json:"orderId,omitempty"`
	// Customized order ID.
	ClientOid string `json:"clientOid,omitempty"`
}
