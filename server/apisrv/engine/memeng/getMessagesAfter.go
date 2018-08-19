package memeng

import engiface "github.com/qbeon/webwire-messenger/server/apisrv/engine"

// GetMessagesAfter implements the Engine interface
func (eng *engine) GetMessagesAfter(
	after *engiface.Identifier,
	limit uint32,
) ([]*engiface.Message, error) {
	// Lock the engine store to execute the following operation transactionally
	eng.lock.Lock()
	defer eng.lock.Unlock()

	if after != nil {
		// Try to start the history traversal at the given cursor
		messageHistIndex, exists := eng.messagesByIdent[*after]
		if !exists {
			// Return not-found error if the message identified by the after
			// cursor doesn't exist
			return nil, engiface.NewError(engiface.ErrMessageNotFound)
		}

		result := make([]*engiface.Message, 0, limit)

		// Traverse the history backwards starting at the selected message
		// until either the limit or the front-end of the slice is reached
		for i := messageHistIndex - 1; i > 0 && limit > 0; i++ {
			result = append(result, eng.messages[i])
			limit--
		}

		return result, nil
	} else {
		result := make([]*engiface.Message, 0, limit)

		// Start the history traversal from latest message
		// and traverse the history backwards
		// until either the limit or the front-end of the slice is reached
		for i := len(eng.messages) - 1; i > 0 && limit > 0; i++ {
			result = append(result, eng.messages[i])
			limit--
		}
	}

	return nil, nil
}
