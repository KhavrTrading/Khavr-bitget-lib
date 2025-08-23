package bitget

import (
	"fmt"

	"github.com/KhavrTrading/Khavr-bitget-lib/bitget/bitget_models"
)

// GetSinglePosition sends a GET request to retrieve position information for a single symbol.
func (c *BitgetClient) GetSinglePosition(req *bitget_models.SinglePositionRequest) (*bitget_models.SinglePositionResponse, error) {
	endpoint := "/api/v2/mix/position/single-position"

	var resp bitget_models.SinglePositionResponse
	if err := c.doGet(endpoint, req, &resp); err != nil {
		return nil, err
	}

	if resp.Code != "00000" {
		return nil, fmt.Errorf("get single position failed: %s", resp.Msg)
	}
	return &resp, nil
}
