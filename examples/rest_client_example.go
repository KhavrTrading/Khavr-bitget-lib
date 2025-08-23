// Package main demonstrates basic REST API usage of the Khavr Bitget library.
//
// To use this library in your own project, you would import it like this:
//
//	import "github.com/KhavrTrading/Khavr-bitget-lib"
//
// But since this example is part of the same module, we use internal imports.
package main

import (
	"fmt"
	"log"

	"github.com/KhavrTrading/Khavr-bitget-lib/bitget"
	"github.com/KhavrTrading/Khavr-bitget-lib/bitget/bitget_models"
)

func main() {
	// Create a new Bitget client with your credentials
	client := bitget.NewBitgetClient("your_api_key", "your_api_secret", "your_passphrase")

	// Example 1: Get account information
	fmt.Println("=== Getting Account Information ===")
	accountReq := bitget_models.AccountRequest{
		Symbol:      "BTCUSDT",
		ProductType: "USDT-FUTURES",
		MarginCoin:  "USDT",
	}

	accountResp, err := client.GetAccount(accountReq)
	if err != nil {
		log.Printf("Error getting account: %v", err)
	} else {
		fmt.Printf("Account Balance: %+v\n", accountResp.Data)
	}

	// Example 2: Place a limit order
	fmt.Println("\n=== Placing a Limit Order ===")
	orderRequest := bitget_models.PlaceOrderRequest{
		Symbol:      "BTCUSDT",
		ProductType: "USDT-FUTURES",
		MarginMode:  "isolated",
		MarginCoin:  "USDT",
		Side:        "buy",
		OrderType:   "limit",
		Size:        "0.01",
		Price:       "50000",
	}

	orderResp, err := client.PlaceOrder(orderRequest)
	if err != nil {
		log.Printf("Error placing order: %v", err)
	} else if orderResp.Code != "00000" {
		log.Printf("API Error: %s - %s", orderResp.Code, orderResp.Msg)
	} else {
		fmt.Printf("Order placed successfully! Order ID: %s\n", orderResp.Data.OrderID)
	}

	// Example 3: Get current positions
	fmt.Println("\n=== Getting Current Positions ===")
	positionReq := bitget_models.SinglePositionRequest{
		Symbol:      "BTCUSDT",
		ProductType: "USDT-FUTURES",
		MarginCoin:  "USDT",
	}

	positionResp, err := client.GetSinglePosition(positionReq)
	if err != nil {
		log.Printf("Error getting position: %v", err)
	} else {
		fmt.Printf("Current Positions: %+v\n", positionResp.Data)
	}

	// Example 4: Get pending orders
	fmt.Println("\n=== Getting Pending Orders ===")
	pendingReq := bitget_models.OrdersPendingRequest{
		Symbol:      "BTCUSDT",
		ProductType: "USDT-FUTURES",
	}

	pendingResp, err := client.GetPendingOrders(pendingReq)
	if err != nil {
		log.Printf("Error getting pending orders: %v", err)
	} else {
		fmt.Printf("Pending Orders Count: %d\n", len(pendingResp.Data.OrderList))
		for i, order := range pendingResp.Data.OrderList {
			fmt.Printf("Order %d: ID=%s, Side=%s, Size=%s, Price=%s\n",
				i+1, order.OrderID, order.Side, order.Size, order.Price)
		}
	}

	// Example 5: Set leverage
	fmt.Println("\n=== Setting Leverage ===")
	leverageReq := bitget_models.SetLeverageRequest{
		Symbol:      "BTCUSDT",
		ProductType: "USDT-FUTURES",
		MarginCoin:  "USDT",
		Leverage:    "10",
	}

	leverageResp, err := client.SetLeverage(leverageReq)
	if err != nil {
		log.Printf("Error setting leverage: %v", err)
	} else {
		fmt.Printf("Leverage set successfully: %+v\n", leverageResp.Data)
	}
}
