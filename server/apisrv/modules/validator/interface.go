package validator

// Validator defines the interface of the validator module
type Validator interface {
	// Name returns an error if the given firstName / lastName is invalid,
	// otherwise returns nil
	Name(name string) error

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
