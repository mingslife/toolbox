package common

import "fmt"

type BaseError interface {
	error
	StatusCode() int
	ErrorCode() uint
	ErrorText() string
}

type baseError struct {
	statusCode int
	errorCode  uint // 10001~99999 前两位模块号，后三位错误号
	errorText  string
}

const (
	errCodeUndefined  uint = 10000 // 未声明错误编码
	errCodeValidation uint = 10001 // 验证错误
)

func (e *baseError) Error() string {
	return fmt.Sprintf("[%d] %s", e.errorCode, e.errorText)
}

func (e *baseError) StatusCode() int {
	return e.statusCode
}

func (e *baseError) ErrorCode() uint {
	return e.errorCode
}

func (e *baseError) ErrorText() string {
	return e.errorText
}
