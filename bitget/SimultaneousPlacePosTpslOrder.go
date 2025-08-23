package bitget

import (
	"context"
	"fmt"
	"time"

	// Adjust the import path as needed.

	"github.com/KhavrTrading/Khavr-bitget-lib/bitget/bitget_models"
	"golang.org/x/time/rate"
)

const (
	maxRetries     = 5
	initialBackoff = 500 * time.Millisecond
)

var limiter = rate.NewLimiter(rate.Every(time.Second/20), 1)

// SimultaneousPlacePosTpslOrder sends a simultaneous stop-profit and stop-loss plan order request to Bitget's API.
// It accepts a PlacePosTpslOrderRequest and returns the parsed PlacePosTpslOrderResponse or an error.
// func (c *BitgetClient) SimultaneousPlacePosTpslOrder(req *bitget_models.PlacePosTpslOrderRequest) (*bitget_models.PlacePosTpslOrderResponse, error) {
// 	endpoint := "/api/v2/mix/order/place-pos-tpsl"

// 	var resp bitget_models.PlacePosTpslOrderResponse
// 	if err := c.doPost(endpoint, req, &resp); err != nil {
// 		return nil, err
// 	}

// 	if resp.Code != "00000" {
// 		return nil, fmt.Errorf("simultaneous place pos tpsl order failed: %s", resp.Msg)
// 	}
// 	return &resp, nil
// }

func (c *BitgetClient) SimultaneousPlacePosTpslOrder(
	req *bitget_models.PlacePosTpslOrderRequest,
) (*bitget_models.PlacePosTpslOrderResponse, error) {
	endpoint := "/api/v2/mix/order/place-pos-tpsl"
	var resp bitget_models.PlacePosTpslOrderResponse
	var err error

	for attempt := 0; attempt < maxRetries; attempt++ {
		// 1) wait for our rate‐limit token
		if err = limiter.Wait(context.Background()); err != nil {
			return nil, fmt.Errorf("rate limiter wait failed: %w", err)
		}

		// 2) do the actual request
		err = c.doPost(endpoint, req, &resp)
		if err == nil && resp.Code == "00000" {
			// success
			return &resp, nil
		}

		// if the API returned a non-zero code, treat it as an error
		if err == nil {
			err = fmt.Errorf("simultaneous place pos tpsl order failed: %s", resp.Msg)
		}

		// 3) exponential backoff before the next try
		backoff := initialBackoff * (1 << attempt)
		time.Sleep(backoff)
	}

	return nil, fmt.Errorf("after %d attempts, last error: %w", maxRetries, err)
}
