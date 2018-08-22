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

	// posts stores references to all posts in chronological order
	posts []*Post

	// postsByIdent stores the post.identifier index
	postsByIdent map[api.Identifier]int

	// postReactionsByIdent stores the postReaction.identifier index
	postReactionsByIdent map[api.Identifier]*PostReaction
}

// New initializes a new in-memory engine implementation
func New(
	passwordHasher passhash.PasswordHasher,
	defaultRootPassword string,
	preallocSessions int,
	preallocPosts int,
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

		// Preallocate the posts store
		posts: make([]*Post, 0, preallocPosts),

		// Preallocate the post.identifier index
		postsByIdent: make(map[api.Identifier]int, preallocPosts),

		// Initialize the postReaction.identifier index
		postReactionsByIdent: make(map[api.Identifier]*PostReaction),
	}

	// Encrypt the default password in order for it to work
	hashedDefaultRootPassword, err := inMemoryEngine.passwordHasher.Hash(
		defaultRootPassword,
	)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't hash the default root password")
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
