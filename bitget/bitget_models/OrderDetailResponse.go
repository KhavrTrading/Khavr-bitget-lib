package bitget_models

type OrderDetailResponse struct {
	Code        string          `json:"code"`        // "00000" indicates success.
	Msg         string          `json:"msg"`         // Response message.
	RequestTime int64           `json:"requestTime"` // Timestamp of the request.
	Data        OrderDetailData `json:"data"`        // Order detail data.
}
