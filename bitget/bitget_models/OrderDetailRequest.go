package bitget_models

// OrderDetailRequest represents the query parameters for getting order details.
// Either OrderId or ClientOid must be provided. If both are provided, OrderId takes precedence.
type OrderDetailRequest struct {
	Symbol      string `json:"symbol"`              // Required: Trading pair, must be capitalized (e.g. "ETHUSDT")
	ProductType string `json:"productType"`         // Required: Product type (e.g. "USDT-FUTURES")
	OrderId     string `json:"orderId,omitempty"`   // Optional: Order ID
	ClientOid   string `json:"clientOid,omitempty"` // Optional: Custom order ID
}

func (r *OrderDetailRequest) ToParams() map[string]string {
	params := make(map[string]string)
	if r.Symbol != "" {
		params["symbol"] = r.Symbol
	}
	if r.ProductType != "" {
		params["productType"] = r.ProductType
	}
	if r.OrderId != "" {
		params["orderId"] = r.OrderId
	}
	if r.ClientOid != "" {
		params["clientOid"] = r.ClientOid
	}
	return params
}
