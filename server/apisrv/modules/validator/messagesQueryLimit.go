package validator

// MessagesQueryLimit implements the Validator interface
func (vld *validator) MessagesQueryLimit(limit uint32) error {
	if limit > vld.MaxMessagesQueryLimit {
		return Errorf(
			"messages query limit exceeds max limit: (%d / %d)",
			limit,
			vld.MaxMessagesQueryLimit,
		)
	}

	return nil
}
