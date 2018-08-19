package memeng

import (
	"github.com/qbeon/webwire-messenger/server/apisrv/api"
)

// GetMessages implements the Engine interface
func (eng *engine) GetMessages(
	identifiers []api.Identifier,
) ([]*api.Message, error) {
	// Lock the engine store to execute the following operation transactionally
	eng.lock.Lock()
	defer eng.lock.Unlock()

	// Preallocate message buffer
	result := make([]*api.Message, len(identifiers))

	// Search by the message.identifier index for each requested identifier
	for i, searchedIdent := range identifiers {
		messageIndex, exists := eng.messagesByIdent[searchedIdent]
		if exists {
			message := eng.messages[messageIndex]
			// Skip archived messages
			if !message.Archived {
				result[i] = &message.Message
			}
		} else {
			// Reset inexistent identifier to nil message
			result[i] = nil
		}
	}

	return nil, nil
}
