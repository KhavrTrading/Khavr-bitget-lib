package bitget_models

// PlaceOrderRequest represents the payload for placing an order.
type PlaceOrderRequest struct {
	Symbol                 string `json:"symbol"`                           // Trading pair, e.g. ETHUSDT
	ProductType            string `json:"productType"`                      // e.g. "USDT-FUTURES"
	MarginMode             string `json:"marginMode"`                       // "isolated" or "crossed"
	MarginCoin             string `json:"marginCoin"`                       // e.g. "USDT"
	Size                   string `json:"size"`                             // Order amount (base coin)
	Price                  string `json:"price,omitempty"`                  // Order price (required for limit orders)
	Side                   string `json:"side"`                             // "buy" or "sell"
	TradeSide              string `json:"tradeSide,omitempty"`              // "open" or "close" (only for hedge-mode)
	OrderType              string `json:"orderType"`                        // "limit" or "market"
	Force                  string `json:"force,omitempty"`                  // Order expiration (e.g., "gtc", "ioc", "fok", "post_only")
	ClientOid              string `json:"clientOid,omitempty"`              // Customized order ID
	ReduceOnly             string `json:"reduceOnly,omitempty"`             // "YES" or "NO" (only for one-way mode)
	PresetStopSurplusPrice string `json:"presetStopSurplusPrice,omitempty"` // Take-profit value (optional)
	PresetStopLossPrice    string `json:"presetStopLossPrice,omitempty"`    // Stop-loss value (optional)
	StpMode                string `json:"stpMode,omitempty"`                // Self trade prevention mode (optional)
}

func (r PlaceOrderRequest) ToParams() map[string]string {
	params := make(map[string]string)
	if r.Symbol != "" {
		params["symbol"] = r.Symbol
	}
	if r.ProductType != "" {
		params["productType"] = r.ProductType
	}
	if r.MarginMode != "" {
		params["marginMode"] = r.MarginMode
	}
	if r.MarginCoin != "" {
		params["marginCoin"] = r.MarginCoin
	}
	if r.Size != "" {
		params["size"] = r.Size
	}
	if r.Price != "" {
		params["price"] = r.Price
	}
	if r.Side != "" {
		params["side"] = r.Side
	}
	if r.TradeSide != "" {
		params["tradeSide"] = r.TradeSide
	}
	if r.OrderType != "" {
		params["orderType"] = r.OrderType
	}
	if r.Force != "" {
		params["force"] = r.Force
	}
	if r.ClientOid != "" {
		params["clientOid"] = r.ClientOid
	}
	if r.ReduceOnly != "" {
		params["reduceOnly"] = r.ReduceOnly
	}
	if r.PresetStopSurplusPrice != "" {
		params["presetStopSurplusPrice"] = r.PresetStopSurplusPrice
	}
	if r.PresetStopLossPrice != "" {
		params["presetStopLossPrice"] = r.PresetStopLossPrice
	}
	if r.StpMode != "" {
		params["stpMode"] = r.StpMode
	}
	return params
}
