package validator

// Username implements the Validator interface
func (vld *validator) Username(username string) error {
	if !vld.UsernameRule.Match([]byte(username)) {
		return Errorf(
			"username (%s) doesn't respect the username rules",
			username,
		)
	}
	return nil
}
