package authorizer

import engiface "github.com/qbeon/webwire-messenger/server/apisrv/engine"

// IsNotResourceOwner represents an implementation of the Condition interface
type IsNotResourceOwner struct {
	ResourceOwner engiface.Identifier
}

// Validate implements the Condition interface.
// It returns an error if the user provided by the given session information
// is the specified owner of the resource
func (val IsNotResourceOwner) Validate(sessionInfo *engiface.SessionInfo) error {
	if sessionInfo.UserIdentifier != val.ResourceOwner {
		return engiface.NewError(engiface.ErrUnauthorized)
	}
	return nil
}
