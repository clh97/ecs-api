package constants

import "errors"

var (
	// ErrNotFound means a generic resource could not be found
	ErrNotFound = errors.New("NOT_FOUND")
)
