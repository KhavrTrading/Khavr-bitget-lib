package bitget

import (
	"fmt"

	"github.com/KhavrTrading/Khavr-bitget-lib/bitget/bitget_models"
)

// ClosePositions sends a request to close positions at market price.
// It accepts a ClosePositionsRequest and API credentials, returning the parsed response or an error.
func (c *BitgetClient) ClosePositions(req *bitget_models.ClosePositionsRequest) (*bitget_models.ClosePositionsResponse, error) {
	endpoint := "/api/v2/mix/order/close-positions"

	var resp bitget_models.ClosePositionsResponse
	if err := c.doPost(endpoint, req, &resp); err != nil {
		return nil, err
	}

	if resp.Code != "00000" {
		return nil, fmt.Errorf("close positions failed: %s", resp.Msg)
	}
	return &resp, nil
}
