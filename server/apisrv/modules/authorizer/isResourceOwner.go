package authorizer

import (
	"github.com/qbeon/webwire-messenger/server/apisrv/api"
	"github.com/qbeon/webwire-messenger/server/apisrv/sessinfo"
)

// IsResourceOwner creates a new condition
func IsResourceOwner(
	resourceOwner api.Identifier,
	format string,
	v ...interface{},
) isResourceOwner {
	return isResourceOwner{
		resourceOwner: resourceOwner,
		err: Errorf(
			"client is not the resource owner, "+format,
			v...,
		),
	}
}

// isResourceOwner represents an implementation of the Condition interface
type isResourceOwner struct {
	resourceOwner api.Identifier
	err           error
}

// Validate implements the Condition interface.
// It returns an error if the user provided by the given session information
// is not a the specified owner of the resource
func (val isResourceOwner) Validate(sessionInfo *sessinfo.SessionInfo) error {
	if sessionInfo.UserIdentifier != val.resourceOwner {
		return val.err
	}
	return nil
}
