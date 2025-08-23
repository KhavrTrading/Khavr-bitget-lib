package bitget

import (
	"fmt"

	"github.com/KhavrTrading/Khavr-bitget-lib/bitget/bitget_models"
)

func (c *BitgetClient) GetHistoryTransactions(req *bitget_models.HistoryTransactionRequest) (*bitget_models.HistoryTransactionResponse, error) {
	endpoint := "/api/v2/mix/market/fills-history"

	var resp bitget_models.HistoryTransactionResponse
	if err := c.doGet(endpoint, req, &resp); err != nil {
		return nil, err
	}

	if resp.Code != "00000" {
		return nil, fmt.Errorf("get history transactions failed: %s", resp.Msg)
	}
	return &resp, nil
}
