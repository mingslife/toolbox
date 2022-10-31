package common

const (
	CodeOK uint64 = iota
	CodeError
)

type BaseDTO struct {
	Code    uint64 `json:"code"`
	Message string `json:"message"`
}
