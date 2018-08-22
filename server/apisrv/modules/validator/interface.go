package validator

// Validator defines the interface of the validator module
type Validator interface {
	// Username returns an error if the given username is invalid,
	// otherwise returns nil
	Username(username string) error

	// Name returns an error if the given firstName / lastName is invalid,
	// otherwise returns nil
	Name(name string) error

	// Password returns an error if the given password is invalid,
	// otherwise returns nil
	Password(password string) error

	// PostContents returns an error `contents`
	// are invalid, otherwise returns nil
	PostContents(contents string) error

	// ReactionDescription returns an error if `description`
	// is invalid, otherwise returns nil
	ReactionDescription(description string) error

	// PostsQueryLimit returns an error if the given posts query limit
	// is invalid, otherwise returns nil
	PostsQueryLimit(limit uint32) error
}
