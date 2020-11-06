package constants

// ServiceError is the error structure returned by services
type ServiceError struct {
	HTTPStatus        int
	HTTPErrorResponse THTTPErrorResponse
}

// ServiceResult is the success structure returned by services
type ServiceResult struct {
	HTTPStatus   int
	HTTPResponse THTTPResponse
}
