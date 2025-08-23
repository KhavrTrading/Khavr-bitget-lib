// bitget\GetAccount.go
package bitget

import (
	"fmt"

	"github.com/KhavrTrading/Khavr-bitget-lib/bitget/bitget_models"
)

func (c *BitgetClient) GetAccount(req *bitget_models.AccountRequest) (*bitget_models.AccountResponse, error) {
	endpoint := "/api/v2/mix/account/account"

	var resp bitget_models.AccountResponse
	if err := c.doGet(endpoint, req, &resp); err != nil {
		return nil, err
	}

	if resp.Code != "00000" {
		return nil, fmt.Errorf("get account failed: %s", resp.Msg)
	}
	return &resp, nil
}
