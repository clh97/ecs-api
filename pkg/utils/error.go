package utils

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

// ValidationError is the base structure for validation errors
type ValidationError struct {
	Message       string
	Specification string `json:",omitempty"`
	Field         string
}

// ParseValidationErrors will return a slice containing user-friendly ValidationError structs
func ParseValidationErrors(errors validator.ValidationErrors) (validationErrors []ValidationError) {
	for _, err := range errors {
		var errorMessage string
		var errorSpecification string

		// empty value means we should say that the field is required
		// while a non-empty value means the information is invalid
		if err.Value() == "" {
			errorMessage = fmt.Sprintf("%s is %s", err.Field(), err.Tag())
		} else {
			errorMessage = fmt.Sprintf("%s is invalid", err.Field())
		}

		if err.ActualTag() == "min" {
			errorSpecification = fmt.Sprintf("%s should be at least %s characters", err.Field(), err.Param())
		}

		if err.ActualTag() == "max" {
			errorSpecification = fmt.Sprintf("%s should be at most %s characters", err.Field(), err.Param())
		}

		validationError := ValidationError{
			Message:       errorMessage,
			Specification: errorSpecification,
			Field:         err.Field(),
		}

		validationErrors = append(validationErrors, validationError)
	}
	return validationErrors
}
