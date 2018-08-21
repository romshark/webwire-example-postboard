package memeng

import (
	"sync"
	"time"

	"github.com/pkg/errors"
	wwr "github.com/qbeon/webwire-go"
	"github.com/qbeon/webwire-messenger/server/apisrv/api"
	engiface "github.com/qbeon/webwire-messenger/server/apisrv/modules/engine"
	"github.com/qbeon/webwire-messenger/server/apisrv/modules/passhash"
)

// engine represents an in-memory implementation of the API Engine interface
type engine struct {
	// lock protects the engine internals from concurrent access
	// and prevents data races
	lock *sync.Mutex

	// passwordHasher is used during the initialization of the engine
	// to encrypt the default root password
	passwordHasher passhash.PasswordHasher

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
	passwordHasher passhash.PasswordHasher,
	defaultRootPassword string,
	preallocSessions int,
	preallocMessages int,
) (engiface.Engine, error) {
	inMemoryEngine := &engine{
		lock:           &sync.Mutex{},
		passwordHasher: passwordHasher,

		// Initialize the users store
		users: make(map[api.Identifier]*UserAccount),

		// Initialize the users.username index
		usersByUsername: make(map[string]*UserAccount),

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

	// Encrypt the default password in order for it to work
	hashedDefaultRootPassword, err := inMemoryEngine.passwordHasher.Hash(
		defaultRootPassword,
	)
	if err != nil {
		return nil, errors.Wrap(
			err,
			"couldn't hash the default root password",
		)
	}

	// Create root administrator user
	if err := inMemoryEngine.createUser(&UserAccount{
		Profile: api.User{
			Identifier:   api.NewIdentifier(),
			Username:     "root",
			Registration: time.Now().UTC(),
			Type:         api.UtAdmin,
		},
		Password: hashedDefaultRootPassword,
	}); err != nil {
		return nil, errors.Wrap(err, "couldn't create root admin user")
	}

	return inMemoryEngine, nil
}
