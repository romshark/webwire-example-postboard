package validator

import (
	"regexp"

	"github.com/pkg/errors"
)

// validator represents an implementation of the Validator interface
type validator struct {
	UsernameRule                 *regexp.Regexp
	NameRule                     *regexp.Regexp
	MinPostLength                uint32
	MaxPostLength                uint32
	MinReactionDescriptionLength uint32
	MaxReactionDescriptionLength uint32
	MaxPostsQueryLimit           uint32
	MinPasswordLength            uint32
	MaxPasswordLength            uint32
}

// NewValidator creates a new validator module instance
func NewValidator(
	usernameRule *regexp.Regexp,
	nameRule *regexp.Regexp,
	minPostLength,
	maxPostLength,
	minReactionDescriptionLength,
	maxReactionDescriptionLength,
	maxPostsQueryLimit,
	minPasswordLength,
	maxPasswordLength uint32,
) Validator {
	if nameRule == nil {
		panic(errors.New("missing validator name rule"))
	}

	return &validator{
		UsernameRule:                 usernameRule,
		NameRule:                     nameRule,
		MinPostLength:                minPostLength,
		MaxPostLength:                maxPostLength,
		MinReactionDescriptionLength: minReactionDescriptionLength,
		MaxReactionDescriptionLength: maxReactionDescriptionLength,
		MaxPostsQueryLimit:           maxPostsQueryLimit,
		MinPasswordLength:            minPasswordLength,
		MaxPasswordLength:            maxPasswordLength,
	}
}
