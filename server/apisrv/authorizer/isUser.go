package authorizer

import engiface "github.com/qbeon/webwire-messenger/server/apisrv/engine"

// IsUser represents an implementation of the Condition interface
type IsUser struct{}

// Validate implements the Condition interface.
// It returns an error if the user provided by the given session information
// is not a regular user
func (val IsUser) Validate(sessionInfo *engiface.SessionInfo) error {
	if sessionInfo.UserType == engiface.UtUser {
		return nil
	}
	return engiface.NewError(engiface.ErrUnauthorized)
}
