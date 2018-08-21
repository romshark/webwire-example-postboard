package memeng

import (
	"errors"

	"github.com/qbeon/webwire-messenger/server/apisrv/api"
)

// postMessageReaction sets a new reaction to the message identified by
// 'messageIdent'
// A not-found error is returned if it either doesn't exist
// or if it was already deleted
func (eng *engine) postMessageReaction(
	messageIdent api.Identifier,
	reaction *api.MessageReaction,
) error {
	if reaction == nil {
		return errors.New("can't post a nil message reaction")
	}

	message, err := eng.findMessageByIdent(messageIdent)
	if err != nil {
		return err
	}

	// Set the reaction
	message.Message.Reactions = append(message.Message.Reactions, *reaction)

	return nil
}

// PostMessageReaction implements the Engine interface
func (eng *engine) PostMessageReaction(
	messageIdent api.Identifier,
	reaction *api.MessageReaction,
) error {
	// Lock the engine store to execute the following operation transactionally
	eng.lock.Lock()
	defer eng.lock.Unlock()

	return eng.postMessageReaction(messageIdent, reaction)
}
