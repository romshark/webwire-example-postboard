package authorizer

import (
	"github.com/qbeon/webwire-messenger/server/apisrv/api"
	engiface "github.com/qbeon/webwire-messenger/server/apisrv/modules/engine"
	"github.com/qbeon/webwire-messenger/server/apisrv/sessinfo"
)

// IsAdmin represents an implementation of the Condition interface
type IsAdmin struct{}

// Validate implements the Condition interface.
// It returns an error if the user provided by the given session information
// is not an administrator
func (val IsAdmin) Validate(sessionInfo *sessinfo.SessionInfo) error {
	if sessionInfo.UserType == api.UtAdmin {
		return nil
	}
	return engiface.NewError(engiface.ErrUnauthorized)
}
