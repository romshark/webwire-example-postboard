package memeng

import engiface "github.com/qbeon/webwire-messenger/server/apisrv/engine"

// GetMessages implements the Engine interface
func (eng *engine) GetMessages(
	identifiers []engiface.Identifier,
) ([]*engiface.Message, error) {
	// Lock the engine store to execute the following operation transactionally
	eng.lock.Lock()
	defer eng.lock.Unlock()

	// Preallocate message buffer
	result := make([]*engiface.Message, len(identifiers))

	// Search by the message.identifier index for each requested identifier
	for i, searchedIdent := range identifiers {
		messageIndex, exists := eng.messagesByIdent[searchedIdent]
		if exists {
			result[i] = eng.messages[messageIndex]
		} else {
			// Reset inexistent identifier to nil message
			result[i] = nil
		}
	}

	return nil, nil
}
