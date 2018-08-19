package authorizer

import (
	"github.com/qbeon/webwire-messenger/server/apisrv/api"
	engiface "github.com/qbeon/webwire-messenger/server/apisrv/modules/engine"
	"github.com/qbeon/webwire-messenger/server/apisrv/sessinfo"
)

// IsAuthenticated represents an implementation of the Condition interface
type IsAuthenticated struct{}

// Validate implements the Condition interface.
// It returns an error if the user provided by the given session information
// is not authenticated (a guest)
func (val IsAuthenticated) Validate(sessionInfo *sessinfo.SessionInfo) error {
	if sessionInfo.UserType == api.UtGuest {
		return engiface.NewError(engiface.ErrUnauthorized)
	}
	return nil
}
