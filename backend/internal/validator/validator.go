package validator

import (
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

//Validator has functions for validating struct and variables
type Validator struct {
	log      *zap.Logger
	validate *validator.Validate
}

//New returns a new Validator
func New(log *zap.Logger) *Validator {
	v := buildValidator()

	return &Validator{
		log:      log,
		validate: v,
	}
}

// IsValidStruct validates a struct and returns a list of invalid fields or nil
func (v *Validator) IsValidStruct(data interface{}) []string {
	err := v.validate.Struct(data)
	if err != nil {
		return HandleValidationError(v.log, err)
	}

	return nil
}

type (
	validatorFn      func(fl validator.FieldLevel) bool
	stringModifierFn func(str string) string
)

// AddCustomValidator adds a custom validator tag
func (v *Validator) AddCustomValidator(name string, f validatorFn) error {
	err := v.validate.RegisterValidation(name, validator.Func(f))
	if err != nil {
		v.log.Error("failed to add custom validator", zap.Error(err), zap.String("custom_validator", name))
	}

	return nil
}

// AddStructLevelValidation adds a custom struct level validation
func (v *Validator) AddStructLevelValidation(fn validator.StructLevelFunc, structType interface{}) {
	v.validate.RegisterStructValidation(fn, structType)
}

// AddStringModifier adds a tag that modifies the string value
func (v *Validator) AddStringModifier(name string, fn stringModifierFn) error {
	return v.AddCustomValidator(name, func(fl validator.FieldLevel) bool {
		if fl.Field().Type().String() == "string" {
			str := fn(fl.Field().String())
			fl.Field().SetString(str)
		}
		return true
	})
}

// buildValidator builds the validator.Validate
func buildValidator() *validator.Validate {
	v := validator.New()

	// register function to get tag name from json tags.
	v.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "" {
			name = strings.SplitN(fld.Tag.Get("schema"), ",", 2)[0]
		}
		if name == "-" {
			return ""
		}
		return name
	})

	return v
}

