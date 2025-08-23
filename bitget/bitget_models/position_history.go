package bitget_models

// PositionHistoryRequest represents the query parameters for retrieving historical positions.
type PositionHistoryRequest struct {
	Symbol      string `json:"symbol,omitempty"`
	ProductType string `json:"productType,omitempty"`
	IdLessThan  string `json:"idLessThan,omitempty"`
	StartTime   string `json:"startTime,omitempty"` // Timestamp in ms
	EndTime     string `json:"endTime,omitempty"`   // Timestamp in ms
	Limit       string `json:"limit,omitempty"`
}

// ToParams converts the PositionHistoryRequest to query parameters.
func (r *PositionHistoryRequest) ToParams() map[string]string {
	params := make(map[string]string)
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

// PositionHistoryItem represents a single historical position.
type PositionHistoryItem struct {
	PositionId    string `json:"positionId"`
	MarginCoin    string `json:"marginCoin"`
	Symbol        string `json:"symbol"`
	HoldSide      string `json:"holdSide"` // long / short
	OpenAvgPrice  string `json:"openAvgPrice"`
	CloseAvgPrice string `json:"closeAvgPrice"`
	MarginMode    string `json:"marginMode"` // isolated / crossed
	OpenTotalPos  string `json:"openTotalPos"`
	CloseTotalPos string `json:"closeTotalPos"`
	PNL           string `json:"pnl"`
	NetProfit     string `json:"netProfit"`
	TotalFunding  string `json:"totalFunding"`
	OpenFee       string `json:"openFee"`
	CloseFee      string `json:"closeFee"`
	CTime         string `json:"cTime"` // Timestamp in ms
	UTime         string `json:"uTime"` // Timestamp in ms
}

// PositionHistoryData holds the response data and pagination ID.
type PositionHistoryData struct {
	List  []PositionHistoryItem `json:"list"`
	EndId string                `json:"endId"`
}

// PositionHistoryResponse represents the full API response for historical positions.
type PositionHistoryResponse struct {
	Code        string              `json:"code"`
	Msg         string              `json:"msg"`
	RequestTime int64               `json:"requestTime"`
	Data        PositionHistoryData `json:"data"`
}
