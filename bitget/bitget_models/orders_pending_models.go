package bitget_models

// OrdersPendingRequest represents the query parameters for retrieving all pending orders.
type OrdersPendingRequest struct {
	// Optional: Order ID. If both OrderId and ClientOid are provided, OrderId takes precedence.
	OrderId string `json:"orderId,omitempty"`
	// Optional: Customized order ID. If both are provided, OrderId takes precedence.
	ClientOid string `json:"clientOid,omitempty"`
	// Optional: Trading pair, e.g. "ETHUSDT".
	Symbol string `json:"symbol,omitempty"`
	// Required: Product type, e.g. "USDT-FUTURES", "COIN-FUTURES", etc.
	ProductType string `json:"productType"`
	// Optional: Order status. If not specified, all pending orders (status "live") are returned.
	Status string `json:"status,omitempty"`
	// Optional: Requests data older than the provided order ID (endId of the previous query).
	IdLessThan string `json:"idLessThan,omitempty"`
	// Optional: Start timestamp (in milliseconds). Maximum time span supported is three months.
	StartTime string `json:"startTime,omitempty"`
	// Optional: End timestamp (in milliseconds). If not set, the default is three months ago.
	EndTime string `json:"endTime,omitempty"`
	// Optional: Number of records to query. Default is 100; maximum is 100.
	Limit string `json:"limit,omitempty"`
}

func (r *OrdersPendingRequest) ToParams() map[string]string {
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
	if r.Status != "" {
		params["status"] = r.Status
	}
	if r.IdLessThan != "" {
		params["idLessThan"] = r.IdLessThan
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

// OrderPending represents a single pending order's details.
type OrderPending struct {
	Symbol                        string `json:"symbol"`                        // Trading pair.
	Size                          string `json:"size"`                          // Order amount.
	OrderId                       string `json:"orderId"`                       // Order ID.
	ClientOid                     string `json:"clientOid"`                     // Customized order ID.
	BaseVolume                    string `json:"baseVolume"`                    // Amount of coins traded.
	Fee                           string `json:"fee"`                           // Transaction fee.
	Price                         string `json:"price"`                         // Order price.
	PriceAvg                      string `json:"priceAvg"`                      // Average order price (empty when status is live).
	Status                        string `json:"status"`                        // Order status ("live", "partially_filled").
	Side                          string `json:"side"`                          // Order direction ("buy" or "sell").
	Force                         string `json:"force"`                         // Order expiration type ("gtc", "ioc", "fok", "post_only").
	TotalProfits                  string `json:"totalProfits"`                  // Total profit and loss (empty when status is live).
	PosSide                       string `json:"posSide"`                       // Position direction (e.g. "long", "short", "net").
	MarginCoin                    string `json:"marginCoin"`                    // Margin coin.
	QuoteVolume                   string `json:"quoteVolume"`                   // Trading amount in quote currency.
	Leverage                      string `json:"leverage"`                      // Leverage.
	MarginMode                    string `json:"marginMode"`                    // Margin mode ("isolated" or "crossed").
	ReduceOnly                    string `json:"reduceOnly"`                    // Whether order is reduce-only ("YES" or "NO").
	EnterPointSource              string `json:"enterPointSource"`              // Order source ("WEB", "API", "SYS", "ANDROID", "IOS").
	TradeSide                     string `json:"tradeSide"`                     // Trade side ("open", "close", etc.).
	PosMode                       string `json:"posMode"`                       // Position mode ("one_way_mode" or "hedge_mode").
	OrderType                     string `json:"orderType"`                     // Order type ("limit" or "market").
	OrderSource                   string `json:"orderSource"`                   // Order source category (e.g., "normal", "market", "profit_market", etc.).
	CTime                         string `json:"cTime"`                         // Creation time (timestamp in ms).
	UTime                         string `json:"uTime"`                         // Last update time (timestamp in ms).
	PresetStopSurplusPrice        string `json:"presetStopSurplusPrice"`        // Take profit trigger price.
	PresetStopSurplusType         string `json:"presetStopSurplusType"`         // Take profit trigger type ("fill_price", "mark_price").
	PresetStopSurplusExecutePrice string `json:"presetStopSurplusExecutePrice"` // Take profit execution price.
	PresetStopLossPrice           string `json:"presetStopLossPrice"`           // Stop loss trigger price.
	PresetStopLossType            string `json:"presetStopLossType"`            // Stop loss trigger type ("fill_price", "mark_price").
	PresetStopLossExecutePrice    string `json:"presetStopLossExecutePrice"`    // Stop loss execution price.
}

// OrdersPendingData holds the data payload for the pending orders response.
type OrdersPendingData struct {
	EntrustedList []OrderPending `json:"entrustedList"` // List of pending orders.
	EndId         string         `json:"endId"`         // The final order ID used for pagination.
}

// OrdersPendingResponse represents the API response for querying pending orders.
type OrdersPendingResponse struct {
	Code        string            `json:"code"`        // "00000" indicates success.
	Data        OrdersPendingData `json:"data"`        // Data payload.
	Msg         string            `json:"msg"`         // Response message.
	RequestTime int64             `json:"requestTime"` // Request timestamp.
}
