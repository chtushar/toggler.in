package teams

import "toggler.in/internal/http/response"

type InternalError struct {}

func (d *InternalError) Message() string {
	return "Error while hashing password"
}

func (d *InternalError) Code() response.ErrorCode {
	return response.DefaultErrorCode
}

func (d *InternalError) Data() interface{} {
	return nil
}