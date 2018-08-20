package api

// CreateUserReturn represents the return type of the CreateUser method
type CreateUserReturn struct {
	UserIdent Identifier `json:"id"`
}

// GetMessagesReturn represents the return type of the GetMessages method
type GetMessagesReturn struct {
	Messages []*Message `json:"messages"`
}

// PostMessageReturn represents the return type of the PostMessage method
type PostMessageReturn struct {
	MessageIdent Identifier `json:"id"`
}

// PostMessageReactionReturn represents the return type
// of the PostMessageReaction method
type PostMessageReactionReturn struct {
	MessageReactionIdent Identifier `json:"id"`
}
