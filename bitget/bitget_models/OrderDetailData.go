package bitget_models

// OrderDetailData holds the detailed information of an order.
type OrderDetailData struct {
	Symbol                 string `json:"symbol"`                 // Trading pair
	Size                   string `json:"size"`                   // Amount
	OrderId                string `json:"orderId"`                // Order ID
	ClientOid              string `json:"clientOid"`              // Custom order ID
	BaseVolume             string `json:"baseVolume"`             // Amount of coins traded
	PriceAvg               string `json:"priceAvg"`               // Average price
	Fee                    string `json:"fee"`                    // Transaction fee
	Price                  string `json:"price"`                  // Order price
	State                  string `json:"state"`                  // Order status (live, partially_filled, filled, canceled)
	Side                   string `json:"side"`                   // Direction (buy or sell)
	Force                  string `json:"force"`                  // Order expiration date (gtc, ioc, fok, post only)
	TotalProfits           string `json:"totalProfits"`           // Total PnL
	PosSide                string `json:"posSide"`                // Position direction (long, short, net)
	MarginCoin             string `json:"marginCoin"`             // Margin coin
	PresetStopSurplusPrice string `json:"presetStopSurplusPrice"` // Set take-profit
	PresetStopLossPrice    string `json:"presetStopLossPrice"`    // Set stop-loss
	QuoteVolume            string `json:"quoteVolume"`            // Trading amount in quoting coin
	OrderType              string `json:"orderType"`              // Order type (limit or market)
	Leverage               string `json:"leverage"`               // Leverage
	MarginMode             string `json:"marginMode"`             // Margin mode (isolated, crossed)
	ReduceOnly             string `json:"reduceOnly"`             // Whether it's a reduce-only order (YES, NO)
	EnterPointSource       string `json:"enterPointSource"`       // Order source (WEB, API, SYS, ANDROID, IOS)
	TradeSide              string `json:"tradeSide"`              // Direction in open/close mode (open, close, etc.)
	PosMode                string `json:"posMode"`                // Position mode (one_way_mode, hedge_mode)
	OrderSource            string `json:"orderSource"`            // Order source (normal, market, etc.)
	CancelReason           string `json:"cancelReason"`           // Cancel reason (normal_cancel, stp_cancel)
	CTime                  string `json:"cTime"`                  // Creation time (ms)
	UTime                  string `json:"uTime"`                  // Update time (ms)
}
