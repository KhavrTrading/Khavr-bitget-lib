package bitget_models

// OrdersHistoryRequest represents the query parameters for retrieving historical orders.
// It only supports data within the past 90 days.
type OrdersHistoryRequest struct {
	// Optional: Order ID. If both OrderId and ClientOid are provided, OrderId takes precedence.
	OrderId string `json:"orderId,omitempty"`
	// Optional: Custom order ID. If both are provided, OrderId takes precedence.
	ClientOid string `json:"clientOid,omitempty"`
	// Optional: Trading pair, e.g. "ETHUSDT".
	Symbol string `json:"symbol,omitempty"`
	// Required: Product type (e.g. "USDT-FUTURES", "COIN-FUTURES", "USDC-FUTURES", etc.).
	ProductType string `json:"productType"`
	// Optional: Requests orders older than the provided order ID.
	IdLessThan string `json:"idLessThan,omitempty"`
	// Optional: Filter by order source. For example: "normal", "market", "profit_market", etc.
	OrderSource string `json:"orderSource,omitempty"`
	// Optional: Start timestamp (in milliseconds). Maximum span supported is one week.
	StartTime string `json:"startTime,omitempty"`
	// Optional: End timestamp (in milliseconds). If not set, default is one week ago.
	EndTime string `json:"endTime,omitempty"`
	// Optional: Number of records to query; default and maximum is 100.
	Limit string `json:"limit,omitempty"`
}

func (r *OrdersHistoryRequest) ToParams() map[string]string {
	params := make(map[string]string)
	if r.OrderId != "" {
		params["orderId"] = r.OrderId
	}
	if r.ClientOid != "" {
		params["clientOid"] = r.ClientOid
	}
	if r.Symbol != "" {
		params["symbol"] = r.Symbol
	}
	if r.ProductType != "" {
		params["productType"] = r.ProductType
	}
	if r.IdLessThan != "" {
		params["idLessThan"] = r.IdLessThan
	}
	if r.OrderSource != "" {
		params["orderSource"] = r.OrderSource
	}
	if r.StartTime != "" {
		params["startTime"] = r.StartTime
	}
	if r.EndTime != "" {
		params["endTime"] = r.EndTime
	}
	if r.Limit != "" {
		params["limit"] = r.Limit
	}
	return params
}

// OrderHistory represents the details of a single historical order.
type OrderHistory struct {
	Symbol                 string `json:"symbol"`                 // Trading pair.
	Size                   string `json:"size"`                   // Order amount.
	OrderId                string `json:"orderId"`                // Order ID.
	ClientOid              string `json:"clientOid"`              // Custom order ID.
	BaseVolume             string `json:"baseVolume"`             // Amount of coins traded.
	Fee                    string `json:"fee"`                    // Transaction fee.
	Price                  string `json:"price"`                  // Order price.
	PriceAvg               string `json:"priceAvg"`               // Average order price.
	Status                 string `json:"status"`                 // Order status (e.g. "filled", "canceled").
	Side                   string `json:"side"`                   // Direction ("buy" or "sell").
	Force                  string `json:"force"`                  // Order expiration type (e.g. "gtc", "ioc", "fok", "post_only").
	TotalProfits           string `json:"totalProfits"`           // Total profit and loss.
	PosSide                string `json:"posSide"`                // Position direction ("long", "short", "net").
	MarginCoin             string `json:"marginCoin"`             // Margin coin.
	QuoteVolume            string `json:"quoteVolume"`            // Trading amount in quote currency.
	Leverage               string `json:"leverage"`               // Leverage.
	MarginMode             string `json:"marginMode"`             // Margin mode ("isolated", "crossed").
	ReduceOnly             string `json:"reduceOnly"`             // Whether the order is reduce-only ("YES", "NO").
	EnterPointSource       string `json:"enterPointSource"`       // Order source (e.g. "WEB", "API", "SYS", "ANDROID", "IOS").
	TradeSide              string `json:"tradeSide"`              // Trade side (e.g. "open", "close", etc.).
	PosMode                string `json:"posMode"`                // Position mode ("one_way_mode", "hedge_mode").
	OrderType              string `json:"orderType"`              // Order type ("limit", "market").
	OrderSource            string `json:"orderSource"`            // Order source category (e.g. "normal", "market", etc.).
	CTime                  string `json:"cTime"`                  // Creation time (ms).
	UTime                  string `json:"uTime"`                  // Last updated time (ms).
	PresetStopSurplusPrice string `json:"presetStopSurplusPrice"` // Take profit trigger price.
	PresetStopLossPrice    string `json:"presetStopLossPrice"`    // Stop loss trigger price.
}

// OrdersHistoryData holds the payload data for historical orders.
type OrdersHistoryData struct {
	EntrustedList []OrderHistory `json:"entrustedList"` // List of historical orders.
	EndId         string         `json:"endId"`         // Last order ID (for pagination).
}

// OrdersHistoryResponse represents the API response for querying historical orders.
type OrdersHistoryResponse struct {
	Code        string            `json:"code"`        // "00000" indicates success.
	Msg         string            `json:"msg"`         // Response message.
	RequestTime int64             `json:"requestTime"` // Request timestamp.
	Data        OrdersHistoryData `json:"data"`        // Data payload.
}
