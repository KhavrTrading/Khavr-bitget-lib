package bitget_models

import (
	"maps"
	"strconv"
)

// AccountRequest holds the query parameters for the account details endpoint.
type AccountRequest struct {
	// Trading pair, e.g. "btcusdt"
	Symbol string `json:"symbol"`
	// Product type, e.g. "USDT-FUTURES", "COIN-FUTURES", etc.
	ProductType string `json:"productType"`
	// Margin coin, e.g. "usdt"
	MarginCoin string `json:"marginCoin"`
}

func (r *AccountRequest) ToParams() map[string]string {
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

// AccountResponse represents the API response for the account details.
type AccountResponse struct {
	Code        string      `json:"code"`        // "00000" indicates success.
	Data        AccountData `json:"data"`        // Account details.
	Msg         string      `json:"msg"`         // Response message.
	RequestTime int64       `json:"requestTime"` // Request timestamp.
}

func (r *AccountResponse) ToParams() map[string]string {
	params := make(map[string]string)
	if r.Code != "" {
		params["code"] = r.Code
	}
	if r.Data != (AccountData{}) {
		dataParams := r.Data.ToParams()
		maps.Copy(params, dataParams)
	}
	if r.Msg != "" {
		params["msg"] = r.Msg
	}
	if r.RequestTime != 0 {
		params["requestTime"] = strconv.FormatInt(r.RequestTime, 10)
	}
	return params
}

// AccountData holds the actual account information.
type AccountData struct {
	MarginCoin            string         `json:"marginCoin"`            // Margin coin.
	Locked                string         `json:"locked"`                // Locked quantity.
	Available             string         `json:"available"`             // Available balance.
	CrossedMaxAvailable   string         `json:"crossedMaxAvailable"`   // Max available under cross margin.
	IsolatedMaxAvailable  string         `json:"isolatedMaxAvailable"`  // Max available under isolated margin.
	MaxTransferOut        string         `json:"maxTransferOut"`        // Max transferable amount.
	AccountEquity         string         `json:"accountEquity"`         // Account equity (includes unrealized PnL).
	UsdtEquity            string         `json:"usdtEquity"`            // Equity in USDT.
	BtcEquity             string         `json:"btcEquity"`             // Equity in BTC.
	CrossedRiskRate       string         `json:"crossedRiskRate"`       // Risk rate in cross margin mode.
	CrossedMarginLeverage StringOrNumber `json:"crossedMarginLeverage"` // Leverage for cross margin.
	IsolatedLongLever     StringOrNumber `json:"isolatedLongLever"`     // Leverage for isolated long positions.
	IsolatedShortLever    StringOrNumber `json:"isolatedShortLever"`    // Leverage for isolated short positions.
	MarginMode            string         `json:"marginMode"`            // Margin mode: "isolated" or "crossed".
	PosMode               string         `json:"posMode"`               // Position mode: "one_way_mode" or "hedge_mode".
	UnrealizedPL          string         `json:"unrealizedPL"`          // Unrealized profit and loss.
	Coupon                string         `json:"coupon"`                // Trading bonus.
	CrossedUnrealizedPL   string         `json:"crossedUnrealizedPL"`   // Unrealized PnL in cross margin.
	IsolatedUnrealizedPL  string         `json:"isolatedUnrealizedPL"`  // Unrealized PnL in isolated margin.
	AssetMode             string         `json:"assetMode"`             // Assets mode ("union" or "single").
}

func (d *AccountData) ToParams() map[string]string {
	params := make(map[string]string)
	if d.MarginCoin != "" {
		params["marginCoin"] = d.MarginCoin
	}
	if d.Locked != "" {
		params["locked"] = d.Locked
	}
	if d.Available != "" {
		params["available"] = d.Available
	}
	if d.CrossedMaxAvailable != "" {
		params["crossedMaxAvailable"] = d.CrossedMaxAvailable
	}
	if d.IsolatedMaxAvailable != "" {
		params["isolatedMaxAvailable"] = d.IsolatedMaxAvailable
	}
	if d.MaxTransferOut != "" {
		params["maxTransferOut"] = d.MaxTransferOut
	}
	if d.AccountEquity != "" {
		params["accountEquity"] = d.AccountEquity
	}
	if d.UsdtEquity != "" {
		params["usdtEquity"] = d.UsdtEquity
	}
	if d.BtcEquity != "" {
		params["btcEquity"] = d.BtcEquity
	}
	if d.CrossedRiskRate != "" {
		params["crossedRiskRate"] = d.CrossedRiskRate
	}
	if d.CrossedMarginLeverage != "" {
		if leverage, err := d.CrossedMarginLeverage.Float64(); err == nil {
			params["crossedMarginLeverage"] = strconv.FormatFloat(leverage, 'f', -1, 64)
		}
	}
	if d.IsolatedLongLever != "" {
		if leverage, err := d.IsolatedLongLever.Float64(); err == nil {
			params["isolatedLongLever"] = strconv.FormatFloat(leverage, 'f', -1, 64)
		}
	}
	if d.IsolatedShortLever != "" {
		if leverage, err := d.IsolatedShortLever.Float64(); err == nil {
			params["isolatedShortLever"] = strconv.FormatFloat(leverage, 'f', -1, 64)
		}
	}
	if d.MarginMode != "" {
		params["marginMode"] = d.MarginMode
	}
	if d.PosMode != "" {
		params["posMode"] = d.PosMode
	}
	if d.UnrealizedPL != "" {
		params["unrealizedPL"] = d.UnrealizedPL
	}
	if d.Coupon != "" {
		params["coupon"] = d.Coupon
	}
	if d.CrossedUnrealizedPL != "" {
		params["crossedUnrealizedPL"] = d.CrossedUnrealizedPL
	}
	if d.IsolatedUnrealizedPL != "" {
		params["isolatedUnrealizedPL"] = d.IsolatedUnrealizedPL
	}
	if d.AssetMode != "" {
		params["assetMode"] = d.AssetMode
	}
	return params
}

// stringToFloat64 converts a string to a float64.
func stringToFloat64(s string) float64 {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0
	}
	return f
}

func (a *AccountData) LockedFloat64() float64 {
	return stringToFloat64(a.Locked)
}

func (a *AccountData) AvailableFloat64() float64 {
	return stringToFloat64(a.Available)
}

func (a *AccountData) CrossedMaxAvailableFloat64() float64 {
	return stringToFloat64(a.CrossedMaxAvailable)
}

func (req *AccountRequest) ToQueryString() string {
	return "symbol=" + req.Symbol + "&productType=" + req.ProductType + "&marginCoin=" + req.MarginCoin
}

// usdtEquity float64
func (a *AccountData) UsdtEquityFloat64() float64 {
	return stringToFloat64(a.UsdtEquity)
}
