package bitget_models

type ErrorResponse struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
}
