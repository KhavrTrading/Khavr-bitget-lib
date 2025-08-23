package bitget_models

// BatchCancelOrderRequest represents the payload for a batch cancel order request.
type BatchCancelOrderRequest struct {
	// Trading pair, e.g. "BTCUSDT". It is required when orderIdList is set.
	Symbol string `json:"symbol,omitempty"`
	// Product type (e.g. "USDT-FUTURES", "COIN-FUTURES", etc.).
	ProductType string `json:"productType"`
	// Margin coin, must be capitalized (e.g. "USDT").
	MarginCoin string `json:"marginCoin,omitempty"`
	// List of orders to cancel (maximum length: 50). Each order can be identified by orderId or clientOid.
	OrderIdList []OrderCancelID `json:"orderIdList,omitempty"`
}
