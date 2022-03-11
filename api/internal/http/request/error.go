package request

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strings"

	"toggler.in/internal/http/response"
)

type (
	// ParseError struct defines the structure of the JSON parsing error
	ParseError struct {
		Msg     string             `json:"message"`
		ErrCode response.ErrorCode `json:"code"`
		ErrData interface{}        `json:"data"`
	}

	// ValidationError struct defines the structure of the Validation errors
	ValidationError struct {
		ErrData []string `json:"data"`
	}
)

func (p *ParseError) Message() string {
	return p.Msg
}

func (p *ParseError) Code() response.ErrorCode {
	return p.ErrCode
}

func (p *ParseError) Data() interface{} {
	return p.ErrData
}

func (v *ValidationError) Message() string {
	return "Invalid data"
}

func (v *ValidationError) Code() response.ErrorCode {
	return response.ValidationFailed
}

func (v *ValidationError) Data() interface{} {
	return v.ErrData
}

// HandleParseError checks the type of the error in request parsing and writes an appropriate response
func (read *Reader) HandleParseError(err error) *ParseError {
	var (
		syntaxError        *json.SyntaxError
		unmarshalTypeError *json.UnmarshalTypeError
	)

	switch {
	// Catch any syntax errors in the JSON and send an error Message
	// which interpolates the location of the problem to make it
	// easier for the client to fix.
	// In some circumstances Decode() may also return an
	// io.ErrUnexpectedEOF error for syntax errors in the JSON. There
	// is an open issue regarding this at
	// https://github.com/golang/go/issues/25956.
	case errors.As(err, &syntaxError), errors.Is(err, io.ErrUnexpectedEOF):
		return &ParseError{
			Msg:     "Request body contains badly-formed JSON",
			ErrCode: response.InvalidJSON,
			ErrData: map[string]interface{}{
				"offset": syntaxError.Offset,
			},
		}

	// Catch any type errors, like trying to assign a string in the
	// JSON request body to an int field in our Person struct. We can
	// interpolate the relevant field name and position into the error
	// Message to make it easier for the client to fix.
	case errors.As(err, &unmarshalTypeError):
		return &ParseError{
			Msg:     fmt.Sprintf("Request body contains an invalid value for the %q field", unmarshalTypeError.Field),
			ErrCode: response.InvalidJSON,
			ErrData: map[string]interface{}{
				"field": unmarshalTypeError.Field,
			},
		}

	// Catch the error caused by extra unexpected fields in the request
	// body. We extract the field name from the error Message and
	// interpolate it in our custom error Message. There is an open
	// issue at https://github.com/golang/go/issues/29035 regarding
	// turning this into a sentinel error.
	case strings.HasPrefix(err.Error(), "json: unknown field "):
		fieldName := strings.TrimPrefix(err.Error(), "json: unknown field ")

		return &ParseError{
			Msg:     fmt.Sprintf("Request body contains unknown field %s", fieldName),
			ErrCode: response.UnKnownJSONField,
			ErrData: map[string]interface{}{
				"field": fieldName,
			},
		}

	// An io.EOF error is returned by Decode() if the request body is
	// empty.
	case errors.Is(err, io.EOF):
		return &ParseError{
			Msg:     "Request body must not be empty",
			ErrCode: response.EmptyRequestBody,
			ErrData: nil,
		}

	// Catch the error caused by the request body being too large. Again
	// there is an open issue regarding turning this into a sentinel
	// error at https://github.com/golang/go/issues/30715.
	case err.Error() == "http: request body too large":
		return &ParseError{
			Msg:     "Request body must not be larger than 1MB",
			ErrCode: response.RequestSizeExceeds,
			ErrData: nil,
		}

	}

	return &ParseError{
		Msg:     "Unknown error occurred",
		ErrCode: response.UnknownParseError,
		ErrData: nil,
	}
}
