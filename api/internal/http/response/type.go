package response

import "time"

type (
	// APIResponse is to be implemented for sending a response to the client
	APIResponse interface {
		Data() interface{}
	}

	// APIError is to be implemented for sending an error to the client
	APIError interface {
		Message() string
		Code() ErrorCode
		Data() interface{}
	}
)

type (
	response struct {
		RequestID string      `json:"request_id"`
		Timestamp time.Time   `json:"timestamp"`
		URI       string      `json:"uri"`
		Success   bool        `json:"success"`
		Data      interface{} `json:"data"`
		Error     *errorWrap  `json:"error"`
	}

	errorWrap struct {
		Message string      `json:"message"`
		Code    ErrorCode   `json:"code"`
		Data    interface{} `json:"data"`
	}

	defaultErr struct{}
)

func (d *defaultErr) Message() string {
	return "An unknown error occurred. If this persists, please contact us."
}

func (d *defaultErr) Code() ErrorCode {
	return DefaultErrorCode
}

func (d *defaultErr) Data() interface{} {
	return nil
}