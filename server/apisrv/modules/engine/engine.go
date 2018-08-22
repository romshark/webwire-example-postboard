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

	// FindSession searches for a session identified by `key`
	FindSession(key string) (wwr.SessionLookupResult, error)

	// CloseSession closes the session identified by `key`
	CloseSession(key string) error

	// AuthenticateUser returns the account of the user identified by `username`
	AuthenticateUser(username string) (
		profile *api.User,
		password string,
		err error,
	)

	// CreateUser permanently saves the new account.
	// Returns an error if the identifier or username are already reserved
	CreateUser(newAccount *api.User, passwordHash string) error

	// GetUser retrieves the user profile by identifier
	GetUser(ident api.Identifier) (*api.User, error)

	// GetPosts retrieves the posts identified by `idents`
	// returning them in the same order as `idents`
	GetPosts(idents []api.Identifier) ([]*api.Post, error)

	// GetPostReactions retrieves the post reactions identified by `idents`
	// returning them in the same order as `idents`
	GetPostReactions(idents []api.Identifier) ([]*api.PostReaction, error)

	// GetReactionsOfPost retrieves all reactions of a certain post
	GetReactionsOfPost(postIdent api.Identifier) ([]*api.PostReaction, error)

	// GetPostsAfter retrieves n posts (n = `limit`) after the post
	// identified by `after`.
	// If `after` is nil then the latest post will be applied
	GetPostsAfter(
		after *api.Identifier,
		limit uint32,
	) ([]*api.Post, error)

	// CreatePost appends a new post to the history
	CreatePost(newPost *api.Post) error

	// EditPost edits the contents of the post identified by `ident`
	// setting the contents to `newContents` updating last-edit time to
	// editTime
	EditPost(
		ident api.Identifier,
		editTime time.Time,
		newContents string,
	) error

	// RemovePost permanently removes the post identified by `ident`
	// from the history
	RemovePost(ident api.Identifier) error

	// RemovePostReaction permanently removes the post reaction
	// identified by `ident` from the history
	RemovePostReaction(ident api.Identifier) error

	// CreatePostReaction posts a reaction to the post identified by `ident`
	CreatePostReaction(
		postIdent api.Identifier,
		reaction *api.PostReaction,
	) error
}
