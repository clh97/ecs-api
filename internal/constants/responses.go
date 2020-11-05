package constants

import "time"

// RequestError is the structure returned in http requests when an error occurs
type RequestError struct {
	Message  string
	Redirect string      `json:",omitempty"`
	Error    interface{} `json:",omitempty"`
	Success  bool
}

// HTTPResponse is the strucure returned in successful http requests
type HTTPResponse struct {
	Message   string `json:",omitempty"`
	Data      interface{}
	Timestamp time.Time
	Success   bool
}

var (
	// MissingAuthHeaderError occurs when Authorization header is not defined
	MissingAuthHeaderError = RequestError{
		Message:  "Missing Authorization header",
		Redirect: "/api/v1/login",
		Success:  false,
	}
	// InvalidTokenError occurs when invalid token is provided
	InvalidTokenError = RequestError{
		Message:  "Invalid token provided",
		Redirect: "/api/v1/login",
		Success:  false,
	}
	// InternalError occurs when some internal method fails to run successfully
	InternalError = RequestError{
		Message: "Internal server error",
		Success: false,
	}
)

// Success is the http response for succesful requests
func Success(data interface{}, message string) HTTPResponse {
	return HTTPResponse{
		Message:   "Successfully executed request",
		Data:      data,
		Success:   true,
		Timestamp: time.Now(),
	}
}

// JSONValidationError occurs when JSON data binding is not valid
func JSONValidationError(validationError error, message string) RequestError {
	return RequestError{
		Message: message,
		Error:   validationError.Error(),
	}
}

// ResourceConflictError occurs when the specified resource already exists and should not be recreated
func ResourceConflictError(err error, message string) RequestError {
	return RequestError{
		Message: message,
	}
}

// ResourceCreated is the http response for successfully created resources
func ResourceCreated(data interface{}, message string) HTTPResponse {
	return HTTPResponse{
		Message:   message,
		Data:      data,
		Success:   true,
		Timestamp: time.Now(),
	}
}
