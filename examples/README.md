# Examples

This directory contains practical examples demonstrating how to use the Khavr Bitget Go Library.

## Available Examples

### 1. `rest_client_example.go`
Basic REST API usage examples including:
- Creating a client
- Getting account information
- Placing orders
- Checking positions
- Setting leverage

### 2. `websocket_example.go`
WebSocket client usage examples including:
- Connecting to WebSocket
- Subscribing to candlestick data
- Handling real-time data streams
- Multiple symbol subscriptions

### 3. `trading_example.go`
Advanced trading operations including:
- Batch order placement
- TP/SL orders
- Order modification
- Order history and fills
- Position management
- Order cancellation

## Running the Examples

1. First, ensure you have Go installed (version 1.21 or higher)

2. Set up the module dependencies:
```bash
cd examples
go mod tidy
```

3. Edit the examples to include your actual API credentials:
```go
client := bitgetlib.NewBitgetClient("your_api_key", "your_api_secret", "your_passphrase")
```

4. Run any example:
```bash
go run rest_client_example.go
go run websocket_example.go
go run trading_example.go
```

## Important Notes

⚠️ **Warning**: These examples use real API endpoints. When testing:
- Use testnet credentials if available
- Start with very small amounts
- Test with market/limit orders you're comfortable with
- Be aware of trading fees and risks

### API Credentials

To get your API credentials:
1. Sign up at [Bitget](https://www.bitget.com/)
2. Go to API Management in your account settings
3. Create a new API key with appropriate permissions
4. Replace the placeholder credentials in the examples

### WebSocket Example Notes

The WebSocket example:
- Uses public endpoints (no credentials needed for market data)
- Subscribes to real-time candlestick data
- Demonstrates graceful shutdown with Ctrl+C
- Shows how to handle multiple symbol subscriptions

### Error Handling

All examples include proper error handling patterns:
- Network-level errors (connection issues)
- API-level errors (invalid parameters, insufficient balance, etc.)
- Response validation

### Rate Limiting

The library includes built-in rate limiting, but be mindful of:
- IP-based limits
- UID-based limits
- Different limits for different endpoints

## Customization

Feel free to modify these examples for your specific use cases:
- Change trading pairs
- Adjust order sizes and prices
- Add more sophisticated logic
- Implement your own trading strategies

## Support

If you encounter issues with the examples:
1. Check that your API credentials are correct
2. Verify you have sufficient balance for trading operations
3. Ensure your API key has the required permissions
4. Check the Bitget API documentation for any updates

## License

These examples are provided as-is for educational purposes. Use at your own risk.
