package validator

// PostContents implements the Validator interface
func (vld *validator) PostContents(contents string) error {
	length := len(contents)
	if uint32(length) < vld.MinPostLength {
		return Errorf(
			"invalid post contents: too short (%d / %d)",
			length,
			vld.MinPostLength,
		)
	} else if uint32(length) > vld.MaxPostLength {
		return Errorf(
			"invalid post contents: too long (%d / %d)",
			length,
			vld.MaxPostLength,
		)
	}
	return nil
}
