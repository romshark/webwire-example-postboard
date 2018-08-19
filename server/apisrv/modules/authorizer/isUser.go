package authorizer

import (
	"github.com/qbeon/webwire-messenger/server/apisrv/api"
	engiface "github.com/qbeon/webwire-messenger/server/apisrv/modules/engine"
	"github.com/qbeon/webwire-messenger/server/apisrv/sessinfo"
)

// IsUser represents an implementation of the Condition interface
type IsUser struct{}

// Validate implements the Condition interface.
// It returns an error if the user provided by the given session information
// is not a regular user
func (val IsUser) Validate(sessionInfo *sessinfo.SessionInfo) error {
	if sessionInfo.UserType == api.UtUser {
		return nil
	}
	return engiface.NewError(engiface.ErrUnauthorized)
}
