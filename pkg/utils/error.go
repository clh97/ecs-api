package utils

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

// ValidationError is the base structure for validation errors, contains user-friendly error data
type ValidationError struct {
	Message       string
	Specification string `json:",omitempty"`
	Field         string
}

func validateFieldError(err validator.FieldError) ValidationError {
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

	return validationError
}

// ParseValidationErrors will return a slice containing ValidationError structs
func ParseValidationErrors(errors validator.ValidationErrors) (validationErrors []ValidationError) {
	for _, err := range errors {
		validationError := validateFieldError(err)

		validationErrors = append(validationErrors, validationError)
	}
	return validationErrors
}

// ParseValidationError will return a single ValidationError instance
func ParseValidationError(err validator.FieldError) ValidationError {
	return validateFieldError(err)
}
