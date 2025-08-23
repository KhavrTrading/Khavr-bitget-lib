package bitget

import (
	"context"
	"fmt"
	"time"

	"github.com/KhavrTrading/Khavr-bitget-lib/bitget/bitget_models"
)

// GetTicker gets ticker data for the specified product type and symbol.
// func (c *BitgetClient) GetTicker(req *bitget_models.TickerRequest) (*bitget_models.TickerResponse, error) {
// 	endpoint := "/api/v2/mix/market/ticker"

// 	var resp bitget_models.TickerResponse
// 	if err := c.doGet(endpoint, req, &resp); err != nil {
// 		return nil, err
// 	}

// 	if resp.Code != "00000" {
// 		return nil, fmt.Errorf("get ticker failed: %s", resp.Msg)
// 	}
// 	return &resp, nil
// }

func (c *BitgetClient) GetTicker(req *bitget_models.TickerRequest) (*bitget_models.TickerResponse, error) {
	endpoint := "/api/v2/mix/market/ticker"
	var (
		resp    bitget_models.TickerResponse
		lastErr error
		backoff = initialBackoff
	)

	// Try up to maxRetries times
	for attempt := 1; attempt <= maxRetries; attempt++ {
		// 1) rate-limit
		if err := limiter.Wait(context.Background()); err != nil {
			return nil, fmt.Errorf("rate limiter error: %w", err)
		}

		// 2) do the HTTP GET
		lastErr = c.doGet(endpoint, req, &resp)
		if lastErr != nil {
			// network / deserialization failure → retry
			time.Sleep(backoff)
			backoff *= 2
			continue
		}

		// 3) check API-level response code
		if resp.Code != "00000" {
			lastErr = fmt.Errorf("bitget error: %s", resp.Msg)
			time.Sleep(backoff)
			backoff *= 2
			continue
		}

		// success!
		return &resp, nil
	}

	// if we fall out of the loop, every attempt failed
	return nil, fmt.Errorf("GetTicker failed after %d attempts: %w", maxRetries, lastErr)
}
