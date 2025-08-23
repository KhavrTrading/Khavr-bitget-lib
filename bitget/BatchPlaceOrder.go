package bitget

import (
	"fmt"

	"github.com/KhavrTrading/Khavr-bitget-lib/bitget/bitget_models"
)

// BatchPlaceOrder sends a batch order request to Bitget's exchange API.
// It accepts a BatchOrderRequest and API credentials, then returns the parsed BatchPlaceOrderResponse or an error.
func (c *BitgetClient) BatchPlaceOrder(req *bitget_models.BatchOrderRequest) (*bitget_models.BatchPlaceOrderResponse, error) {
	endpoint := "/api/v2/mix/order/batch-place-order"

	var resp bitget_models.BatchPlaceOrderResponse
	if err := c.doPost(endpoint, req, &resp); err != nil {
		return nil, err
	}

	if resp.Code != "00000" {
		return nil, fmt.Errorf("batch order failed: %s", resp.Msg)
	}
	return &resp, nil
}
