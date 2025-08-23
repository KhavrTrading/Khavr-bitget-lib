package bitget

import (
	"fmt"

	"github.com/KhavrTrading/Khavr-bitget-lib/bitget/bitget_models"
)

// CancelOrder sends a cancel order request to Bitget's exchange API.
// It takes a CancelOrderRequest and API credentials, then returns the parsed CancelOrderResponse or an error.
func (c *BitgetClient) CancelOrder(req *bitget_models.CancelOrderRequest) (*bitget_models.CancelOrderResponse, error) {
	endpoint := "/api/v2/mix/order/cancel-order"

	var resp bitget_models.CancelOrderResponse
	if err := c.doPost(endpoint, req, &resp); err != nil {
		return nil, err
	}

	if resp.Code != "00000" {
		return nil, fmt.Errorf("cancel order failed: %s", resp.Msg)
	}
	return &resp, nil
}
