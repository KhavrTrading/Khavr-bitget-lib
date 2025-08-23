package bitget

import (
	"fmt"

	"github.com/KhavrTrading/Khavr-bitget-lib/bitget/bitget_models"
)

// GetOrderDetail sends a GET request to Bitget's order detail endpoint.
// It accepts an OrderDetailRequest and API credentials, then returns the parsed OrderDetailResponse or an error.
func (c *BitgetClient) GetOrderDetail(req *bitget_models.OrderDetailRequest) (*bitget_models.OrderDetailResponse, error) {
	endpoint := "/api/v2/mix/order/detail"

	var resp bitget_models.OrderDetailResponse
	if err := c.doGet(endpoint, req, &resp); err != nil {
		return nil, err
	}

	if resp.Code != "00000" {
		return nil, fmt.Errorf("get order detail failed: %s", resp.Msg)
	}
	return &resp, nil
}
