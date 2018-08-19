package validator

import (
	"regexp"

	"github.com/pkg/errors"
)

// validator represents an implementation of the Validator interface
type validator struct {
	NameRule                     *regexp.Regexp
	MinMessageLength             uint32
	MaxMessageLength             uint32
	MinReactionDescriptionLength uint32
	MaxReactionDescriptionLength uint32
	MaxMessagesQueryLimit        uint32
}

func NewValidator(
	nameRule *regexp.Regexp,
	minMessageLength,
	maxMessageLength,
	minReactionDescriptionLength,
	maxReactionDescriptionLength,
	maxMessagesQueryLimit uint32,
) Validator {
	if nameRule == nil {
		panic(errors.New("missing validator name rule"))
	}

	return &validator{
		NameRule:                     nameRule,
		MinMessageLength:             minMessageLength,
		MaxMessageLength:             maxMessageLength,
		MinReactionDescriptionLength: minReactionDescriptionLength,
		MaxReactionDescriptionLength: maxReactionDescriptionLength,
		MaxMessagesQueryLimit:        maxMessagesQueryLimit,
	}
}
