package validator

// Name implements the Validator interface
func (vld *validator) Name(name string) error {
	if !vld.NameRule.Match([]byte(name)) {
		return Errorf(
			"name (%s) doesn't respect the name rule",
			name,
		)
	}
	return nil
}
