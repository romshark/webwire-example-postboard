package validator

import "github.com/pkg/errors"

// Name implements the Validator interface
func (vld *validator) Name(name string) error {
	if !vld.NameRule.Match([]byte(name)) {
		return errors.Errorf(
			"name (%s) doesn't respect the name rule",
			name,
		)
	}
	return nil
}
