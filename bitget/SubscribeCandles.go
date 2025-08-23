package bitget

// Adjust the import path according to your project structure.

// SubscribeCandles connects to the given WebSocket endpoint, subscribes to the candle channel,
// and then calls the provided callback for each received message.
// func SubscribeCandles(wsURL string, subReq bitget_models.CandleSubscriptionRequest, callback func(pushData bitget_models.CandlePushData)) error {
// 	// Connect to the WebSocket.
// 	conn, _, err := websocket.DefaultDialer.Dial(wsURL, nil)
// 	if err != nil {
// 		return fmt.Errorf("failed to dial websocket: %w", err)
// 	}
// 	defer conn.Close()
// 	log.Println("Connected to WebSocket.")

// 	// Marshal and send the subscription message.
// 	msgBytes, err := json.Marshal(subReq)
// 	if err != nil {
// 		return fmt.Errorf("failed to marshal subscription request: %w", err)
// 	}

// 	if err := conn.WriteMessage(websocket.TextMessage, msgBytes); err != nil {
// 		return fmt.Errorf("failed to send subscription request: %w", err)
// 	}
// 	log.Println("Subscription request sent.")

// 	// Listen for incoming messages.
// 	for {
// 		_, message, err := conn.ReadMessage()
// 		if err != nil {
// 			return fmt.Errorf("failed to read message: %w", err)
// 		}

// 		// Optionally log raw message for debugging.
// 		log.Printf("Received raw message: %s", message)

// 		// Unmarshal the JSON message.
// 		var pushData bitget_models.CandlePushData
// 		if err := json.Unmarshal(message, &pushData); err != nil {
// 			log.Printf("Failed to unmarshal message: %v", err)
// 			continue
// 		}

// 		// Call the callback with the parsed push data.
// 		callback(pushData)

// 		// Sleep briefly if needed.
// 		time.Sleep(500 * time.Millisecond)
// 	}
// }
