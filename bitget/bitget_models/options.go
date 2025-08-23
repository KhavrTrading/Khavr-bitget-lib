package bitget_models

// Product types available.
const (
	ProductTypeUSDTFutures  = "USDT-FUTURES"
	ProductTypeCOINFutures  = "COIN-FUTURES"
	ProductTypeUSDCFutures  = "USDC-FUTURES"
	ProductTypeSUSDTFutures = "SUSDT-FUTURES"
	ProductTypeSCOINFutures = "SCOIN-FUTURES"
	ProductTypeSUSDCFutures = "SUSDC-FUTURES"
)

// Margin modes.
const (
	MarginModeIsolated = "isolated"
	MarginModeCrossed  = "crossed"
)

// Trade sides.
const (
	SideBuy  = "buy"
	SideSell = "sell"
)

// Hedge mode trade sides.
const (
	TradeSideOpen  = "open"
	TradeSideClose = "close"
)

// Order types.
const (
	OrderTypeLimit  = "limit"
	OrderTypeMarket = "market"
)

// Force options for limit orders.
const (
	ForceGTC      = "gtc"
	ForceIOC      = "ioc"
	ForceFOK      = "fok"
	ForcePostOnly = "post_only"
)

// Reduce only options.
const (
	ReduceOnlyYes = "YES"
	ReduceOnlyNo  = "NO"
)

// STP Modes (Self Trade Prevention).
const (
	STPModeNone        = "none"
	STPModeCancelTaker = "cancel_taker"
	STPModeCancelMaker = "cancel_maker"
	STPModeCancelBoth  = "cancel_both"
)

const (
	// DefaultMarginCoin is the default margin coin.
	DefaultMarginCoin = "usdt"
)

const (
	// hold side simbol long or short
	HoldSideLong  = "long"
	HoldSideShort = "short"
)

// live: New order, waiting for a match in orderbook
// partially_filled: Partially filled
// filled: All filled
// canceled: the order is cancelled

const (
	OrderStateLive            = "live"
	OrderStatePartiallyFilled = "partially_filled"
	OrderStateFilled          = "filled"
	OrderStateCanceled        = "canceled"
)

// fill_price: market price;
// mark_price: mark price

const (
	TriggerTypeFillPrice = "fill_price"
	TriggerTypeMarkPrice = "mark_price"
)
