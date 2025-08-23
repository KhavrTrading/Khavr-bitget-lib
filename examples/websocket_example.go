// Package main demonstrates WebSocket usage of the Khavr Bitget library.
//
// To use this library in your own project, you would import it like this:
//   import "github.com/KhavrTrading/Khavr-bitget-lib"
//
// But since this example is part of the same module, we use internal imports.
package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/KhavrTrading/Khavr-bitget-lib/bitget"
	"github.com/KhavrTrading/Khavr-bitget-lib/bitget/bitget_models"
)

func main() {
	// Create WebSocket client (empty credentials for public endpoints)
	wsClient := bitget.NewBitgetWsClient("", "", "")

	// Connect to WebSocket
	fmt.Println("Connecting to Bitget WebSocket...")
	err := wsClient.Connect()
	if err != nil {
		log.Fatal("Failed to connect:", err)
	}
	defer wsClient.Close()

	fmt.Println("Connected successfully!")

	// Set up graceful shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// Example 1: Subscribe to 1-minute candlestick data for BTCUSDT
	fmt.Println("Subscribing to BTCUSDT 1m candlestick data...")
	
	subscription := &bitget_models.CandleSubscriptionRequest{
		Op: "subscribe",
		Args: []bitget_models.CandleSubscriptionArg{
			{
				InstType: "USDT-FUTURES",
				Channel:  "candle1m",
				InstId:   "BTCUSDT",
			},
		},
	}	err = wsClient.SubscribeCandles(subscription)
	if err != nil {
		log.Fatal("Failed to subscribe to candles:", err)
	}

	fmt.Println("Subscription successful! Waiting for data...")

	// Example 2: Subscribe to multiple symbols
	fmt.Println("Subscribing to multiple symbols...")
	
	multiSubscription := &bitget_models.CandleSubscriptionRequest{
		Op: "subscribe",
		Args: []bitget_models.CandleSubscriptionArg{
			{
				InstType: "USDT-FUTURES",
				Channel:  "candle5m",
				InstId:   "ETHUSDT",
			},
			{
				InstType: "USDT-FUTURES",
				Channel:  "candle5m",
				InstId:   "ADAUSDT",
			},
		},
	}	err = wsClient.SubscribeCandles(multiSubscription)
	if err != nil {
		log.Printf("Failed to subscribe to multiple symbols: %v", err)
	}

	// Start reading messages in a goroutine
	go func() {
		for {
			data, err := wsClient.ReadMessage()
			if err != nil {
				log.Printf("Error reading message: %v", err)
				continue
			}

			// Handle received candle data
			fmt.Printf("� %s Candle Data:\n", data.Arg.InstId)
			fmt.Printf("   Action: %s\n", data.Action)
			fmt.Printf("   Channel: %s\n", data.Arg.Channel)
			fmt.Printf("   Timestamp: %d\n", data.Ts)

			for i, candle := range data.Data {
				if len(candle) >= 7 {
					fmt.Printf("   Candle %d: [Time=%s, Open=%s, High=%s, Low=%s, Close=%s, Volume=%s, Quote Volume=%s]\n",
						i+1, candle[0], candle[1], candle[2], candle[3], candle[4], candle[5], candle[6])
				}
			}
			fmt.Println("   ---")
		}
	}()

	// Keep the connection alive and listen for data
	fmt.Println("🔄 Listening for real-time data... Press Ctrl+C to stop")

	// Use a ticker to show connection status
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-sigChan:
			fmt.Println("\n🛑 Received shutdown signal, closing connection...")
			return

		case <-ticker.C:
			fmt.Println("⏰ Connection alive, waiting for more data...")

		default:
			// Keep the main goroutine alive
			time.Sleep(100 * time.Millisecond)
		}
	}
}
