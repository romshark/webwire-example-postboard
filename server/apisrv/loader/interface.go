package loader

import engiface "github.com/qbeon/webwire-messenger/server/apisrv/engine"

// Loader defines the interface of a loader module
type Loader interface {
	// LoadMessage loads a message identified by 'identifier'
	// blocking until the message is loaded by the next batch
	LoadMessage(identifier engiface.Identifier) (*engiface.Message, error)
}
