package bitget

import (
	"fmt"

	"github.com/KhavrTrading/Khavr-bitget-lib/bitget/bitget_models"
)

// ModifyOrder sends a modify order request to Bitget's exchange API.
// It takes a ModifyOrderRequest and API credentials, and returns the parsed ModifyOrderResponse or an error.
func (c *BitgetClient) ModifyOrder(req *bitget_models.ModifyOrderRequest) (*bitget_models.ModifyOrderResponse, error) {
	endpoint := "/api/v2/mix/order/modify-order"

	var resp bitget_models.ModifyOrderResponse
	if err := c.doPost(endpoint, req, &resp); err != nil {
		return nil, err
	}

	if resp.Code != "00000" {
		return nil, fmt.Errorf("modify order failed: %s", resp.Msg)
	}
	return &resp, nil
}
