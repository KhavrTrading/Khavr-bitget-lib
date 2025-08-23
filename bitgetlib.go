// Package bitgetlib provides a Go client library for the Bitget cryptocurrency exchange API.
// It includes both REST API and WebSocket functionality for futures and spot trading.
//
// The library provides three main components:
//   - BitgetClient: REST API client for synchronous operations
//   - BitgetWsClient: WebSocket client for real-time data streams
//   - BitgetModels: Complete type definitions for all API requests and responses
//
// Example usage:
//
//	import "github.com/KhavrTrading/Khavr-bitget-lib"
//
//	// Create REST client
//	client := bitgetlib.NewBitgetClient("api_key", "api_secret", "passphrase")
//
//	// Create WebSocket client
//	wsClient := bitgetlib.NewBitgetWsClient("api_key", "api_secret", "passphrase")
//
//	// Use models for type safety
//	request := bitgetlib.PlaceOrderRequest{
//	    Symbol:   "BTCUSDT",
//	    Side:     "buy",
//	    OrderType: "limit",
//	    Size:     "0.01",
//	    Price:    "50000",
//	}
package bitgetlib

import (
	"github.com/KhavrTrading/Khavr-bitget-lib/bitget"
	"github.com/KhavrTrading/Khavr-bitget-lib/bitget/bitget_models"
)

// BitgetClient is the main REST API client for Bitget exchange.
// It provides methods for all trading operations including orders, positions, account management.
type BitgetClient = bitget.BitgetClient

// BitgetWsClient is the WebSocket client for real-time data streams from Bitget.
// It supports subscribing to candlestick data, order book updates, and other real-time feeds.
type BitgetWsClient = bitget.BitgetWsClient

// NewBitgetClient creates a new REST API client with the provided credentials.
//
// Parameters:
//   - apiKey: Your Bitget API key
//   - apiSecret: Your Bitget API secret
//   - passphrase: Your Bitget API passphrase
//
// Returns a configured BitgetClient ready for use.
func NewBitgetClient(apiKey, apiSecret, passphrase string) *BitgetClient {
	return bitget.NewBitgetClient(apiKey, apiSecret, passphrase)
}

// NewBitgetWsClient creates a new WebSocket client with the provided credentials.
//
// Parameters:
//   - apiKey: Your Bitget API key (may be empty for public endpoints)
//   - apiSecret: Your Bitget API secret (may be empty for public endpoints)
//   - passphrase: Your Bitget API passphrase (may be empty for public endpoints)
//
// Returns a configured BitgetWsClient ready for connection.
func NewBitgetWsClient(apiKey, apiSecret, passphrase string) *BitgetWsClient {
	return bitget.NewBitgetWsClient(apiKey, apiSecret, passphrase)
}

// BitgetModels contains all type definitions for Bitget API requests and responses.
// This allows users to import and use all model types directly.
type BitgetModels struct{}

// Re-export all commonly used model types for easy access
type (
	// Account models
	AccountRequest  = bitget_models.AccountRequest
	AccountResponse = bitget_models.AccountResponse
	AccountData     = bitget_models.AccountData

	// Order models
	PlaceOrderRequest   = bitget_models.PlaceOrderRequest
	PlaceOrderResponse  = bitget_models.PlaceOrderResponse
	OrderData           = bitget_models.OrderData
	CancelOrderRequest  = bitget_models.CancelOrderRequest
	CancelOrderResponse = bitget_models.CancelOrderResponse
	ModifyOrderRequest  = bitget_models.ModifyOrderRequest
	ModifyOrderResponse = bitget_models.ModifyOrderResponse
	OrderDetailRequest  = bitget_models.OrderDetailRequest
	OrderDetailResponse = bitget_models.OrderDetailResponse
	OrderDetailData     = bitget_models.OrderDetailData

	// Batch order models
	BatchOrderRequest       = bitget_models.BatchOrderRequest
	BatchPlaceOrderResponse = bitget_models.BatchPlaceOrderResponse
	BatchPlaceOrderData     = bitget_models.BatchPlaceOrderData
	BatchOrderOrderItem     = bitget_models.BatchOrderOrderItem
	BatchOrderSuccess       = bitget_models.BatchOrderSuccess
	BatchOrderFailure       = bitget_models.BatchOrderFailure

	BatchCancelOrderRequest  = bitget_models.BatchCancelOrderRequest
	BatchCancelOrderResponse = bitget_models.BatchCancelOrderResponse
	BatchCancelOrderData     = bitget_models.BatchCancelOrderData

	// Position models
	SinglePositionRequest  = bitget_models.SinglePositionRequest
	SinglePositionResponse = bitget_models.SinglePositionResponse
	SinglePosition         = bitget_models.SinglePosition

	PositionHistoryRequest  = bitget_models.PositionHistoryRequest
	PositionHistoryResponse = bitget_models.PositionHistoryResponse
	PositionHistoryData     = bitget_models.PositionHistoryData

	ClosePositionsRequest  = bitget_models.ClosePositionsRequest
	ClosePositionsResponse = bitget_models.ClosePositionsResponse
	ClosePositionsData     = bitget_models.ClosePositionsData
	ClosePositionsSuccess  = bitget_models.ClosePositionsSuccess
	ClosePositionsFailure  = bitget_models.ClosePositionsFailure

	// TPSL order models
	PlaceTpslOrderRequest  = bitget_models.PlaceTpslOrderRequest
	PlaceTpslOrderResponse = bitget_models.PlaceTpslOrderResponse
	PlaceTpslOrderData     = bitget_models.PlaceTpslOrderData

	SimultaneousPlacePosTpslOrderRequest  = bitget_models.PlacePosTpslOrderRequest
	SimultaneousPlacePosTpslOrderResponse = bitget_models.PlacePosTpslOrderResponse

	// History models
	HistoryTransactionRequest  = bitget_models.HistoryTransactionRequest
	HistoryTransactionResponse = bitget_models.HistoryTransactionResponse
	HistoryTransactionData     = bitget_models.HistoryTransactionData

	OrdersHistoryRequest  = bitget_models.OrdersHistoryRequest
	OrdersHistoryResponse = bitget_models.OrdersHistoryResponse
	OrdersHistoryData     = bitget_models.OrdersHistoryData

	OrdersPendingRequest  = bitget_models.OrdersPendingRequest
	OrdersPendingResponse = bitget_models.OrdersPendingResponse
	OrdersPendingData     = bitget_models.OrdersPendingData

	// Order fill models
	OrderFillsRequest    = bitget_models.OrderFillsRequest
	OrderFillsResponse   = bitget_models.OrderFillsResponse
	OrderFillsData       = bitget_models.OrderFillsData
	OrderFillHistoryData = bitget_models.OrderFillHistoryData

	// Market data models
	TickerRequest  = bitget_models.TickerRequest
	TickerResponse = bitget_models.TickerResponse
	TickerData     = bitget_models.TickerData

	// WebSocket models
	CandleSubscriptionRequest = bitget_models.CandleSubscriptionRequest
	CandleSubscriptionArg     = bitget_models.CandleSubscriptionArg
	CandlePushData            = bitget_models.CandlePushData

	// Leverage models
	SetLeverageRequest  = bitget_models.SetLeverageRequest
	SetLeverageResponse = bitget_models.SetLeverageResponse
	LeverageData        = bitget_models.LeverageData

	// Cancel all orders models
	CancelAllOrdersRequest  = bitget_models.CancelAllOrdersRequest
	CancelAllOrdersResponse = bitget_models.CancelAllOrdersResponse
	CancelAllOrdersData     = bitget_models.CancelAllOrdersData

	// Common types
	ErrorResponse = bitget_models.ErrorResponse
)

// Constants from bitget_models for convenience
const (
	// Product types
	ProductTypeUSDTFutures  = "USDT-FUTURES"
	ProductTypeCOINFutures  = "COIN-FUTURES"
	ProductTypeUSDCFutures  = "USDC-FUTURES"
	ProductTypeSUSDTFutures = "SUSDT-FUTURES"
	ProductTypeSCOINFutures = "SCOIN-FUTURES"
	ProductTypeSUSDCFutures = "SUSDC-FUTURES"

	// Order sides
	SideBuy  = "buy"
	SideSell = "sell"

	// Order types
	OrderTypeLimit  = "limit"
	OrderTypeMarket = "market"

	// Margin modes
	MarginModeIsolated = "isolated"
	MarginModeCrossed  = "crossed"
)
