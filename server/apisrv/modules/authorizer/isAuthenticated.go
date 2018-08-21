package authorizer

import (
	"github.com/qbeon/webwire-messenger/server/apisrv/api"
	"github.com/qbeon/webwire-messenger/server/apisrv/sessinfo"
)

// IsAuthenticated creates a new condition
func IsAuthenticated(format string, v ...interface{}) isAuthenticated {
	return isAuthenticated{
		err: Errorf(
			"client is not authenticated "+format,
			v...,
		),
	}
}

// isAuthenticated represents an implementation of the Condition interface
type isAuthenticated struct {
	err error
}

// Validate implements the Condition interface.
// It returns an error if the user provided by the given session information
// is not authenticated (a guest)
func (val isAuthenticated) Validate(sessionInfo *sessinfo.SessionInfo) error {
	if sessionInfo.UserType == api.UtGuest {
		return val.err
	}
	return nil
}
