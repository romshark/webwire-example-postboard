package authorizer

import engiface "github.com/qbeon/webwire-messenger/server/apisrv/engine"

// IsAdmin represents an implementation of the Condition interface
type IsAdmin struct{}

// Validate implements the Condition interface.
// It returns an error if the user provided by the given session information
// is not an administrator
func (val IsAdmin) Validate(sessionInfo *engiface.SessionInfo) error {
	if sessionInfo.UserType == engiface.UtAdmin {
		return nil
	}
	return engiface.NewError(engiface.ErrUnauthorized)
}
