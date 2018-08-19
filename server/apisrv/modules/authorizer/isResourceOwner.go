package authorizer

import (
	"github.com/qbeon/webwire-messenger/server/apisrv/api"
	engiface "github.com/qbeon/webwire-messenger/server/apisrv/modules/engine"
	"github.com/qbeon/webwire-messenger/server/apisrv/sessinfo"
)

// IsResourceOwner represents an implementation of the Condition interface
type IsResourceOwner struct {
	ResourceOwner api.Identifier
}

// Validate implements the Condition interface.
// It returns an error if the user provided by the given session information
// is not a the specified owner of the resource
func (val IsResourceOwner) Validate(sessionInfo *sessinfo.SessionInfo) error {
	if sessionInfo.UserIdentifier != val.ResourceOwner {
		return engiface.NewError(engiface.ErrUnauthorized)
	}
	return nil
}
