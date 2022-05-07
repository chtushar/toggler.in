package handlers

import "toggler.in/internal/http/response"

type UnauthorizedError struct {}

func (d *UnauthorizedError) Message() string {
	return "Unauthorized"
}

func (d *UnauthorizedError) Code() response.ErrorCode {
	return response.Unauthorized
}

func (d *UnauthorizedError) Data() interface{} {
	return nil
}