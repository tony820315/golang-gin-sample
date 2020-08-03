package resp

import (
	"fmt"
	"net/http"
)

const (
	StatusNotFound = iota
)

var statusText = map[int]string{
	StatusNotFound: "No todo found!",
}

type ExtendError interface {
	GetAttachment() interface{}
	GetCode() int
	GetMessage() string
	Error() string
}

type BaseError struct {
	Code       int    `json:"status"`
	Message    string `json:"message"`
	attachment interface{}
}

func NewBaseError(code int, message string, attachment interface{}) ExtendError {
	if len(message) == 0 {
		if 0 <= code && code < http.StatusContinue {
			message = statusText[code]
		} else if http.StatusContinue <= code && code <= http.StatusNetworkAuthenticationRequired {
			message = http.StatusText(code)
		}
	}
	return &BaseError{Code: code, Message: message, attachment: attachment}
}

func (b *BaseError) Error() string {
	return fmt.Sprintf("Code:%d: Message:%s", b.Code, b.Message)
}

func (b *BaseError) GetAttachment() interface{} {
	return b.attachment
}

func (b *BaseError) GetCode() int {
	return b.Code
}

func (b *BaseError) GetMessage() string {
	return b.Message
}
