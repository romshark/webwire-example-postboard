package validator

// PostsQueryLimit implements the Validator interface
func (vld *validator) PostsQueryLimit(limit uint32) error {
	if limit > vld.MaxPostsQueryLimit {
		return Errorf(
			"posts query limit exceeds max limit: (%d / %d)",
			limit,
			vld.MaxPostsQueryLimit,
		)
	}

	return nil
}
