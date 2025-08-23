package bitget

import (
	"fmt"

	"github.com/KhavrTrading/Khavr-bitget-lib/bitget/bitget_models"
)

// CancelAllOrders sends a cancel-all-orders request to Bitget's API.
// It accepts a CancelAllOrdersRequest and API credentials, returning a CancelAllOrdersResponse or an error.
func (c *BitgetClient) CancelAllOrders(req *bitget_models.CancelAllOrdersRequest) (*bitget_models.CancelAllOrdersResponse, error) {
	endpoint := "/api/v2/mix/order/cancel-all-orders"

	var resp bitget_models.CancelAllOrdersResponse
	if err := c.doPost(endpoint, req, &resp); err != nil {
		return nil, err
	}

	if resp.Code != "00000" {
		return nil, fmt.Errorf("cancel all orders failed: %s", resp.Msg)
	}
	return &resp, nil
}
