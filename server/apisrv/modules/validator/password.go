package validator

// Password implements the Validator interface
func (vld *validator) Password(password string) error {
	length := len(password)
	if uint32(length) < vld.MinPasswordLength {
		return Errorf(
			"invalid password: too short (%d / %d)",
			length,
			vld.MinPasswordLength,
		)
	} else if uint32(length) > vld.MaxPasswordLength {
		return Errorf(
			"invalid password: too long (%d / %d)",
			length,
			vld.MaxPasswordLength,
		)
	}
	return nil
}
