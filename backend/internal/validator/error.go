package validator

import (
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

//HandleValidationError takes the validation error and read the failed fields
func HandleValidationError(log *zap.Logger, err error) []string {
	fields := make([]string, 0)

	if invalidErr, ok := err.(*validator.InvalidValidationError); ok {
		log.Panic("error in validation", zap.Error(invalidErr))
		return nil
	}

	for _, validationErr := range err.(validator.ValidationErrors) {
		fields = append(fields, validationErr.Field())
	}

	return fields
}