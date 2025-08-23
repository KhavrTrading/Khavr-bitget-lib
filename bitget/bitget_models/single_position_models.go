package bitget_models

// SinglePositionRequest holds the query parameters for the single-position endpoint.
type SinglePositionRequest struct {
	// Trading pair, e.g. "btcusdt"
	Symbol string `json:"symbol"`
	// Product type, e.g. "USDT-FUTURES", "COIN-FUTURES", etc.
	ProductType string `json:"productType"`
	// Margin coin, must be capitalized, e.g. "USDT"
	MarginCoin string `json:"marginCoin"`
}

func (r *SinglePositionRequest) ToParams() map[string]string {
	params := make(map[string]string)
	if r.Symbol != "" {
		params["symbol"] = r.Symbol
	}
	if r.ProductType != "" {
		params["productType"] = r.ProductType
	}
	if r.MarginCoin != "" {
		params["marginCoin"] = r.MarginCoin
	}
	return params
}

// SinglePosition represents the details of a single position.
type SinglePosition struct {
	MarginCoin       string `json:"marginCoin"`       // Margin coin
	Symbol           string `json:"symbol"`           // Trading pair name
	HoldSide         string `json:"holdSide"`         // Position direction ("long" or "short")
	OpenDelegateSize string `json:"openDelegateSize"` // Amount to be filled of the current order (base coin)
	MarginSize       string `json:"marginSize"`       // Margin amount (margin coin)
	Available        string `json:"available"`        // Available amount for positions (base coin)
	Locked           string `json:"locked"`           // Frozen amount (base coin)
	Total            string `json:"total"`            // Total amount (available + locked)
	Leverage         string `json:"leverage"`         // Leverage
	AchievedProfits  string `json:"achievedProfits"`  // Realized PnL (exclude fees)
	OpenPriceAvg     string `json:"openPriceAvg"`     // Average entry price
	MarginMode       string `json:"marginMode"`       // Margin mode ("isolated" or "crossed")
	PosMode          string `json:"posMode"`          // Position mode ("one_way_mode" or "hedge_mode")
	UnrealizedPL     string `json:"unrealizedPL"`     // Unrealized profit and loss
	LiquidationPrice string `json:"liquidationPrice"` // Estimated liquidation price (if <=0, no liquidation risk)
	KeepMarginRate   string `json:"keepMarginRate"`   // Tiered maintenance margin rate
	MarkPrice        string `json:"markPrice"`        // Mark price
	BreakEvenPrice   string `json:"breakEvenPrice"`   // Breakeven price
	TotalFee         string `json:"totalFee"`         // Funding fee (accumulated)
	DeductedFee      string `json:"deductedFee"`      // Deducted transaction fees
	MarginRatio      string `json:"marginRatio"`      // Maintenance margin rate (e.g. 0.1 means 10%)
	AssetMode        string `json:"assetMode"`        // Asset mode ("single" or "union")
	UTime            string `json:"uTime"`            // Last updated time (ms)
	AutoMargin       string `json:"autoMargin"`       // Auto Margin ("on" or "off")
	CTime            string `json:"cTime"`            // Creation time (ms)
}

// SinglePositionResponse represents the API response for querying a single position.
type SinglePositionResponse struct {
	Code        string           `json:"code"`        // "00000" indicates success.
	Msg         string           `json:"msg"`         // Response message.
	RequestTime int64            `json:"requestTime"` // Request timestamp.
	Data        []SinglePosition `json:"data"`        // Slice of position information
}
