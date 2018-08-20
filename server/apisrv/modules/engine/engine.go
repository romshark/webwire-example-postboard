package engine

import (
	"time"

	wwr "github.com/qbeon/webwire-go"
	"github.com/qbeon/webwire-messenger/server/apisrv/api"
)

// Engine defines the service engine interface
type Engine interface {
	// SaveSession permanently saves the given new session
	SaveSession(newSession *wwr.Session) error

	// FindSession searches for a session identified by 'key'
	FindSession(key string) (wwr.SessionLookupResult, error)

	// CloseSession closes the session identified by 'key'
	CloseSession(key string) error

	// AuthenticateUser returns the account of the user identified by 'username'
	AuthenticateUser(username string) (
		profile *api.User,
		password string,
		err error,
	)

	// CreateUser permanently saves the new account.
	// Returns an error if the identifier or username are already reserved
	CreateUser(newAccount *api.User, passwordHash string) error

	// GetMessages retrieves the messages identified by the given 'identifiers'
	// in the same order
	GetMessages(
		identifiers []api.Identifier,
	) ([]*api.Message, error)

	// GetMessagesAfter retrieves n messages (n = 'limit') after the message
	// identified by 'after'.
	// If 'after' is nil then the latest message will be applied
	GetMessagesAfter(
		after *api.Identifier,
		limit uint32,
	) ([]*api.Message, error)

	// PostMessage appends a new message to the history
	PostMessage(newMessage *api.Message) error

	// EditMessage edits the contents of the message identified by 'ident'
	EditMessage(
		ident api.Identifier,
		editTime time.Time,
		newContents string,
	) error

	// RemoveMessage permanently removes the message identified by 'ident'
	// from the message history
	RemoveMessage(ident api.Identifier) error

	// PostMessageReaction posts a reaction to the message identified by 'ident'
	PostMessageReaction(
		messageIdent api.Identifier,
		reaction *api.MessageReaction,
	) error
}
