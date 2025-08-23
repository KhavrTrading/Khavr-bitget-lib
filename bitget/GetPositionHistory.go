package bitget

import (
	"fmt"

	"github.com/KhavrTrading/Khavr-bitget-lib/bitget/bitget_models"
)

func (c *BitgetClient) GetPositionHistory(req *bitget_models.PositionHistoryRequest) (*bitget_models.PositionHistoryResponse, error) {
	endpoint := "/api/v2/mix/position/history-position"

	var resp bitget_models.PositionHistoryResponse
	if err := c.doGet(endpoint, req, &resp); err != nil {
		return nil, err
	}

	if resp.Code != "00000" {
		return nil, fmt.Errorf("query historical positions failed: %s", resp.Msg)
	}

	return &resp, nil
}
