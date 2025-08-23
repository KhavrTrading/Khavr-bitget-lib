package bitget_models

// SetLeverageRequest represents the request parameters for adjusting leverage.
type SetLeverageRequest struct {
	Symbol      string `json:"symbol"`                // Trading pair, e.g. "btcusdt"
	ProductType string `json:"productType"`           // Product type, e.g. "USDT-FUTURES"
	MarginCoin  string `json:"marginCoin"`            // Margin coin (must be capitalized in response, but request can be lowercase), e.g. "usdt"
	Leverage    string `json:"leverage"`              // Leverage value, e.g. "20"
	HoldSide    string `json:"holdSide,omitempty"`    // Optional: Position direction ("long" or "short")
}

// LeverageData holds the leverage details returned by the API.
type LeverageData struct {
	Symbol              string `json:"symbol"`              // Trading pair name.
	MarginCoin          string `json:"marginCoin"`          // Margin coin.
	LongLeverage        string `json:"longLeverage"`        // Leverage for long positions.
	ShortLeverage       string `json:"shortLeverage"`       // Leverage for short positions.
	CrossMarginLeverage string `json:"crossMarginLeverage"` // Leverage for cross margin mode.
	MarginMode          string `json:"marginMode"`          // Margin mode ("isolated" or "crossed").
}

// SetLeverageResponse represents the API response for setting leverage.
type SetLeverageResponse struct {
	Code        string       `json:"code"`        // "00000" indicates success.
	Data        LeverageData `json:"data"`        // Leverage details.
	Msg         string       `json:"msg"`         // Response message.
	RequestTime int64        `json:"requestTime"` // Timestamp of the request.
}
