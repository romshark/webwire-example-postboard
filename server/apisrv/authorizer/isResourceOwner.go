package authorizer

import engiface "github.com/qbeon/webwire-messenger/server/apisrv/engine"

// IsResourceOwner represents an implementation of the Condition interface
type IsResourceOwner struct {
	ResourceOwner engiface.Identifier
}

// Validate implements the Condition interface.
// It returns an error if the user provided by the given session information
// is not a the specified owner of the resource
func (val IsResourceOwner) Validate(sessionInfo *engiface.SessionInfo) error {
	if sessionInfo.UserIdentifier != val.ResourceOwner {
		return engiface.NewError(engiface.ErrUnauthorized)
	}
	return nil
}
