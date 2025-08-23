package bitget_models

// PlaceTpslOrderRequest represents the payload for placing a stop-profit and stop-loss plan order.
type PlaceTpslOrderRequest struct {
	// Required: Margin currency (capitalized, e.g., "USDT").
	MarginCoin string `json:"marginCoin"`
	// Required: Product type (e.g., "usdt-futures", "coin-futures", etc.).
	ProductType string `json:"productType"`
	// Required: Trading pair, e.g., "ethusdt".
	Symbol string `json:"symbol"`
	// Required: Order plan type.
	// Valid values:
	// profit_plan: take profit plan;
	// loss_plan: stop loss plan;
	// moving_plan: trailing stop;
	// pos_profit: position take profit;
	// pos_loss: position stop loss.
	PlanType string `json:"planType"`
	// Required: Trigger price.
	TriggerPrice string `json:"triggerPrice"`
	// Optional: Trigger type.
	// fill_price: market price;
	// mark_price: mark price.
	TriggerType string `json:"triggerType,omitempty"`
	// Optional: Execution price.
	// If 0 or not provided, it means market price execution.
	ExecutePrice string `json:"executePrice,omitempty"`
	// Required: Position side.
	// For two-way positions: "long" or "short".
	HoldSide string `json:"holdSide"`
	// Required when planType is profit_plan, loss_plan or moving_plan.
	// Represents order quantity (base coin).
	Size string `json:"size,omitempty"`
	// Optional: Callback range, required only when planType is moving_plan.
	RangeRate string `json:"rangeRate,omitempty"`
	// Optional: Customized order ID.
	ClientOid string `json:"clientOid,omitempty"`
	// Optional: STP Mode.
	// Valid values: none, cancel_taker, cancel_maker, cancel_both.
	StpMode string `json:"stpMode,omitempty"`
}

// PlaceTpslOrderData holds the order information returned in a successful response.
type PlaceTpslOrderData struct {
	// Order ID.
	OrderId string `json:"orderId"`
	// Customized order ID.
	ClientOid string `json:"clientOid"`
}

// PlaceTpslOrderResponse represents the API response for placing a stop-profit and stop-loss plan order.
type PlaceTpslOrderResponse struct {
	// "00000" indicates success.
	Code string `json:"code"`
	// Response message.
	Msg string `json:"msg"`
	// Timestamp of the request.
	RequestTime int64 `json:"requestTime"`
	// Order details.
	Data PlaceTpslOrderData `json:"data"`
}
