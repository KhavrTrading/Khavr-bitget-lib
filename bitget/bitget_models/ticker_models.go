package bitget_models

// TickerRequest holds the query parameters for the ticker endpoint.
type TickerRequest struct {
	// Trading pair, e.g. "ETHUSDM24"
	Symbol string `json:"symbol"`
	// Product type, e.g. "COIN-FUTURES", "USDT-FUTURES", etc.
	ProductType string `json:"productType"`
}

// ToParams converts TickerRequest to a map of query parameters.
func (r *TickerRequest) ToParams() map[string]string {
	params := make(map[string]string)
	if r.Symbol != "" {
		params["symbol"] = r.Symbol
	}
	if r.ProductType != "" {
		params["productType"] = r.ProductType
	}
	return params
}

// TickerResponse represents the API response for the ticker data.
type TickerResponse struct {
	Code        string       `json:"code"`        // "00000" indicates success.
	Msg         string       `json:"msg"`         // Response message.
	RequestTime int64        `json:"requestTime"` // Request timestamp.
	Data        []TickerData `json:"data"`        // Array of ticker data.
}

// TickerData holds the ticker details.
type TickerData struct {
	Symbol            string `json:"symbol"`            // Trading pair name.
	LastPr            string `json:"lastPr"`            // Last price.
	AskPr             string `json:"askPr"`             // Ask price.
	BidPr             string `json:"bidPr"`             // Bid price.
	BidSz             string `json:"bidSz"`             // Buying amount.
	AskSz             string `json:"askSz"`             // Selling amount.
	High24h           string `json:"high24h"`           // 24h high.
	Low24h            string `json:"low24h"`            // 24h low.
	Ts                string `json:"ts"`                // Current timestamp in milliseconds.
	Change24h         string `json:"change24h"`         // Price change in 24 hours.
	BaseVolume        string `json:"baseVolume"`        // Trading volume of the coin.
	QuoteVolume       string `json:"quoteVolume"`       // Trading volume of quote currency.
	UsdtVolume        string `json:"usdtVolume"`        // Trading volume in USDT.
	OpenUtc           string `json:"openUtc"`           // UTC opening price.
	ChangeUtc24h      string `json:"changeUtc24h"`      // UTC 24h price change.
	IndexPrice        string `json:"indexPrice"`        // Index price.
	FundingRate       string `json:"fundingRate"`       // Funding rate.
	HoldingAmount     string `json:"holdingAmount"`     // Current holding positions.
	DeliveryStartTime string `json:"deliveryStartTime"` // Delivery start time (for delivery contracts).
	DeliveryTime      string `json:"deliveryTime"`      // Delivery time (for delivery contracts).
	DeliveryStatus    string `json:"deliveryStatus"`    // Delivery status.
	Open24h           string `json:"open24h"`           // Entry price of the last 24 hours.
	MarkPrice         string `json:"markPrice"`         // Mark price.
}

func (td TickerData) BidPrice() float64 {
	return stringToFloat64(td.BidPr)
}

func (td TickerData) AskPrice() float64 {
	return stringToFloat64(td.AskPr)
}
