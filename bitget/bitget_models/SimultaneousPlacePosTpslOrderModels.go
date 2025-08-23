package bitget_models

// PlacePosTpslOrderRequest represents the payload for placing simultaneous stop-profit and stop-loss plan orders.
type PlacePosTpslOrderRequest struct {
	// Required: Margin currency (e.g., "USDT").
	MarginCoin string `json:"marginCoin"`
	// Required: Product type (e.g., "usdt-futures", "coin-futures", etc.).
	ProductType string `json:"productType"`
	// Required: Trading pair, e.g. "BTCUSDT".
	Symbol string `json:"symbol"`
	// Required: Take Profit Trigger price.
	StopSurplusTriggerPrice string `json:"stopSurplusTriggerPrice"`
	// Optional: Take Profit Trigger type.
	// Valid values: "fill_price" (market price), "mark_price" (mark price).
	StopSurplusTriggerType string `json:"stopSurplusTriggerType,omitempty"`
	// Optional: Take Profit Execution price.
	// If set to 0 or not provided, market price execution is assumed.
	StopSurplusExecutePrice string `json:"stopSurplusExecutePrice,omitempty"`
	// Required: Stop Loss Trigger price.
	StopLossTriggerPrice string `json:"stopLossTriggerPrice"`
	// Optional: Stop Loss Trigger type.
	// Valid values: "fill_price" (market price), "mark_price" (mark price).
	StopLossTriggerType string `json:"stopLossTriggerType,omitempty"`
	// Optional: Stop Loss Execution price.
	// If set to 0 or not provided, market price execution is assumed.
	StopLossExecutePrice string `json:"stopLossExecutePrice,omitempty"`
	// Required: Position side.
	// For two-way positions use "long" or "short"; for one-way positions use "buy" or "sell".
	HoldSide string `json:"holdSide"`
	// Optional: STP Mode.
	// Valid values: "none", "cancel_taker", "cancel_maker", "cancel_both".
	StpMode string `json:"stpMode,omitempty"`
}

// PlacePosTpslOrderData represents an individual order's response data.
type PlacePosTpslOrderData struct {
	// Order ID.
	OrderId string `json:"orderId"`
}

// PlacePosTpslOrderResponse represents the API response for placing simultaneous stop-profit and stop-loss plan orders.
type PlacePosTpslOrderResponse struct {
	// "00000" indicates success.
	Code string `json:"code"`
	// Response message.
	Msg string `json:"msg"`
	// Timestamp of the request.
	RequestTime int64 `json:"requestTime"`
	// List of order responses.
	Data []PlacePosTpslOrderData `json:"data"`
}
