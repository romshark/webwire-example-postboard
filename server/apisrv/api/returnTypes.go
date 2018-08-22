package api

// CreateUserReturn represents the return type of the CreateUser method
type CreateUserReturn struct {
	UserIdent Identifier `json:"id"`
}

// CreatePostReturn represents the return type of api.CreatePost
type CreatePostReturn struct {
	PostIdent Identifier `json:"id"`
}

// CreatePostReactionReturn represents the return type
// of api.CreatePostReaction
type CreatePostReactionReturn struct {
	PostReactionIdent Identifier `json:"id"`
}
