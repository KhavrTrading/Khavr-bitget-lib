package bitget_models

import "encoding/json"

// OrderFillsRequest represents the query parameters for retrieving order fill details.
type OrderFillsRequest struct {
	OrderId     string `json:"orderId,omitempty"`    // Optional: Order ID.
	Symbol      string `json:"symbol,omitempty"`     // Optional: Trading pair, e.g. "ETHUSDT".
	ProductType string `json:"productType"`          // Required: Product type, e.g. "USDT-FUTURES".
	IdLessThan  string `json:"idLessThan,omitempty"` // Optional: Request fills older than this trade ID.
	StartTime   string `json:"startTime,omitempty"`  // Optional: Start time (timestamp in ms).
	EndTime     string `json:"endTime,omitempty"`    // Optional: End time (timestamp in ms).
	Limit       string `json:"limit,omitempty"`      // Optional: Number of records to query; default is 100, max 100.
}

func (r *OrderFillsRequest) ToParams() map[string]string {
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

// FeeDetail represents fee details for a fill.
type FeeDetail struct {
	Deduction         string `json:"deduction"`         // Whether or not fee deduction applies.
	FeeCoin           string `json:"feeCoin"`           // Crypto ticker of the fee.
	TotalDeductionFee string `json:"totalDeductionFee"` // Total fee deduction.
	TotalFee          string `json:"totalFee"`          // Total fee.
}

// OrderFill represents the details of a single fill (transaction).
type OrderFill struct {
	TradeId          string      `json:"tradeId"`          // Transaction ID.
	Symbol           string      `json:"symbol"`           // Trading pair.
	OrderId          string      `json:"orderId"`          // Order number.
	Price            string      `json:"price"`            // Order price.
	BaseVolume       string      `json:"baseVolume"`       // Amount of coins traded.
	FeeDetail        []FeeDetail `json:"feeDetail"`        // Transaction fee details.
	Side             string      `json:"side"`             // Transaction type (buy or sell).
	QuoteVolume      string      `json:"quoteVolume"`      // Trading amount in quote currency.
	Profit           string      `json:"profit"`           // Profit.
	EnterPointSource string      `json:"enterPointSource"` // Order source (e.g. "api").
	TradeSide        string      `json:"tradeSide"`        // Direction (e.g. "close", "open", etc.).
	PosMode          string      `json:"posMode"`          // Position mode ("one_way_mode" or "hedge_mode").
	TradeScope       string      `json:"tradeScope"`       // Trader tag ("taker" or "maker").
	CTime            string      `json:"cTime"`            // Transaction timestamp (ms).
}

// OrderFillsData holds the response data for order fills.
type OrderFillsData struct {
	FillList []OrderFill `json:"fillList"` // List of order fills.
	EndId    string      `json:"endId"`    // Final trade ID in the returned range.
}

// OrderFillsResponse represents the full API response for order fill details.
type OrderFillsResponse struct {
	Code        string         `json:"code"`        // "00000" indicates success.
	Data        OrderFillsData `json:"data"`        // Order fills data.
	Msg         string         `json:"msg"`         // Response message.
	RequestTime int64          `json:"requestTime"` // Timestamp of the request.
}

// UnmarshalJSON provides a custom unmarshal for OrderFillsResponse if needed.
func (r *OrderFillsResponse) UnmarshalJSON(data []byte) error {
	type Alias OrderFillsResponse
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(r),
	}
	return json.Unmarshal(data, &aux)
}
