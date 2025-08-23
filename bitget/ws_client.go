package bitget

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/KhavrTrading/Khavr-bitget-lib/bitget/bitget_models"

	"github.com/gorilla/websocket"
)

// BitgetWsClient holds the credentials (if needed), the WebSocket connection, and endpoint info.
type BitgetWsClient struct {
	APIKey     string
	APISecret  string
	Passphrase string

	WsConn    *websocket.Conn
	BaseWsURL string // e.g. "wss://ws.bitget.com/v2/ws/public"
	stopChan  chan struct{}
}

// NewBitgetWsClient creates a new BitgetWsClient.
func NewBitgetWsClient(apiKey, apiSecret, passphrase string) *BitgetWsClient {
	return &BitgetWsClient{
		APIKey:     apiKey,
		APISecret:  apiSecret,
		Passphrase: passphrase,
		BaseWsURL:  "wss://ws.bitget.com/v2/ws/public",
		stopChan:   make(chan struct{}),
	}
}

func (c *BitgetWsClient) Connect() error {
	// Copy the default dialer…
	dialer := *websocket.DefaultDialer
	dialer.EnableCompression = true
	// …then bump up its read/write buffer sizes to 16 KiB (or more)
	dialer.ReadBufferSize = 16 * 1024
	dialer.WriteBufferSize = 16 * 1024

	// Now dial using that larger-buffered dialer
	var err error
	c.WsConn, _, err = dialer.Dial(c.BaseWsURL, nil)
	if err != nil {
		return fmt.Errorf("failed to connect to WebSocket: %w", err)
	}
	return nil
}

// Close closes the WebSocket connection.
func (c *BitgetWsClient) Close() error {
	if err := c.WsConn.Close(); err != nil {
		return fmt.Errorf("failed to close WebSocket: %w", err)
	}
	return nil
}

func (c *BitgetWsClient) ListenLoop(
	subReq *bitget_models.CandleSubscriptionRequest,
	callback func(pushData bitget_models.CandlePushData),
) {
	defer func() {
		if r := recover(); r != nil {
			c.Stop()
			c.Close()
		}
	}()

	if err := c.SubscribeCandles(subReq); err != nil {
	}

MainLoop:
	for {
		select {
		case <-c.stopChan:
			c.Close()
			return

		default:
			_, msg, err := c.WsConn.ReadMessage()
			if err != nil {
				errStr := err.Error()
				// skip harmless protocol‐extension frames
				if strings.Contains(errStr, "RSV2 set") ||
					strings.Contains(errStr, "RSV3 set") ||
					strings.Contains(errStr, "bad opcode") {
					continue MainLoop
				}
			ReconnectLoop:
				for {
					select {
					case <-c.stopChan:
						c.Close()
						return
					default:
						if err := c.reconnect(subReq); err != nil {
							time.Sleep(2 * time.Second)
							continue
						}
						break ReconnectLoop
					}
				}
				continue MainLoop
			}

			// only try to decode text→JSON errors
			var pushData bitget_models.CandlePushData
			if err := json.Unmarshal(msg, &pushData); err != nil {
				continue
			}
			callback(pushData)
		}
	}
}

func (c *BitgetWsClient) Stop() {
	select {
	case <-c.stopChan: // already closed
	default:
		close(c.stopChan)
	}
}

func (c *BitgetWsClient) reconnect(subReq *bitget_models.CandleSubscriptionRequest) error {
	// Close any existing connection.
	if err := c.Close(); err != nil {
		log.Printf("Error closing WebSocket: %v", err)
	}

	// Wait a moment before reconnecting.
	time.Sleep(1 * time.Second)

	// Establish a new connection.
	if err := c.Connect(); err != nil {
		return fmt.Errorf("failed to reconnect: %w", err)
	}

	// Re-subscribe to channels.
	if err := c.SubscribeCandles(subReq); err != nil {
		return fmt.Errorf("failed to re-subscribe: %w", err)
	}
	return nil
}

// GetOrderFillMessage is your existing ReadMessage method adapted for ws client.
func (c *BitgetWsClient) ReadMessage() (bitget_models.CandlePushData, error) {
	var pushData bitget_models.CandlePushData
	_, msg, err := c.WsConn.ReadMessage()
	if err != nil {
		return pushData, fmt.Errorf("failed to read message: %w", err)
	}
	if err := json.Unmarshal(msg, &pushData); err != nil {
		return pushData, fmt.Errorf("failed to unmarshal message: %w", err)
	}
	return pushData, nil
}

// SubscribeCandles sends the subscription message.
func (c *BitgetWsClient) SubscribeCandles(req *bitget_models.CandleSubscriptionRequest) error {
	payload, err := json.Marshal(req)
	if err != nil {
		return fmt.Errorf("failed to marshal subscription request: %w", err)
	}
	if err := c.WsConn.WriteMessage(websocket.TextMessage, payload); err != nil {
		return fmt.Errorf("failed to send subscription request: %w", err)
	}
	return nil
}
