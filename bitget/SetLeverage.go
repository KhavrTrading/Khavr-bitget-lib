package bitget

import (
	"fmt"

	"github.com/KhavrTrading/Khavr-bitget-lib/bitget/bitget_models"
)

// SetLeverage sends a request to adjust the leverage for a given symbol, productType, and marginCoin.
func (c *BitgetClient) SetLeverage(req *bitget_models.SetLeverageRequest) (*bitget_models.SetLeverageResponse, error) {
	endpoint := "/api/v2/mix/account/set-leverage"

	var resp bitget_models.SetLeverageResponse
	if err := c.doPost(endpoint, req, &resp); err != nil {
		return nil, err
	}

	if resp.Code != "00000" {
		return nil, fmt.Errorf("set leverage failed: %s", resp.Msg)
	}
	return &resp, nil
}
