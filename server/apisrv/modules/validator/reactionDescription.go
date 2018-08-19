package validator

import "github.com/pkg/errors"

// ReactionDescription implements the Validator interface
func (vld *validator) ReactionDescription(description string) error {
	length := len(description)
	if uint32(length) < vld.MinReactionDescriptionLength {
		return errors.Errorf(
			"invalid message reaction description: too short (%d / %d)",
			length,
			vld.MinReactionDescriptionLength,
		)
	} else if uint32(length) > vld.MaxReactionDescriptionLength {
		return errors.Errorf(
			"invalid message reaction description: too long (%d / %d)",
			length,
			vld.MaxReactionDescriptionLength,
		)
	}
	return nil
}
