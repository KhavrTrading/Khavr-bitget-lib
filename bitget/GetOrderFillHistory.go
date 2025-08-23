package bitget

import (
	"fmt"

	"github.com/KhavrTrading/Khavr-bitget-lib/bitget/bitget_models"
)

// GetOrderFillHistory sends a GET request to Bitget's historical order fill endpoint.
// It accepts an OrderFillHistoryRequest and API credentials, then returns the parsed OrderFillHistoryResponse or an error.
func (c *BitgetClient) GetOrderFillHistory(req *bitget_models.OrderFillHistoryRequest) (*bitget_models.OrderFillHistoryResponse, error) {
	endpoint := "/api/v2/mix/order/fill-history"

	var resp bitget_models.OrderFillHistoryResponse
	if err := c.doGet(endpoint, req, &resp); err != nil {
		return nil, err
	}

	if resp.Code != "00000" {
		return nil, fmt.Errorf("get historical transaction details failed: %s", resp.Msg)
	}
	return &resp, nil
}
