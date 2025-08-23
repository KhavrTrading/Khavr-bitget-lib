package bitget_models

import "encoding/json"

// OrderFillHistoryRequest represents the query parameters for retrieving historical order fill details.
type OrderFillHistoryRequest struct {
	// Optional: Order ID. Either orderId or clientOid is required; if both are provided, orderId takes precedence.
	OrderId string `json:"orderId,omitempty"`
	// Optional: Trading pair, e.g. "ETHUSDT".
	Symbol string `json:"symbol,omitempty"`
	// Required: Product type, e.g. "USDT-FUTURES", "COIN-FUTURES", etc.
	ProductType string `json:"productType"`
	// Optional: Start timestamp (in milliseconds). Maximum time span supported is a week.
	StartTime string `json:"startTime,omitempty"`
	// Optional: End timestamp (in milliseconds). Default is one week if not set.
	EndTime string `json:"endTime,omitempty"`
	// Optional: Requests data older than the given trade ID.
	IdLessThan string `json:"idLessThan,omitempty"`
	// Optional: Number of records to query. Default: 100; maximum: 100.
	Limit string `json:"limit,omitempty"`
}

func (r *OrderFillHistoryRequest) ToParams() map[string]string {
	params := make(map[string]string)
	if r.OrderId != "" {
		params["orderId"] = r.OrderId
	}
	if r.Symbol != "" {
		params["symbol"] = r.Symbol
	}
	if r.ProductType != "" {
		params["productType"] = r.ProductType
	}
	if r.StartTime != "" {
		params["startTime"] = r.StartTime
	}
	if r.EndTime != "" {
		params["endTime"] = r.EndTime
	}
	if r.IdLessThan != "" {
		params["idLessThan"] = r.IdLessThan
	}
	if r.Limit != "" {
		params["limit"] = r.Limit
	}
	return params
}

// OrderFillHistory represents a single historical order fill.
type OrderFillHistory struct {
	TradeId          string      `json:"tradeId"`          // Transaction ID.
	Symbol           string      `json:"symbol"`           // Trading pair.
	OrderId          string      `json:"orderId"`          // Order ID.
	Price            string      `json:"price"`            // Deal price.
	BaseVolume       string      `json:"baseVolume"`       // Amount of coins traded.
	FeeDetail        []FeeDetail `json:"feeDetail"`        // Transaction fee details.
	Side             string      `json:"side"`             // Direction: "buy" or "sell".
	QuoteVolume      string      `json:"quoteVolume"`      // Trading amount in quote currency.
	Profit           string      `json:"profit"`           // Profit.
	EnterPointSource string      `json:"enterPointSource"` // Order source (e.g. "WEB", "API", "IOS").
	TradeSide        string      `json:"tradeSide"`        // Transaction direction (e.g. "close", "open", etc.).
	PosMode          string      `json:"posMode"`          // Position mode ("one_way_mode" or "hedge_mode").
	TradeScope       string      `json:"tradeScope"`       // Trader tag ("taker" or "maker").
	CTime            string      `json:"cTime"`            // Transaction timestamp (ms).
}

// OrderFillHistoryData holds the data payload for historical order fills.
type OrderFillHistoryData struct {
	FillList []OrderFillHistory `json:"fillList"` // List of order fills.
	EndId    string             `json:"endId"`    // Last order ID in the query.
}

// OrderFillHistoryResponse represents the API response for historical order fill details.
type OrderFillHistoryResponse struct {
	Code        string               `json:"code"`        // "00000" indicates success.
	Msg         string               `json:"msg"`         // Response message.
	RequestTime int64                `json:"requestTime"` // Timestamp of the request.
	Data        OrderFillHistoryData `json:"data"`        // Data payload.
}

// UnmarshalJSON is provided to allow custom unmarshaling if needed.
func (r *OrderFillHistoryResponse) UnmarshalJSON(data []byte) error {
	type Alias OrderFillHistoryResponse
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(r),
	}
	return json.Unmarshal(data, &aux)
}
