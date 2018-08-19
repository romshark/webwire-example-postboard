package engine

import (
	"time"

	wwr "github.com/qbeon/webwire-go"
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
	AuthenticateUser(username string) (*User, error)

	// GetMessages retrieves the messages identified by the given 'identifiers'
	// in the same order
	GetMessages(
		identifiers []Identifier,
	) ([]*Message, error)

	// GetMessagesAfter retrieves n messages (n = 'limit') after the message
	// identified by 'after'.
	// If 'after' is nil then the latest message will be applied
	GetMessagesAfter(
		after *Identifier,
		limit uint32,
	) ([]*Message, error)

	// PostMessage appends a new message to the history
	PostMessage(newMessage *Message) error

	// EditMessage edits the contents of the message identified by 'ident'
	EditMessage(
		ident Identifier,
		editTime time.Time,
		newContents string,
	) error

	// RemoveMessage permanently removes the message identified by 'ident'
	// from the message history
	RemoveMessage(ident Identifier) error

	// PostMessageReaction posts a reaction to the message identified by 'ident'
	PostMessageReaction(
		messageIdent Identifier,
		reaction *MessageReaction,
	) error
}
