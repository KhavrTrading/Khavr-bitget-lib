package bitget

import (
	"fmt"

	"github.com/KhavrTrading/Khavr-bitget-lib/bitget/bitget_models"
)

// PlaceTpslOrder sends a stop-profit and stop-loss plan order request to Bitget's API.
// It accepts a PlaceTpslOrderRequest and returns the parsed PlaceTpslOrderResponse or an error.
func (c *BitgetClient) PlaceTpslOrder(req *bitget_models.PlaceTpslOrderRequest) (*bitget_models.PlaceTpslOrderResponse, error) {
	endpoint := "/api/v2/mix/order/place-tpsl-order"

	var resp bitget_models.PlaceTpslOrderResponse
	if err := c.doPost(endpoint, req, &resp); err != nil {
		return nil, err
	}

	if resp.Code != "00000" {
		return nil, fmt.Errorf("place tpsl order failed: %s", resp.Msg)
	}
	return &resp, nil
}
