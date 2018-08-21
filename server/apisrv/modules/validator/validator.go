package validator

import (
	"regexp"

	"github.com/pkg/errors"
)

// validator represents an implementation of the Validator interface
type validator struct {
	UsernameRule                 *regexp.Regexp
	NameRule                     *regexp.Regexp
	MinMessageLength             uint32
	MaxMessageLength             uint32
	MinReactionDescriptionLength uint32
	MaxReactionDescriptionLength uint32
	MaxMessagesQueryLimit        uint32
	MinPasswordLength            uint32
	MaxPasswordLength            uint32
}

// NewValidator creates a new validator module instance
func NewValidator(
	usernameRule *regexp.Regexp,
	nameRule *regexp.Regexp,
	minMessageLength,
	maxMessageLength,
	minReactionDescriptionLength,
	maxReactionDescriptionLength,
	maxMessagesQueryLimit,
	minPasswordLength,
	maxPasswordLength uint32,
) Validator {
	if nameRule == nil {
		panic(errors.New("missing validator name rule"))
	}

	return &validator{
		UsernameRule:                 usernameRule,
		NameRule:                     nameRule,
		MinMessageLength:             minMessageLength,
		MaxMessageLength:             maxMessageLength,
		MinReactionDescriptionLength: minReactionDescriptionLength,
		MaxReactionDescriptionLength: maxReactionDescriptionLength,
		MaxMessagesQueryLimit:        maxMessagesQueryLimit,
		MinPasswordLength:            minPasswordLength,
		MaxPasswordLength:            maxPasswordLength,
	}
}
