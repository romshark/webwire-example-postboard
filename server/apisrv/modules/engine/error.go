package engine

import (
	"fmt"
)

// ErrorCode represents an engine error code
type ErrorCode int

const (
	_ ErrorCode = iota

	// ErrInternal indicates that an internal error occurred
	ErrInternal

	// ErrInvalidRequest indicates that the received request is invalid
	// and cannot be processed
	ErrInvalidRequest

	// ErrUnauthorized indicates that the client is not authorized
	// to execute any given operation
	ErrUnauthorized

	// ErrWrongCredentials indicates that the provided login credentials
	// are incorrect
	ErrWrongCredentials

	// ErrUsernameReserved indicates that a username is already reserved
	// by another user's account
	ErrUsernameReserved

	// ErrMessageNotFound indicates that the targeted message doesn't exist
	ErrMessageNotFound

	// ErrMessageAlreadyExists indicates that a message
	// with a similar identifier already exists
	ErrMessageAlreadyExists

	// ErrUserNotFound indicates that the targeted user account doesn't exist
	ErrUserNotFound
)

// String stringifies the engine error code
func (errc ErrorCode) String() string {
	switch errc {
	case ErrInternal:
		return "INTERNAL_ERR"
	case ErrInvalidRequest:
		return "INVALID_REQ"
	case ErrUnauthorized:
		return "UNAUTHORIZED"
	case ErrWrongCredentials:
		return "WRONG_CREDENTIALS"
	case ErrUsernameReserved:
		return "USERNAME_RESERVED"
	case ErrMessageNotFound:
		return "MSG_NOT_FOUND"
	case ErrMessageAlreadyExists:
		return "MSG_EXISTS"
	case ErrUserNotFound:
		return "USER_NOT_FOUND"
	}
	return ""
}

// EngineError represents an implementation of the standard Go error interface.
// It represents an error returned by the engine implementation
type EngineError struct {
	Code ErrorCode
}

// Error implements the standard Go error interface
func (err EngineError) Error() string {
	return fmt.Sprintf("engine error (%s)", err.Code.String())
}

// NewError creates a new engine error with the given code
func NewError(code ErrorCode) error {
	return EngineError{
		Code: code,
	}
}

// ToEngineError returns an engine error if the given error interface
// houses one, otherwise returns nil
func ToEngineError(err error) *EngineError {
	switch engineErr := err.(type) {
	case EngineError:
		return &engineErr
	case *EngineError:
		return engineErr
	default:
		return nil
	}
}
