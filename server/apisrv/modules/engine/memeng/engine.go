package memeng

import (
	"sync"

	wwr "github.com/qbeon/webwire-go"
	"github.com/qbeon/webwire-messenger/server/apisrv/api"
	engiface "github.com/qbeon/webwire-messenger/server/apisrv/modules/engine"
)

// engine represents an in-memory implementation of the API Engine interface
type engine struct {
	// lock protects the engine internals from concurrent access
	// and prevents data races
	lock *sync.Mutex

	// users stores all user accounts indexed by identifier
	users map[api.Identifier]*UserAccount

	// usersByUsername stores the user.username index
	usersByUsername map[string]*UserAccount

	// sessions stores all currently open sessions indexed by key
	sessions map[string]*wwr.Session

	// messages stores references to all messages in chronological order
	messages []*Message

	// messagesByIdent stores the message.identifier index
	messagesByIdent map[api.Identifier]int
}

// New initializes a new in-memory engine implementation
func New(
	preallocSessions int,
	preallocMessages int,
) engiface.Engine {
	return &engine{
		lock: &sync.Mutex{},

		// Initialize the users store
		users: make(map[api.Identifier]*UserAccount, 0),

		// Initialize the users.username index
		usersByUsername: make(map[string]*UserAccount, 0),

		// Preallocate the sessions store
		sessions: make(map[string]*wwr.Session, preallocSessions),

		// Preallocate the message store
		messages: make([]*Message, 0, preallocMessages),

		// Preallocate the message.identifier index
		messagesByIdent: make(
			map[api.Identifier]int,
			preallocMessages,
		),
	}
}
