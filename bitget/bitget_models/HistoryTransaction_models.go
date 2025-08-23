package bitget_models

// HistoryTransactionRequest holds the query parameters for the fills-history endpoint.
type HistoryTransactionRequest struct {
	// Trading pair, e.g. "BTCUSDT"
	Symbol string `json:"symbol"`
	// Product type, e.g. "usdt-futures", "coin-futures", etc.
	ProductType string `json:"productType"`
	// Optional: Number of queries (default 500, maximum 1000)
	Limit string `json:"limit,omitempty"`
	// Optional: Request older data before this ID (should be the endId of the previous page)
	IdLessThan string `json:"idLessThan,omitempty"`
	// Optional: Start timestamp (milliseconds)
	StartTime string `json:"startTime,omitempty"`
	// Optional: End timestamp (milliseconds)
	EndTime string `json:"endTime,omitempty"`
}

// ToParams converts the HistoryTransactionRequest into a map of query parameters.
func (r *HistoryTransactionRequest) ToParams() map[string]string {
	params := make(map[string]string)
	if r.Symbol != "" {
		params["symbol"] = r.Symbol
	}
	if r.ProductType != "" {
		params["productType"] = r.ProductType
	}
	if r.Limit != "" {
		params["limit"] = r.Limit
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
	return params
}

// ToQueryString converts the request parameters to a URL query string.
func (r *HistoryTransactionRequest) ToQueryString() string {
	params := r.ToParams()
	qs := ""
	first := true
	for k, v := range params {
		if first {
			qs += k + "=" + v
			first = false
		} else {
			qs += "&" + k + "=" + v
		}
	}
	return qs
}

// HistoryTransactionData represents a single transaction record.
type HistoryTransactionData struct {
	TradeId string `json:"tradeId"` // Trade ID in descending order.
	Price   string `json:"price"`   // Price.
	Size    string `json:"size"`    // Amount (base coin).
	Side    string `json:"side"`    // Trading direction: "Sell" or "Buy".
	Ts      string `json:"ts"`      // Timestamp (Unix in ms).
	Symbol  string `json:"symbol"`  // Trading pair.
}

// HistoryTransactionResponse represents the API response for fills-history.
type HistoryTransactionResponse struct {
	Code        string                   `json:"code"`        // "00000" indicates success.
	Msg         string                   `json:"msg"`         // Response message.
	RequestTime int64                    `json:"requestTime"` // Request timestamp.
	Data        []HistoryTransactionData `json:"data"`        // List of transaction records.
}
