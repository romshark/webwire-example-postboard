package authorizer

import (
	"github.com/qbeon/webwire-example-postboard/server/apisrv/api"
	"github.com/qbeon/webwire-example-postboard/server/apisrv/sessinfo"
)

// IsNotResourceOwner creates a new condition
func IsNotResourceOwner(
	resourceOwner api.Identifier,
	format string,
	v ...interface{},
) isNotResourceOwner {
	return isNotResourceOwner{
		resourceOwner: resourceOwner,
		err: Errorf(
			"client is resource owner, "+format,
			v...,
		),
	}
}

// isNotResourceOwner represents an implementation of the Condition interface
type isNotResourceOwner struct {
	resourceOwner api.Identifier
	err           error
}

// Validate implements the Condition interface.
// It returns an error if the user provided by the given session information
// is the specified owner of the resource
func (val isNotResourceOwner) Validate(sessionInfo *sessinfo.SessionInfo) error {
	if sessionInfo.UserIdentifier == val.resourceOwner {
		return val.err
	}
	return nil
}
