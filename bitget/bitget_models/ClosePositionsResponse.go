package bitget_models

// ClosePositionsResponse represents the API response for closing positions.
type ClosePositionsResponse struct {
	Code        string             `json:"code"`        // "00000" indicates success.
	Msg         string             `json:"msg"`         // Response message.
	RequestTime int64              `json:"requestTime"` // Timestamp of the request.
	Data        ClosePositionsData `json:"data"`        // Data containing success and failure lists.
}
