package authorizer

import "fmt"

// AuthorizationError represents an authorization error
type AuthorizationError struct {
	err error
}

// Error implements the standard Go error interface
func (err AuthorizationError) Error() string {
	return err.err.Error()
}

// Errorf creates a new formatted validator error
func Errorf(format string, v ...interface{}) AuthorizationError {
	return AuthorizationError{
		err: fmt.Errorf(format, v...),
	}
}

// IsAuthorizationError returns true if the given error represents
// an authorization error and originates from an authorizer module
func IsAuthorizationError(err error) bool {
	_, isAuthorizationError := err.(AuthorizationError)
	return isAuthorizationError
}
