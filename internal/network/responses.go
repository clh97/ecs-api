package network

// RequestError is the structure returned in http requests when an authentication error occurs
type RequestError struct {
	Message  string
	Redirect string `json:",omitempty"`
	Error    string
}

var (
	// MissingAuthHeaderError occurs when Authentication header is not defined
	MissingAuthHeaderError = RequestError{
		Message:  "Missing Authentication header",
		Redirect: "/api/v1/login",
	}
)

// JSONValidationError occurs when JSON data binding is not valid
func JSONValidationError(validationError error, message string) RequestError {
	return RequestError{
		Message: message,
		Error:   validationError.Error(),
	}
}
