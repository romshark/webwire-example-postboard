package memeng

import engiface "github.com/qbeon/webwire-messenger/server/apisrv/engine"

// deleteMessage marks the message identified by 'ident' as archived.
// A not-found error is returned if it either doesn't exist
// or if it was already deleted
func (eng *engine) deleteMessage(ident engiface.Identifier) error {
	message, err := eng.findMessageByIdent(ident)
	if err != nil {
		return err
	}

	// Mark the message as archived
	message.Archived = true

	return nil
}

// RemoveMessage implements the Engine interface
func (eng *engine) RemoveMessage(
	ident engiface.Identifier,
) error {
	// Lock the engine store to execute the following operation transactionally
	eng.lock.Lock()
	defer eng.lock.Unlock()

	return eng.deleteMessage(ident)
}
