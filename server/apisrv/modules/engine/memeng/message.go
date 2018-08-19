package memeng

import "github.com/qbeon/webwire-messenger/server/apisrv/api"

// Message defines the message in-memory storage representations
type Message struct {
	// Message stores the actual message data
	Message api.Message

	// Archived indicates whether this message is archived or not
	Archived bool
}
