package resolver

import (
	wwr "github.com/qbeon/webwire-go"
	"github.com/qbeon/webwire-messenger/server/apisrv/api"
	"github.com/qbeon/webwire-messenger/server/apisrv/sessinfo"
)

// Resolver defines the interface of the request resolver module
type Resolver interface {
	// Resolve resolves an incoming request message
	Resolve(
		connection wwr.Connection,
		session *sessinfo.SessionInfo,
		message wwr.Message,
	) (wwr.Payload, error)

	// Login signs the client into either an administrator or a user account
	// creating a session
	Login(
		session *sessinfo.SessionInfo,
		connection wwr.Connection,
		params *api.LoginParams,
	) (wwr.Payload, error)

	// Logout signs the client out making him a guest
	Logout(
		session *sessinfo.SessionInfo,
		connection wwr.Connection,
	) (wwr.Payload, error)

	// GetMessages looks for n messages (n = limit) after
	// the message identified by the given identifier.
	// After is set to the latest message if not explicitly specified
	GetMessages(
		session *sessinfo.SessionInfo,
		params *api.GetMessagesParams,
	) (wwr.Payload, error)

	// PostMessage posts a new message.
	// Requires
	PostMessage(
		session *sessinfo.SessionInfo,
		params *api.PostMessageParams,
	) (wwr.Payload, error)

	// EditMessage permanently changes the contents of a message.
	// Requires message owner permissions
	EditMessage(
		session *sessinfo.SessionInfo,
		params *api.EditMessageParams,
	) (wwr.Payload, error)

	// RemoveMessage permanently removes a messages.
	// Requires either message owner or administrator permissions
	RemoveMessage(
		session *sessinfo.SessionInfo,
		params *api.RemoveMessageParams,
	) (wwr.Payload, error)

	// PostMessageReaction adds a reaction to the message
	PostMessageReaction(
		session *sessinfo.SessionInfo,
		params *api.PostMessageReactionParams,
	) (wwr.Payload, error)
}
