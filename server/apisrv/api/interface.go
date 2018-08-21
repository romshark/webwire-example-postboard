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

	// GetMessage finds a specific message by the given identifier
	GetMessage(
		ctx context.Context,
		params GetMessageParams,
	) (*Message, error)

	// GetMessages looks for n messages (n = 'limit') after
	// the message identified by the given identifier.
	// After is set to the latest message if not explicitly specified
	GetMessages(
		ctx context.Context,
		params GetMessagesParams,
	) ([]*Message, error)

	// PostMessage posts a new message.
	// Requires
	PostMessage(
		ctx context.Context,
		params PostMessageParams,
	) (Identifier, error)

	// EditMessage permanently changes the contents of a message.
	// Requires message owner permissions
	EditMessage(
		ctx context.Context,
		params EditMessageParams,
	) error

	// RemoveMessage permanently removes a messages.
	// Requires either message owner or administrator permissions
	RemoveMessage(
		ctx context.Context,
		params RemoveMessageParams,
	) error

	// PostMessageReaction adds a reaction to the message
	PostMessageReaction(
		ctx context.Context,
		params PostMessageReactionParams,
	) (Identifier, error)
}
