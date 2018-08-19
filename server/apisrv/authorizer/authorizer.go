package authorizer

import engiface "github.com/qbeon/webwire-messenger/server/apisrv/engine"

// authorizer represents an implementation of the Authorizer module interface
type authorizer struct{}

// New creates a new authorizer module instance
func New() Authorizer {
	return &authorizer{}
}

// MeetsAll implements the Authorizer interface
func (auth *authorizer) MeetsAll(
	sessionInfo *engiface.SessionInfo,
	conditions ...Condition,
) (err error) {
	for _, condition := range conditions {
		if err = condition.Validate(sessionInfo); err != nil {
			return err
		}
	}
	return nil
}

// MeetsEitherOf implements the Authorizer interface
func (auth *authorizer) MeetsEitherOf(
	sessionInfo *engiface.SessionInfo,
	conditions ...Condition,
) (err error) {
	var lastError error
	for _, condition := range conditions {
		if err = condition.Validate(sessionInfo); err == nil {
			return nil
		}
		lastError = err
	}
	return lastError
}
