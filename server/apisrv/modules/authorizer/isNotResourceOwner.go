package authorizer

import (
	"github.com/qbeon/webwire-messenger/server/apisrv/api"
	engiface "github.com/qbeon/webwire-messenger/server/apisrv/modules/engine"
	"github.com/qbeon/webwire-messenger/server/apisrv/sessinfo"
)

// IsNotResourceOwner represents an implementation of the Condition interface
type IsNotResourceOwner struct {
	ResourceOwner api.Identifier
}

// Validate implements the Condition interface.
// It returns an error if the user provided by the given session information
// is the specified owner of the resource
func (val IsNotResourceOwner) Validate(sessionInfo *sessinfo.SessionInfo) error {
	if sessionInfo.UserIdentifier != val.ResourceOwner {
		return engiface.NewError(engiface.ErrUnauthorized)
	}
	return nil
}
