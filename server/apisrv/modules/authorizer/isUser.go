package authorizer

import (
	"github.com/qbeon/webwire-example-postboard/server/apisrv/api"
	"github.com/qbeon/webwire-example-postboard/server/apisrv/sessinfo"
)

// IsUser creates a new condition
func IsUser(format string, v ...interface{}) isUser {
	return isUser{
		err: Errorf(
			"client is not a user "+format,
			v...,
		),
	}
}

// isUser represents an implementation of the Condition interface
type isUser struct {
	err error
}

// Validate implements the Condition interface.
// It returns an error if the user provided by the given session information
// is not a regular user
func (val isUser) Validate(sessionInfo *sessinfo.SessionInfo) error {
	if sessionInfo.UserType == api.UtUser {
		return nil
	}
	return val.err
}
