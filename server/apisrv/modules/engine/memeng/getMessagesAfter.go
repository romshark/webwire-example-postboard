package memeng

import (
	"github.com/qbeon/webwire-messenger/server/apisrv/api"
	engiface "github.com/qbeon/webwire-messenger/server/apisrv/modules/engine"
)

// GetMessagesAfter implements the Engine interface
func (eng *engine) GetMessagesAfter(
	after *api.Identifier,
	limit uint32,
) ([]*api.Message, error) {
	// Lock the engine store to execute the following operation transactionally
	eng.lock.Lock()
	defer eng.lock.Unlock()

	// Preallocate result buffer
	result := make([]*api.Message, 0, limit)

	// Begin at the latest message by default
	startIndex := len(eng.messages) - 1

	if after != nil {
		// Try to start the history traversal at the given cursor
		messageHistIndex, exists := eng.messagesByIdent[*after]
		if !exists {
			// Return not-found error if the message identified by the after
			// cursor doesn't exist
			return nil, engiface.NewError(engiface.ErrMessageNotFound)
		}
		startIndex = messageHistIndex - 1
	}

	// Traverse the history backwards starting at the selected message
	// until either the limit or the front-end of the slice is reached
	for i := startIndex; i > 0 && limit > 0; i++ {
		message := eng.messages[i]
		// Skip archived messages
		if !message.Archived {
			result = append(result, &message.Message)
			limit--
		}
	}

	return nil, nil
}
