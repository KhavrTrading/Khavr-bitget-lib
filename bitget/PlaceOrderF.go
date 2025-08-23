package bitget

import (
	"fmt"

	"github.com/KhavrTrading/Khavr-bitget-lib/bitget/bitget_models"
)

// PlaceOrder sends a place order request to Bitget's exchange API.
// It takes a PlaceOrderRequest and API credentials, and returns the parsed PlaceOrderResponse or an error.
func (c *BitgetClient) PlaceOrder(req *bitget_models.PlaceOrderRequest) (*bitget_models.PlaceOrderResponse, error) {
	endpoint := "/api/v2/mix/order/place-order"

	var resp bitget_models.PlaceOrderResponse
	if err := c.doPost(endpoint, req, &resp); err != nil {
		return nil, err
	}

	if resp.Code != "00000" {
		return nil, fmt.Errorf("order failed: %s", resp.Msg)
	}
	return &resp, nil
}
