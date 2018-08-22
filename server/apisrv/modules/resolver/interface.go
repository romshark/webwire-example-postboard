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

	// GetPosts looks for n posts (n = limit) after
	// the post identified by `params.Ident`.
	// After is set to the latest post if not explicitly specified
	GetPosts(
		session *sessinfo.SessionInfo,
		params *api.GetPostsParams,
	) (interface{}, error)

	// GetReactionsOfPost find all reactions of a certain post
	// identified by `params.PostIdent`
	GetReactionsOfPost(
		session *sessinfo.SessionInfo,
		params *api.GetReactionsOfPostParams,
	) (interface{}, error)

	// CreatePost creates a new post.
	// Requires
	CreatePost(
		session *sessinfo.SessionInfo,
		params *api.CreatePostParams,
	) (interface{}, error)

	// EditPost permanently changes the contents of a post.
	// Requires post author permissions
	EditPost(
		session *sessinfo.SessionInfo,
		params *api.EditPostParams,
	) (interface{}, error)

	// RemovePost permanently removes a post.
	// Requires either post author or administrator permissions
	RemovePost(
		session *sessinfo.SessionInfo,
		params *api.RemovePostParams,
	) (interface{}, error)

	// CreatePostReaction adds a reaction to the post
	CreatePostReaction(
		session *sessinfo.SessionInfo,
		params *api.CreatePostReactionParams,
	) (interface{}, error)

	// RemovePostReaction permanently removes a post reaction.
	// Requires either reaction author or administrator permissions
	RemovePostReaction(
		session *sessinfo.SessionInfo,
		params *api.RemovePostReactionParams,
	) (interface{}, error)
}
