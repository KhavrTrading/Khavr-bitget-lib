package bitget_models

type CandleSubscriptionArg struct {
	InstType string `json:"instType"` // e.g. "USDT-FUTURES"
	Channel  string `json:"channel"`  // e.g. "candle1m", "candle5m", etc.
	InstId   string `json:"instId"`   // e.g. "BTCUSDT"
}

// CandleSubscriptionRequest is the request format for subscribing.
type CandleSubscriptionRequest struct {
	Op   string                    `json:"op"`   // "subscribe" or "unsubscribe"
	Args []CandleSubscriptionArg   `json:"args"` // List of channels to subscribe
}

// CandlePushData represents the pushed candlestick data.
type CandlePushData struct {
	Action string `json:"action"` // e.g. "snapshot" or "update"
	Arg    struct {
		InstType string `json:"instType"` // Product type.
		Channel  string `json:"channel"`  // Channel name.
		InstId   string `json:"instId"`   // Product ID.
	} `json:"arg"`
	Data [][]string `json:"data"` // Each inner array contains: [timestamp, open, high, low, close, volume, quoteVolume, usdtVolume]
	Ts   int64      `json:"ts"`   // Server timestamp.
}
