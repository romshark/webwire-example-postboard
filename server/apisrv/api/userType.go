package api

import (
	"strings"

	"github.com/pkg/errors"
)

// UserType represents the role of a user
type UserType int

const (
	// UtGuest represents an unauthenticated guest user
	UtGuest UserType = iota

	// UtUser represents a regular user
	UtUser

	// UtAdmin represents an administrator user
	UtAdmin
)

// String stringifies the value
func (tp UserType) String() string {
	switch tp {
	case UtGuest:
		return "guest"
	case UtUser:
		return "user"
	case UtAdmin:
		return "admin"
	}
	return ""
}

// FromString parses the client type from string
func (tp *UserType) FromString(str string) error {
	switch strings.ToLower(str) {
	case "guest":
		*tp = UtGuest
	case "user":
		*tp = UtUser
	case "admin":
		*tp = UtAdmin
	default:
		return errors.Errorf(
			"invalid string representation of UserType type: %s",
			str,
		)
	}
	return nil
}
