package bitget

import (
	"fmt"

	"github.com/KhavrTrading/Khavr-bitget-lib/bitget/bitget_models"
)

// GetOrdersHistory sends a GET request to Bitget's "orders-history" endpoint to retrieve historical orders.
// It accepts an OrdersHistoryRequest and API credentials, and returns the parsed OrdersHistoryResponse or an error.
func (c *BitgetClient) GetOrdersHistory(req *bitget_models.OrdersHistoryRequest) (*bitget_models.OrdersHistoryResponse, error) {
	endpoint := "/api/v2/mix/order/orders-history"

	var resp bitget_models.OrdersHistoryResponse
	if err := c.doGet(endpoint, req, &resp); err != nil {
		return nil, err
	}

	if resp.Code != "00000" {
		return nil, fmt.Errorf("query historical orders failed: %s", resp.Msg)
	}

	return &resp, nil
}
