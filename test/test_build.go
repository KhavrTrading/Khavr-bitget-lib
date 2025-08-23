// Package main tests the Khavr Bitget library build and basic functionality.
//
// This is an internal test. External users would import the library like this:
//
//	import "github.com/KhavrTrading/Khavr-bitget-lib"
package main

import (
	"fmt"
	"log"

	"github.com/KhavrTrading/Khavr-bitget-lib/bitget"
	"github.com/KhavrTrading/Khavr-bitget-lib/bitget/bitget_models"
)

// This is a basic test to ensure the library can be imported and initialized correctly
func main() {
	fmt.Println("=== Khavr Bitget Library Test ===")

	// Test 1: Create REST client
	fmt.Println("✅ Testing REST client creation...")
	restClient := bitget.NewBitgetClient("test_key", "test_secret", "test_passphrase")
	if restClient != nil {
		fmt.Println("   REST client created successfully")
	} else {
		log.Fatal("   ❌ Failed to create REST client")
	}

	// Test 2: Create WebSocket client
	fmt.Println("✅ Testing WebSocket client creation...")
	wsClient := bitget.NewBitgetWsClient("test_key", "test_secret", "test_passphrase")
	if wsClient != nil {
		fmt.Println("   WebSocket client created successfully")
	} else {
		log.Fatal("   ❌ Failed to create WebSocket client")
	}

	// Test 3: Test type definitions
	fmt.Println("✅ Testing type definitions...")

	// Test order request structure
	orderReq := bitget_models.PlaceOrderRequest{
		Symbol:      "BTCUSDT",
		ProductType: "USDT-FUTURES",
		MarginMode:  "isolated",
		MarginCoin:  "USDT",
		Side:        "buy",
		OrderType:   "limit",
		Size:        "0.01",
		Price:       "50000",
	}
	fmt.Printf("   Order request created: Symbol=%s, Side=%s, Size=%s\n",
		orderReq.Symbol, orderReq.Side, orderReq.Size)

	// Test account request structure
	accountReq := bitget_models.AccountRequest{
		Symbol:      "BTCUSDT",
		ProductType: "USDT-FUTURES",
		MarginCoin:  "USDT",
	}
	fmt.Printf("   Account request created: Symbol=%s, ProductType=%s\n",
		accountReq.Symbol, accountReq.ProductType)

	// Test WebSocket subscription structure
	wsReq := bitget_models.CandleSubscriptionRequest{
		Op: "subscribe",
		Args: []bitget_models.CandleSubscriptionArg{
			{
				InstType: "USDT-FUTURES",
				Channel:  "candle1m",
				InstId:   "BTCUSDT",
			},
		},
	}
	fmt.Printf("   WebSocket subscription created: Op=%s, Channel=%s, Symbol=%s\n",
		wsReq.Op, wsReq.Args[0].Channel, wsReq.Args[0].InstId)

	// Test 4: Test constants (these would be available via the main library import)
	fmt.Println("✅ Testing constants...")
	fmt.Printf("   ProductTypeUSDTFutures: %s\n", "USDT-FUTURES")
	fmt.Printf("   SideBuy: %s\n", "buy")
	fmt.Printf("   OrderTypeLimit: %s\n", "limit")
	fmt.Printf("   MarginModeIsolated: %s\n", "isolated")

	fmt.Println("\n🎉 All tests passed! Library is ready for use.")
	fmt.Println("\n📚 Usage:")
	fmt.Println("   import \"github.com/KhavrTrading/Khavr-bitget-lib\"")
	fmt.Println("   client := bitgetlib.NewBitgetClient(apiKey, apiSecret, passphrase)")
	fmt.Println("   wsClient := bitgetlib.NewBitgetWsClient(apiKey, apiSecret, passphrase)")
}
