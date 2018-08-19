package authorizer

import engiface "github.com/qbeon/webwire-messenger/server/apisrv/engine"

// IsAuthenticated represents an implementation of the Condition interface
type IsAuthenticated struct{}

// Validate implements the Condition interface.
// It returns an error if the user provided by the given session information
// is not authenticated (a guest)
func (val IsAuthenticated) Validate(sessionInfo *engiface.SessionInfo) error {
	if sessionInfo.UserType == engiface.UtGuest {
		return engiface.NewError(engiface.ErrUnauthorized)
	}
	return nil
}
