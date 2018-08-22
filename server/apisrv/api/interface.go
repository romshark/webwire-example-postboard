package api

import "context"

// Api defines the interface of the messenger API
type Api interface {
	// Login signs the client into either an administrator or a user account
	// creating a session
	Login(
		ctx context.Context,
		params LoginParams,
	) error

	// Logout signs the client out making him a guest
	Logout(
		ctx context.Context,
	) error

	// CreateUser creates a new user account
	CreateUser(
		ctx context.Context,
		params CreateUserParams,
	) (Identifier, error)

	// GetUser finds and returns the user associated with the given identifier
	GetUser(
		ctx context.Context,
		params GetUserParams,
	) (*User, error)

	// GetPost finds a specific post by the given identifier
	GetPost(
		ctx context.Context,
		params GetPostParams,
	) (*Post, error)

	// GetPostReaction finds a specific post reaction by `params.ReactionIdent`
	GetPostReaction(
		ctx context.Context,
		params GetPostReactionParams,
	) (*PostReaction, error)

	// GetReactionsOfPost returns the reactions to a specific post
	GetReactionsOfPost(
		ctx context.Context,
		params GetReactionsOfPostParams,
	) ([]*PostReaction, error)

	// GetPosts looks for n posts (n = 'limit') after
	// the post identified by the given identifier.
	// After is set to the latest post if not explicitly specified
	GetPosts(
		ctx context.Context,
		params GetPostsParams,
	) ([]*Post, error)

	// CreatePost creates a new post.
	// Requires
	CreatePost(
		ctx context.Context,
		params CreatePostParams,
	) (Identifier, error)

	// EditPost permanently changes the contents of a post.
	// Requires post author permissions
	EditPost(
		ctx context.Context,
		params EditPostParams,
	) error

	// RemovePost removes a post.
	// Requires either post author or administrator permissions
	RemovePost(
		ctx context.Context,
		params RemovePostParams,
	) error

	// CreatePostReaction creates a new reaction on the post
	CreatePostReaction(
		ctx context.Context,
		params CreatePostReactionParams,
	) (Identifier, error)

	// RemovePostReaction removes a post reaction.
	// Requires either reaction author or administrator permissions
	RemovePostReaction(
		ctx context.Context,
		params RemovePostReactionParams,
	) error
}
