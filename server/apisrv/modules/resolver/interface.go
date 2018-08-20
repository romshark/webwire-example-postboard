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
	) (interface{}, error)

	// Logout signs the client out making him a guest
	Logout(
		session *sessinfo.SessionInfo,
		connection wwr.Connection,
	) (interface{}, error)

	// CreateUser creates a new user account
	CreateUser(
		session *sessinfo.SessionInfo,
		params *api.CreateUserParams,
	) (interface{}, error)

	// GetMessages looks for n messages (n = limit) after
	// the message identified by the given identifier.
	// After is set to the latest message if not explicitly specified
	GetMessages(
		session *sessinfo.SessionInfo,
		params *api.GetMessagesParams,
	) (interface{}, error)

	// PostMessage posts a new message.
	// Requires
	PostMessage(
		session *sessinfo.SessionInfo,
		params *api.PostMessageParams,
	) (interface{}, error)

	// EditMessage permanently changes the contents of a message.
	// Requires message owner permissions
	EditMessage(
		session *sessinfo.SessionInfo,
		params *api.EditMessageParams,
	) (interface{}, error)

	// RemoveMessage permanently removes a messages.
	// Requires either message owner or administrator permissions
	RemoveMessage(
		session *sessinfo.SessionInfo,
		params *api.RemoveMessageParams,
	) (interface{}, error)

	// PostMessageReaction adds a reaction to the message
	PostMessageReaction(
		session *sessinfo.SessionInfo,
		params *api.PostMessageReactionParams,
	) (interface{}, error)
}
