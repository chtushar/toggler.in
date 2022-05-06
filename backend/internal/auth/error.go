package auth

import (
	"toggler.in/internal/http/response"
)

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


type UserNotFoundError struct {}

func (d *UserNotFoundError) Message() string {
	return "User not found"
}

func (d *UserNotFoundError) Code() response.ErrorCode {
	return response.NotFound
}

func (d *UserNotFoundError) Data() interface{} {
	return nil
}

type IncorrectPasswordError struct {}

func (d *IncorrectPasswordError) Message() string {
	return "Incorrect password"
}

func (d *IncorrectPasswordError) Code() response.ErrorCode {
	return response.Unauthorized
}

func (d *IncorrectPasswordError) Data() interface{} {
	return nil
}