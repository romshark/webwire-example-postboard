package resolver

import (
	wwr "github.com/qbeon/webwire-go"
	engiface "github.com/qbeon/webwire-messenger/server/apisrv/engine"
)

// Resolver defines the interface of the request resolver module
type Resolver interface {
	// Resolve resolves an incoming request message
	Resolve(
		connection wwr.Connection,
		session *engiface.SessionInfo,
		message wwr.Message,
	) (wwr.Payload, error)

	// Login signs the client into either an administrator or a user account
	// creating a session
	Login(
		session *engiface.SessionInfo,
		connection wwr.Connection,
		params *LoginParams,
	) (wwr.Payload, error)

	// Logout signs the client out making him a guest
	Logout(
		session *engiface.SessionInfo,
		connection wwr.Connection,
	) (wwr.Payload, error)

	// GetMessages looks for n messages (n = limit) after
	// the message identified by the given identifier.
	// After is set to the latest message if not explicitly specified
	GetMessages(
		session *engiface.SessionInfo,
		params *GetMessagesParams,
	) (wwr.Payload, error)

	// PostMessage posts a new message.
	// Requires
	PostMessage(
		session *engiface.SessionInfo,
		params *PostMessageParams,
	) (wwr.Payload, error)

	// EditMessage permanently changes the contents of a message.
	// Requires message owner permissions
	EditMessage(
		session *engiface.SessionInfo,
		params *EditMessageParams,
	) (wwr.Payload, error)

	// RemoveMessage permanently removes a messages.
	// Requires either message owner or administrator permissions
	RemoveMessage(
		session *engiface.SessionInfo,
		params *RemoveMessageParams,
	) (wwr.Payload, error)

	// PostMessageReaction adds a reaction to the message
	PostMessageReaction(
		session *engiface.SessionInfo,
		params *PostMessageReactionParams,
	) (wwr.Payload, error)
}

// LoginParams defines the parameters
// required by the login mutation
type LoginParams struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// GetMessagesParams defines the parameters
// required by the message reaction creation mutation
type GetMessagesParams struct {
	After *engiface.Identifier `json:"after"`
	Limit uint32               `json:"limit"`
}

// PostMessageParams defines the parameters
// required by the message creation mutation
type PostMessageParams struct {
	Contents string `json:"contents"`
}

// EditMessageParams defines the parameters
// required by the message editing mutation
type EditMessageParams struct {
	MessageIdent engiface.Identifier `json:"messageId"`
	NewContents  string              `json:"newContents"`
}

// RemoveMessageParams defines the parameters
// required by the message deletion mutation
type RemoveMessageParams struct {
	MessageIdent engiface.Identifier `json:"messageId"`
}

// PostMessageReactionParams defines the parameters
// required by the message reaction creation mutation
type PostMessageReactionParams struct {
	MessageIdent engiface.Identifier          `json:"messageId"`
	AuthorIdent  engiface.Identifier          `json:"authorId"`
	Type         engiface.MessageReactionType `json:"type"`
	Description  string                       `json:"description"`
}
