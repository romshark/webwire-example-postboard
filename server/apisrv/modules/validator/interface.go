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

	// MessageContents returns an error if the given message contents
	// are invalid, otherwise returns nil
	MessageContents(contents string) error

	// ReactionDescription returns an error if the reaction description
	// is invalid, otherwise returns nil
	ReactionDescription(description string) error

	// MessagesQueryLimit returns an error if the given messages query limit
	// is invalid, otherwise returns nil
	MessagesQueryLimit(limit uint32) error
}
