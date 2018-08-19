package memeng

import (
	"errors"

	"github.com/qbeon/webwire-messenger/server/apisrv/api"
	engiface "github.com/qbeon/webwire-messenger/server/apisrv/modules/engine"
)

// postMessage posts the new message 'newMessage' to the store
// updating the identifier index.
// Returns an error if a message with a similar identifier already exists
func (eng *engine) postMessage(newMessage *api.Message) error {
	if newMessage == nil {
		return errors.New("can't post nil message")
	}

	// Return error if a message with a similar identifier is already defined,
	// ignore the returned error because it'll definitely be a not-found error
	message, _ := eng.findMessageByIdent(newMessage.Identifier)
	if message != nil {
		return engiface.NewError(engiface.ErrMessageAlreadyExists)
	}

	// Commit the message to the chronological history
	eng.messages = append(eng.messages, &Message{
		Message:  *newMessage,
		Archived: false,
	})

	// Update message identifier index
	eng.messagesByIdent[newMessage.Identifier] = len(eng.messages) - 1

	return nil
}

// PostMessage implements the Engine interface
func (eng *engine) PostMessage(newMessage *api.Message) error {
	// Lock the engine store to execute the following operation transactionally
	eng.lock.Lock()
	defer eng.lock.Unlock()

	return eng.postMessage(newMessage)
}
