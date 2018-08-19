package validator

import "github.com/pkg/errors"

// MessagesQueryLimit implements the Validator interface
func (vld *validator) MessagesQueryLimit(limit uint32) error {
	if limit > vld.MaxMessagesQueryLimit {
		return errors.Errorf(
			"messages query limit exceeds max limit: (%d / %d)",
			limit,
			vld.MaxMessagesQueryLimit,
		)
	}

	return nil
}
