package validator

// ReactionDescription implements the Validator interface
func (vld *validator) ReactionDescription(description string) error {
	length := len(description)
	if uint32(length) < vld.MinReactionDescriptionLength {
		return Errorf(
			"invalid post reaction description: too short (%d / %d)",
			length,
			vld.MinReactionDescriptionLength,
		)
	} else if uint32(length) > vld.MaxReactionDescriptionLength {
		return Errorf(
			"invalid post reaction description: too long (%d / %d)",
			length,
			vld.MaxReactionDescriptionLength,
		)
	}
	return nil
}
