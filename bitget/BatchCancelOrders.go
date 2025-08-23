package bitget

import (
	"fmt"

	"github.com/KhavrTrading/Khavr-bitget-lib/bitget/bitget_models"
)

// BatchCancelOrders sends a batch cancel order request to Bitget's exchange API.
// It takes a BatchCancelOrderRequest and API credentials, then returns the parsed BatchCancelOrderResponse or an error.
func (c *BitgetClient) BatchCancelOrders(req *bitget_models.BatchCancelOrderRequest) (*bitget_models.BatchCancelOrderResponse, error) {
	endpoint := "/api/v2/mix/order/batch-cancel-orders"

	var resp bitget_models.BatchCancelOrderResponse
	if err := c.doPost(endpoint, req, &resp); err != nil {
		return nil, err
	}

	if resp.Code != "00000" {
		return nil, fmt.Errorf("batch cancel order failed: %s", resp.Msg)
	}
	return &resp, nil
}
