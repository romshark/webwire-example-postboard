package memeng

import (
	"sync"

	wwr "github.com/qbeon/webwire-go"
	engiface "github.com/qbeon/webwire-messenger/server/apisrv/engine"
)

// engine represents an in-memory implementation of the API Engine interface
type engine struct {
	// lock protects the engine internals from concurrent access
	// and prevents data races
	lock *sync.Mutex

	// users stores all user accounts indexed by identifier
	users map[engiface.Identifier]*engiface.User

	// usersByUsername stores the user.username index
	usersByUsername map[string]*engiface.User

	// sessions stores all currently open sessions indexed by key
	sessions map[string]*wwr.Session

	// messages stores references to all messages in chronological order
	messages []*engiface.Message

	// messagesByIdent stores the message.identifier index
	messagesByIdent map[engiface.Identifier]int
}

// New initializes a new in-memory engine implementation
func New(
	preallocSessions int,
	preallocMessages int,
) engiface.Engine {
	return &engine{
		lock: &sync.Mutex{},

		// Initialize the users store
		users: make(map[engiface.Identifier]*engiface.User, 0),

		// Initialize the users.username index
		usersByUsername: make(map[string]*engiface.User, 0),

		// Preallocate the sessions store
		sessions: make(map[string]*wwr.Session, preallocSessions),

		// Preallocate the message store
		messages: make([]*engiface.Message, 0, preallocMessages),

		// Preallocate the message.identifier index
		messagesByIdent: make(
			map[engiface.Identifier]int,
			preallocMessages,
		),
	}
}

// findMessageByIdent finds and returns the message identified by 'ident'.
// Returns a not-found error if the message doesn't exist or is archived
func (eng *engine) findMessageByIdent(
	ident engiface.Identifier,
) (*engiface.Message, error) {
	messageIndex, exists := eng.messagesByIdent[ident]
	if !exists {
		// Return not-found error if the message doesn't exist
		return nil, engiface.NewError(engiface.ErrMessageNotFound)
	}
	message := eng.messages[messageIndex]
	if message.Archived {
		// Return not-found error if the message is archived
		return nil, engiface.NewError(engiface.ErrMessageNotFound)
	}
	return message, nil
}
