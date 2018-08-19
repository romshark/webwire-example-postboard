package api

// Api defines the interface of the messenger API
type Api interface {
	// Login signs the client into either an administrator or a user account
	// creating a session
	Login(params *LoginParams) error

	// Logout signs the client out making him a guest
	Logout() error

	// GetMessages looks for n messages (n = limit) after
	// the message identified by the given identifier.
	// After is set to the latest message if not explicitly specified
	GetMessages(params *GetMessagesParams) ([]*Message, error)

	// PostMessage posts a new message.
	// Requires
	PostMessage(params *PostMessageParams) (Identifier, error)

	// EditMessage permanently changes the contents of a message.
	// Requires message owner permissions
	EditMessage(params *EditMessageParams) error

	// RemoveMessage permanently removes a messages.
	// Requires either message owner or administrator permissions
	RemoveMessage(params *RemoveMessageParams) error

	// PostMessageReaction adds a reaction to the message
	PostMessageReaction(
		params *PostMessageReactionParams,
	) (Identifier, error)
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
	After *Identifier `json:"after"`
	Limit uint32      `json:"limit"`
}

// PostMessageParams defines the parameters
// required by the message creation mutation
type PostMessageParams struct {
	Contents string `json:"contents"`
}

// EditMessageParams defines the parameters
// required by the message editing mutation
type EditMessageParams struct {
	MessageIdent Identifier `json:"messageId"`
	NewContents  string     `json:"newContents"`
}

// RemoveMessageParams defines the parameters
// required by the message deletion mutation
type RemoveMessageParams struct {
	MessageIdent Identifier `json:"messageId"`
}

// PostMessageReactionParams defines the parameters
// required by the message reaction creation mutation
type PostMessageReactionParams struct {
	MessageIdent Identifier          `json:"messageId"`
	AuthorIdent  Identifier          `json:"authorId"`
	Type         MessageReactionType `json:"type"`
	Description  string              `json:"description"`
}
