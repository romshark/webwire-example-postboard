package loader

import (
	"github.com/qbeon/webwire-messenger/server/apisrv/api"
)

// Loader defines the interface of a loader module
type Loader interface {
	// LoadMessage loads a message identified by 'identifier'
	// blocking until the message is loaded by the next batch
	LoadMessage(identifier api.Identifier) (*api.Message, error)
}
