package request

import (
	"encoding/json"
	"net/http"

	"toggler.in/internal/http/response"
	"toggler.in/internal/validator"

	"go.uber.org/zap"
)

type Reader struct {
	log *zap.Logger
	jw *response.JSONWriter
	validator *validator.Validator
}

// NewReader returns a new instance of Reader
func NewReader(log *zap.Logger, jw *response.JSONWriter, validator *validator.Validator) *Reader {
	return &Reader{
		log:       log,
		jw:        jw,
		validator: validator,
	}
}

// ReadJSONAndValidate reads a json request body into the given struct and the validates the struct data
func (read *Reader) ReadJSONAndValidate(w http.ResponseWriter, r *http.Request, v interface{}) bool {
	err := read.ReadJSONRequest(r, &v)
	if err != nil {
		parseErr := read.HandleParseError(err)
		read.jw.BadRequest(w, r, parseErr)
		return false
	}

	ve := read.validate(v)
	if ve != nil {
		read.jw.UnprocessableEntity(w, r, ve)
		return false
	}

	return true
}

// ReadJSONRequest reads a json request body into the given struct
func (read *Reader) ReadJSONRequest(r *http.Request, v interface{}) error {
	return json.NewDecoder(r.Body).Decode(v)
}

//validate functions uses the validator to test issues with the given data
func (read *Reader) validate(v interface{}) response.APIError {
	fields := read.validator.IsValidStruct(v)
	if len(fields) == 0 {
		return nil
	}

	return &ValidationError{
		ErrData: fields,
	}
}