package bitget

import (
	"fmt"

	"github.com/KhavrTrading/Khavr-bitget-lib/bitget/bitget_models"
)

// GetOrderFills sends a GET request to Bitget's order fills endpoint to retrieve order fill details.
// It accepts an OrderFillsRequest and API credentials, returning the parsed OrderFillsResponse or an error.
func (c *BitgetClient) GetOrderFills(req *bitget_models.OrderFillsRequest) (*bitget_models.OrderFillsResponse, error) {
	endpoint := "/api/v2/mix/order/fills"

	var resp bitget_models.OrderFillsResponse
	if err := c.doGet(endpoint, req, &resp); err != nil {
		return nil, err
	}

	if resp.Code != "00000" {
		return nil, fmt.Errorf("get order fills failed: %s", resp.Msg)
	}
	return &resp, nil
}
