package api

// LoginParams defines the parameters
// required by the login mutation
type LoginParams struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// CreateUserParams defines the parameters
// required by the CreateUser mutation
type CreateUserParams struct {
	FirstName string   `json:"firstName"`
	LastName  string   `json:"lastName"`
	Username  string   `json:"username"`
	Password  string   `json:"password"`
	Type      UserType `json:"type"`
}

// GetUserParams defines the parameters required by api.GetUser
type GetUserParams struct {
	Ident Identifier `json:"id"`
}

// GetPostParams defines the parameters required by api.GetPost
type GetPostParams struct {
	Ident Identifier `json:"id"`
}

// GetPostsParams defines the parameters required by api.GetPosts
type GetPostsParams struct {
	After *Identifier `json:"after"`
	Limit uint32      `json:"limit"`
}

// CreatePostParams defines the parameters required by api.CreatePost
type CreatePostParams struct {
	Contents string `json:"contents"`
}

// EditPostParams defines the parameters required by api.EditPost
type EditPostParams struct {
	PostIdent   Identifier `json:"postId"`
	NewContents string     `json:"newContents"`
}

// RemovePostParams defines the parameters required by api.RemovePost
type RemovePostParams struct {
	PostIdent Identifier `json:"postId"`
}

// CreatePostReactionParams defines the parameters
// required by api.CreatePostReaction
type CreatePostReactionParams struct {
	PostIdent   Identifier       `json:"postId"`
	Type        PostReactionType `json:"type"`
	Description string           `json:"description"`
}
