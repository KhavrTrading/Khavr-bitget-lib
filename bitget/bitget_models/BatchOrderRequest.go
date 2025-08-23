package bitget_models

// BatchOrderRequest represents the payload for batch placing orders.
type BatchOrderRequest struct {
	Symbol      string                `json:"symbol"`      // Trading pair, e.g. BTCUSDT
	ProductType string                `json:"productType"` // e.g. "USDT-FUTURES"
	MarginMode  string                `json:"marginMode"`  // "isolated" or "crossed"
	MarginCoin  string                `json:"marginCoin"`  // e.g. "USDT"
	OrderList   []BatchOrderOrderItem `json:"orderList"`   // List of orders (max length: 50)
}
