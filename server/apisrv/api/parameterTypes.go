package api

// LoginParams defines the parameters
// required by the login mutation
type LoginParams struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// CreateUserParams defines the parameters
// requires by the CreateUser mutation
type CreateUserParams struct {
	FirstName string   `json:"firstName"`
	LastName  string   `json:"lastName"`
	Username  string   `json:"username"`
	Password  string   `json:"password"`
	Type      UserType `json:"type"`
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
