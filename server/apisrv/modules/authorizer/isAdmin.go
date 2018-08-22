package authorizer

import (
	"github.com/qbeon/webwire-example-postboard/server/apisrv/api"
	"github.com/qbeon/webwire-example-postboard/server/apisrv/sessinfo"
)

// IsAdmin creates a new condition
func IsAdmin(format string, v ...interface{}) isAdmin {
	return isAdmin{
		err: Errorf(
			"client is not an administrator "+format,
			v...,
		),
	}
}

// isAdmin represents an implementation of the Condition interface
type isAdmin struct {
	err error
}

// Validate implements the Condition interface.
// It returns an error if the user provided by the given session information
// is not an administrator
func (val isAdmin) Validate(sessionInfo *sessinfo.SessionInfo) error {
	if sessionInfo.UserType == api.UtAdmin {
		return nil
	}
	return val.err
}
