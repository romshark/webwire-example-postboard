package authorizer

import "github.com/qbeon/webwire-messenger/server/apisrv/sessinfo"

// Condition represents an authorization condition interface
type Condition interface {
	Validate(sessionInfo *sessinfo.SessionInfo) error
}

// Authorizer defines the interface of the authorizer module
type Authorizer interface {
	// MeetsAll returns an error if either of the given conditions don't apply.
	// If all of the given conditions apply collectively then nil is returned
	MeetsAll(
		sessionInfo *sessinfo.SessionInfo,
		conditions ...Condition,
	) error

	// MeetsEitherOf returns an error if all of the given conditions
	// don't apply.
	// If at least one of the given conditions applies then nil is returned
	MeetsEitherOf(
		sessionInfo *sessinfo.SessionInfo,
		conditions ...Condition,
	) error
}
