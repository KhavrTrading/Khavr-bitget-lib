package bitget_models

// ClosePositionsRequest represents the payload for closing positions at market price.
type ClosePositionsRequest struct {
	// Trading pair, e.g. "BTCUSDT". In one-way mode this field can be omitted.
	Symbol string `json:"symbol,omitempty"`
	// Product type, e.g. "USDT-FUTURES", "COIN-FUTURES", etc.
	ProductType string `json:"productType"`
	// Position direction: "long" for long positions, "short" for short positions.
	// In one-way mode this field is ignored; if left blank in hedge-mode, all positions will be closed.
	HoldSide string `json:"holdSide,omitempty"`
}
