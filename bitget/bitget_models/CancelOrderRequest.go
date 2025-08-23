package bitget_models

// CancelOrderRequest represents the payload for canceling an order.
// Either orderId or clientOid must be provided. If both are present, orderId takes precedence.
type CancelOrderRequest struct {
	OrderId     string `json:"orderId,omitempty"`    // Optional: Order ID.
	ClientOid   string `json:"clientOid,omitempty"`  // Optional: Customized order ID.
	Symbol      string `json:"symbol"`               // Required: Trading pair, e.g. "BTCUSDT".
	ProductType string `json:"productType"`          // Required: Product type (e.g. "USDT-FUTURES").
	MarginCoin  string `json:"marginCoin,omitempty"` // Optional: Margin coin (capitalized), e.g. "USDT".
}
