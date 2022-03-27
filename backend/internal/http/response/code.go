package response

type ErrorCode int

const (
	// DefaultErrorCode - in case other codes are irrelevant
	DefaultErrorCode ErrorCode = 500

	// EmptyRequestBody - when the request body is empty
	EmptyRequestBody ErrorCode = 1000
	// InvalidJSON - when the json data in request in invalid
	InvalidJSON ErrorCode = 1001
	// InvalidJSONField - when a json field in request in invalid
	InvalidJSONField ErrorCode = 1002
	// UnKnownJSONField - when json request body contains unnecessary field for the request
	UnKnownJSONField ErrorCode = 1003
	// RequestSizeExceeds - when request body's size is large
	RequestSizeExceeds ErrorCode = 1004
	// UnknownParseError - when the parse error is none of the above categories
	UnknownParseError ErrorCode = 1005
	// ValidationFailed - when the parse error is none of the above categories
	ValidationFailed ErrorCode = 1006
)