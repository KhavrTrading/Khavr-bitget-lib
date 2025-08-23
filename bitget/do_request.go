// do_request.go
package bitget

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/KhavrTrading/Khavr-bitget-lib/bitget/bitget_models"
)

func (c *BitgetClient) doGet(endpoint string, pb bitget_models.ParamBuilder, out interface{}) error {
	params := pb.ToParams()
	queryString := BuildGetParams(params)

	url := c.BaseURL + endpoint + queryString

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return fmt.Errorf("failed to create HTTP request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("locale", "en-US")

	timestamp := fmt.Sprintf("%d", time.Now().UnixMilli())
	bodyString := ""
	sign := c.SignPayload("GET", endpoint+queryString, bodyString, timestamp)

	req.Header.Set("ACCESS-KEY", c.APIKey)
	req.Header.Set("ACCESS-SIGN", sign)
	req.Header.Set("ACCESS-PASSPHRASE", c.Passphrase)
	req.Header.Set("ACCESS-TIMESTAMP", timestamp)

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send HTTP request: %w", err)
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response: %w", err)
	}

	if err := json.Unmarshal(bodyBytes, out); err != nil {
		return fmt.Errorf("failed to decode response: %w", err)
	}

	return nil
}

func (c *BitgetClient) doPost(endpoint string, body interface{}, out interface{}) error {
	// 1) Marshal the request body to JSON.
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return fmt.Errorf("failed to marshal request body: %w", err)
	}
	bodyString := string(bodyBytes)

	// 2) Build the final URL.
	url := c.BaseURL + endpoint

	// 3) Create a new POST request.
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyBytes))
	if err != nil {
		return fmt.Errorf("failed to create HTTP request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("locale", "en-US") // or "en-US", etc.

	// 4) Sign the request.
	timestamp := fmt.Sprintf("%d", time.Now().UnixMilli())
	sign := c.SignPayload("POST", endpoint, bodyString, timestamp)

	req.Header.Set("ACCESS-KEY", c.APIKey)
	req.Header.Set("ACCESS-SIGN", sign)
	req.Header.Set("ACCESS-PASSPHRASE", c.Passphrase)
	req.Header.Set("ACCESS-TIMESTAMP", timestamp)

	// 5) Send the request.
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send POST request: %w", err)
	}
	defer resp.Body.Close()

	// 6) Read the response body.
	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response: %w", err)
	}

	// 7) Unmarshal into 'out'.
	if err := json.Unmarshal(respBytes, out); err != nil {
		return fmt.Errorf("failed to decode response: %w", err)
	}

	return nil
}
