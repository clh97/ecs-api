package constants

import (
	"time"

	"github.com/clh97/ecs/pkg/utils"
	"github.com/go-playground/validator/v10"
)

// THTTPErrorResponse is the structure returned in http requests when an error occurs
type THTTPErrorResponse struct {
	Message          string
	Redirect         string                  `json:",omitempty"`
	Error            interface{}             `json:"-"`
	ValidationErrors []utils.ValidationError `json:",omitempty"`
	Timestamp        time.Time
	Success          bool
}

// THTTPResponse is the strucure returned in successful http requests
type THTTPResponse struct {
	Message   string      `json:",omitempty"`
	Data      interface{} `json:",omitempty"`
	Timestamp time.Time
	Success   bool
}

// HTTPResponse is the constructor for a successful http response
func HTTPResponse(payload interface{}, message string) THTTPResponse {
	return THTTPResponse{
		Message:   message,
		Timestamp: time.Now(),
		Data:      payload,
		Success:   true,
	}
}

// HTTPErrorResponse is the constructor for a unsuccessful http response
func HTTPErrorResponse(err error, message string, redirect string) THTTPErrorResponse {
	var responseValidationErrors []utils.ValidationError

	if validationError, ok := err.(validator.FieldError); ok {
		responseValidationError := utils.ParseValidationError(validationError)
		responseValidationErrors = append(responseValidationErrors, responseValidationError)
	}

	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		responseValidationErrors = utils.ParseValidationErrors(validationErrors)
	}

	return THTTPErrorResponse{
		Message:          message,
		Error:            err,
		Redirect:         redirect,
		ValidationErrors: responseValidationErrors,
		Timestamp:        time.Now(),
		Success:          false,
	}
}
