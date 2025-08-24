# Khavr Bitget Go Library

[![Go Reference](https://pkg.go.dev/badge/github.com/KhavrTrading/Khavr-bitget-lib.svg)](https://pkg.go.dev/github.com/KhavrTrading/Khavr-bitget-lib)
[![Go Report Card](https://goreportcard.com/badge/github.com/KhavrTrading/Khavr-bitget-lib)](https://goreportcard.com/report/github.com/KhavrTrading/Khavr-bitget-lib)

A comprehensive Go client library for the Bitget cryptocurrency exchange API. This library provides both REST API and WebSocket functionality for futures and spot trading with complete type safety.

## Features

- 🚀 **Complete REST API Coverage**: All major Bitget API endpoints
- 📡 **WebSocket Support**: Real-time data streams and subscriptions
- 🔒 **Type Safe**: Full Go type definitions for all API requests and responses
- ⚡ **Rate Limiting**: Built-in rate limiting to respect API limits
- 🛡️ **Error Handling**: Comprehensive error handling and response validation
- 📊 **Futures Trading**: Support for USDT-FUTURES, COIN-FUTURES, and more
- 💼 **Account Management**: Position management, leverage adjustment, account details

## Installation

```bash
go get github.com/KhavrTrading/Khavr-bitget-lib
```

## Quick Start

### REST API Client

```go
package main

import (
    "fmt"
    "log"
    
    "github.com/KhavrTrading/Khavr-bitget-lib"
)

func main() {
    // Create a new Bitget client
    client := bitgetlib.NewBitgetClient("your_api_key", "your_api_secret", "your_passphrase")
    
    // Place a limit order
    orderRequest := bitgetlib.PlaceOrderRequest{
        Symbol:      "BTCUSDT",
        ProductType: bitgetlib.ProductTypeUSDTFutures,
        MarginMode:  bitgetlib.MarginModeIsolated,
        MarginCoin:  "USDT",
        Side:        bitgetlib.SideBuy,
        OrderType:   bitgetlib.OrderTypeLimit,
        Size:        "0.01",
        Price:       "50000",
    }
    
    response, err := client.PlaceOrder(orderRequest)
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("Order placed: %+v\n", response)
}
```

### WebSocket Client

```go
package main

import (
    "log"
    "time"
    
    "github.com/KhavrTrading/Khavr-bitget-lib"
)

func main() {
    // Create WebSocket client
    wsClient := bitgetlib.NewBitgetWsClient("", "", "") // Empty for public endpoints
    
    // Connect to WebSocket
    err := wsClient.Connect()
    if err != nil {
        log.Fatal(err)
    }
    defer wsClient.Close()
    
    // Subscribe to candlestick data
    subscription := bitgetlib.CandleSubscriptionRequest{
        Op: "subscribe",
        Args: []bitgetlib.CandleSubscriptionArg{
            {
                InstType: bitgetlib.ProductTypeUSDTFutures,
                Channel:  "candle1m",
                InstId:   "BTCUSDT",
            },
        },
    }
    
    err = wsClient.SubscribeCandles(subscription, func(data bitgetlib.CandlePushData) {
        fmt.Printf("Received candle data: %+v\n", data)
    })
    if err != nil {
        log.Fatal(err)
    }
    
    // Keep connection alive
    time.Sleep(30 * time.Second)
}
```

## API Coverage

### Account & Position Management
- ✅ Get Account Details
- ✅ Get Single Position
- ✅ Get Position History
- ✅ Close Positions
- ✅ Set Leverage

### Order Management
- ✅ Place Order
- ✅ Cancel Order
- ✅ Modify Order
- ✅ Batch Place Orders
- ✅ Batch Cancel Orders
- ✅ Cancel All Orders
- ✅ Get Order Details
- ✅ Get Pending Orders
- ✅ Get Order History
- ✅ Get Order Fills

### Advanced Orders
- ✅ Place TP/SL Orders
- ✅ Simultaneous Position TP/SL Orders

### Market Data
- ✅ Get Ticker Information
- ✅ WebSocket Candlestick Data

### Transaction History
- ✅ Get Transaction History
- ✅ Get Order Fill History

## Type Definitions

The library provides complete type definitions for all API interactions:

```go
// All request/response types are available
type PlaceOrderRequest struct {
    Symbol      string `json:"symbol"`
    ProductType string `json:"productType"`
    MarginMode  string `json:"marginMode"`
    MarginCoin  string `json:"marginCoin"`
    Side        string `json:"side"`
    OrderType   string `json:"orderType"`
    Size        string `json:"size"`
    Price       string `json:"price,omitempty"`
    // ... and more fields
}

type PlaceOrderResponse struct {
    Code        string    `json:"code"`
    Data        OrderData `json:"data"`
    Msg         string    `json:"msg"`
    RequestTime int64     `json:"requestTime"`
}
```

## Constants

The library provides convenient constants for common values:

```go
// Product Types
bitgetlib.ProductTypeUSDTFutures  // "USDT-FUTURES"
bitgetlib.ProductTypeCOINFutures  // "COIN-FUTURES"
bitgetlib.ProductTypeUSDCFutures  // "USDC-FUTURES"

// Order Sides
bitgetlib.SideBuy   // "buy"
bitgetlib.SideSell  // "sell"

// Order Types
bitgetlib.OrderTypeLimit   // "limit"
bitgetlib.OrderTypeMarket  // "market"

// Margin Modes
bitgetlib.MarginModeIsolated // "isolated"
bitgetlib.MarginModeCrossed  // "crossed"
```

## Error Handling

The library provides structured error handling:

```go
response, err := client.PlaceOrder(request)
if err != nil {
    // Handle network or API errors
    log.Printf("Error: %v", err)
    return
}

if response.Code != "00000" {
    // Handle API-level errors
    log.Printf("API Error: %s - %s", response.Code, response.Msg)
    return
}

// Success case
fmt.Printf("Order ID: %s", response.Data.OrderID)
```

## Rate Limiting

The library includes built-in rate limiting to respect Bitget's API limits:

- IP-based rate limiting
- UID-based rate limiting for private endpoints
- Automatic retry mechanisms
- Connection management for WebSocket

## Requirements

- Go 1.21 or higher
- Valid Bitget API credentials (for private endpoints)

## Getting API Credentials

1. Sign up at [Bitget](https://bonus.bitget.com/TNH1AZ)
2. Go to API Management in your account settings
3. Create a new API key with appropriate permissions
4. Save your API Key, Secret Key, and Passphrase securely

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## Support

If you encounter any issues or have questions, please open an issue on GitHub.

## Disclaimer

This library is not officially affiliated with Bitget. Use at your own risk. Always test thoroughly before using in production environments.

## Changelog

### v0.1.0
- Initial release
- Complete REST API coverage
- WebSocket support for real-time data
- Type-safe Go implementations
- Built-in rate limiting
- Comprehensive error handling
