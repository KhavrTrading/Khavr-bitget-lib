package bitget_models

type ParamBuilder interface {
	ToParams() map[string]string
}
