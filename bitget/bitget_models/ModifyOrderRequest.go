package bitget_models

// ModifyOrderRequest represents the payload for modifying an order.
// Either orderId or clientOid must be provided (if both are provided, orderId takes precedence).
type ModifyOrderRequest struct {
	OrderId                   string `json:"orderId,omitempty"`                   // Order ID (optional)
	ClientOid                 string `json:"clientOid,omitempty"`                 // Custom order ID (optional)
	Symbol                    string `json:"symbol"`                              // Trading pair, e.g. "ETHUSDT"
	ProductType               string `json:"productType"`                         // Product type (e.g. "USDT-FUTURES")
	MarginCoin                string `json:"marginCoin"`                          // Margin coin (capitalized), e.g. "USDT"
	NewClientOid              string `json:"newClientOid"`                        // New customized order ID after modification
	NewSize                   string `json:"newSize,omitempty"`                   // Modified order amount (if omitted, size stays unchanged)
	NewPrice                  string `json:"newPrice,omitempty"`                  // Modified price for new order (for limit orders)
	NewPresetStopSurplusPrice string `json:"newPresetStopSurplusPrice,omitempty"` // Modified take-profit value
	NewPresetStopLossPrice    string `json:"newPresetStopLossPrice,omitempty"`    // Modified stop-loss value
}
