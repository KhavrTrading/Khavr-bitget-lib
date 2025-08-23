// Package main demonstrates advanced trading operations using the Khavr Bitget library.
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
	"time"

	"github.com/KhavrTrading/Khavr-bitget-lib/bitget"
	"github.com/KhavrTrading/Khavr-bitget-lib/bitget/bitget_models"
)

func main() {
	// Create a new Bitget client with your credentials
	client := bitget.NewBitgetClient("your_api_key", "your_api_secret", "your_passphrase")

	symbol := "BTCUSDT"
	productType := "USDT-FUTURES"
	marginCoin := "USDT"

	fmt.Printf("=== Advanced Trading Examples for %s ===\n", symbol)

	// Example 1: Batch place multiple orders
	fmt.Println("\n1. Batch Placing Multiple Orders")
	batchOrderReq := bitget_models.BatchOrderRequest{
		Symbol:      symbol,
		ProductType: productType,
		MarginMode:  "isolated",
		MarginCoin:  marginCoin,
		OrderList: []bitget_models.BatchOrderOrderItem{
			{
				Side:      "buy",
				OrderType: "limit",
				Size:      "0.01",
				Price:     "45000",
			},
			{
				Side:      "buy",
				OrderType: "limit",
				Size:      "0.02",
				Price:     "44000",
			},
			{
				Side:      "sell",
				OrderType: "limit",
				Size:      "0.01",
				Price:     "55000",
			},
		},
	}

	batchResp, err := client.BatchPlaceOrder(batchOrderReq)
	if err != nil {
		log.Printf("Error batch placing orders: %v", err)
	} else {
		fmt.Printf("Batch order response: %+v\n", batchResp.Data)
		fmt.Printf("Success count: %d, Failure count: %d\n",
			len(batchResp.Data.SuccessList), len(batchResp.Data.FailureList))
	}

	// Example 2: Place TP/SL order
	fmt.Println("\n2. Placing Take Profit / Stop Loss Order")
	tpslReq := bitget_models.PlaceTpslOrderRequest{
		MarginCoin:   marginCoin,
		ProductType:  productType,
		Symbol:       symbol,
		PlanType:     "profit_plan", // Take profit plan
		TriggerPrice: "60000",
		Size:         "0.01",
		HoldSide:     "long", // Position side: long or short
		TriggerType:  "mark_price",
	}

	tpslResp, err := client.PlaceTpslOrder(tpslReq)
	if err != nil {
		log.Printf("Error placing TP/SL order: %v", err)
	} else {
		fmt.Printf("TP/SL order placed: %+v\n", tpslResp.Data)
	}

	// Example 3: Modify existing order
	fmt.Println("\n3. Modifying Order (if you have an order ID)")
	// Note: You would need a real order ID here
	modifyReq := bitgetlib.ModifyOrderRequest{
		Symbol:      symbol,
		ProductType: productType,
		OrderId:     "example_order_id", // Replace with real order ID
		NewSize:     "0.02",
		NewPrice:    "46000",
	}

	modifyResp, err := client.ModifyOrder(modifyReq)
	if err != nil {
		log.Printf("Error modifying order: %v", err)
	} else {
		fmt.Printf("Order modified: %+v\n", modifyResp.Data)
	}

	// Example 4: Get order history
	fmt.Println("\n4. Getting Order History")
	historyReq := bitgetlib.OrdersHistoryRequest{
		Symbol:      symbol,
		ProductType: productType,
		StartTime:   fmt.Sprintf("%d", time.Now().Add(-24*time.Hour).UnixMilli()),
		EndTime:     fmt.Sprintf("%d", time.Now().UnixMilli()),
		Limit:       "20",
	}

	historyResp, err := client.GetOrdersHistory(historyReq)
	if err != nil {
		log.Printf("Error getting order history: %v", err)
	} else {
		fmt.Printf("Order history count: %d\n", len(historyResp.Data.OrderList))
		for i, order := range historyResp.Data.OrderList {
			if i < 3 { // Show first 3 orders
				fmt.Printf("  Order %d: ID=%s, Side=%s, Size=%s, Price=%s, Status=%s\n",
					i+1, order.OrderID, order.Side, order.Size, order.Price, order.State)
			}
		}
	}

	// Example 5: Get order fills
	fmt.Println("\n5. Getting Order Fills")
	fillsReq := bitgetlib.OrderFillsRequest{
		Symbol:      symbol,
		ProductType: productType,
		StartTime:   fmt.Sprintf("%d", time.Now().Add(-24*time.Hour).UnixMilli()),
		EndTime:     fmt.Sprintf("%d", time.Now().UnixMilli()),
		Limit:       "20",
	}

	fillsResp, err := client.GetOrderFills(fillsReq)
	if err != nil {
		log.Printf("Error getting order fills: %v", err)
	} else {
		fmt.Printf("Order fills count: %d\n", len(fillsResp.Data.FillList))
		for i, fill := range fillsResp.Data.FillList {
			if i < 3 { // Show first 3 fills
				fmt.Printf("  Fill %d: OrderID=%s, Size=%s, Price=%s, Fee=%s\n",
					i+1, fill.OrderID, fill.Size, fill.Price, fill.Fee)
			}
		}
	}

	// Example 6: Close all positions for a symbol
	fmt.Println("\n6. Closing Positions (if any exist)")
	closeReq := bitgetlib.ClosePositionsRequest{
		Symbol:      symbol,
		ProductType: productType,
		HoldSide:    "long", // Optional: specify position side
	}

	closeResp, err := client.ClosePositions(closeReq)
	if err != nil {
		log.Printf("Error closing positions: %v", err)
	} else {
		fmt.Printf("Close positions response: %+v\n", closeResp.Data)
		fmt.Printf("Success count: %d, Failure count: %d\n",
			len(closeResp.Data.SuccessList), len(closeResp.Data.FailureList))
	}

	// Example 7: Cancel all orders for a symbol
	fmt.Println("\n7. Canceling All Orders")
	cancelAllReq := bitgetlib.CancelAllOrdersRequest{
		ProductType: productType,
		MarginCoin:  marginCoin,
	}

	cancelAllResp, err := client.CancelAllOrders(cancelAllReq)
	if err != nil {
		log.Printf("Error canceling all orders: %v", err)
	} else {
		fmt.Printf("Cancel all orders response: %+v\n", cancelAllResp.Data)
	}

	// Example 8: Get position history
	fmt.Println("\n8. Getting Position History")
	posHistoryReq := bitgetlib.PositionHistoryRequest{
		Symbol:      symbol,
		ProductType: productType,
		StartTime:   fmt.Sprintf("%d", time.Now().Add(-7*24*time.Hour).UnixMilli()),
		EndTime:     fmt.Sprintf("%d", time.Now().UnixMilli()),
		Limit:       "10",
	}

	posHistoryResp, err := client.GetPositionHistory(posHistoryReq)
	if err != nil {
		log.Printf("Error getting position history: %v", err)
	} else {
		fmt.Printf("Position history count: %d\n", len(posHistoryResp.Data.List))
		for i, position := range posHistoryResp.Data.List {
			if i < 2 { // Show first 2 positions
				fmt.Printf("  Position %d: Symbol=%s, Size=%s, OpenTime=%s, CloseTime=%s\n",
					i+1, position.Symbol, position.Size, position.OpenTime, position.CloseTime)
			}
		}
	}

	fmt.Println("\n=== Trading Examples Complete ===")
}
