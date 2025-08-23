package bitget

import (
	"fmt"

	"github.com/KhavrTrading/Khavr-bitget-lib/bitget/bitget_models"
)

// GetPendingOrders sends a GET request to Bitget's "orders-pending" endpoint to query all existing pending orders.
// It accepts an OrdersPendingRequest and API credentials, and returns the parsed OrdersPendingResponse or an error.
func (c *BitgetClient) GetPendingOrders(req *bitget_models.OrdersPendingRequest) (*bitget_models.OrdersPendingResponse, error) {
	endpoint := "/api/v2/mix/order/orders-pending"

	var resp bitget_models.OrdersPendingResponse
	if err := c.doGet(endpoint, req, &resp); err != nil {
		return nil, err
	}

	if resp.Code != "00000" {
		return nil, fmt.Errorf("query pending orders failed: %s", resp.Msg)
	}

	return &resp, nil
}
