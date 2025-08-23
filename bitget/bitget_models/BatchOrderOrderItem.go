package bitget_models

type BatchOrderOrderItem struct {
	Size                   string `json:"size"`                             // Order amount
	Price                  string `json:"price,omitempty"`                  // Order price (required for limit orders)
	Side                   string `json:"side"`                             // "buy" or "sell"
	TradeSide              string `json:"tradeSide,omitempty"`              // "open" or "close" (only required in hedge-mode)
	OrderType              string `json:"orderType"`                        // "limit" or "market"
	Force                  string `json:"force,omitempty"`                  // Order expiration: "gtc", "ioc", "fok", "post_only"
	ClientOid              string `json:"clientOid,omitempty"`              // Customized order ID
	ReduceOnly             string `json:"reduceOnly,omitempty"`             // "YES" or "NO" (applicable only in one-way mode)
	PresetStopSurplusPrice string `json:"presetStopSurplusPrice,omitempty"` // Take-profit value (optional)
	PresetStopLossPrice    string `json:"presetStopLossPrice,omitempty"`    // Stop-loss value (optional)
	StpMode                string `json:"stpMode,omitempty"`                // STP mode (optional)
}
