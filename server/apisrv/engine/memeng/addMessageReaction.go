package memeng

import (
	"errors"

	engiface "github.com/qbeon/webwire-messenger/server/apisrv/engine"
)

// postMessageReaction sets a new reaction to the message identified by
// 'messageIdent'
// A not-found error is returned if it either doesn't exist
// or if it was already deleted
func (eng *engine) postMessageReaction(
	messageIdent engiface.Identifier,
	reaction *engiface.MessageReaction,
) error {
	if reaction == nil {
		return errors.New("can't post a nil message reaction")
	}

	message, err := eng.findMessageByIdent(messageIdent)
	if err != nil {
		return err
	}

	// Set the reaction
	message.Reactions[reaction.Author] = reaction

	return nil
}

// PostMessageReaction implements the Engine interface
func (eng *engine) PostMessageReaction(
	messageIdent engiface.Identifier,
	reaction *engiface.MessageReaction,
) error {
	// Lock the engine store to execute the following operation transactionally
	eng.lock.Lock()
	defer eng.lock.Unlock()

	return eng.postMessageReaction(messageIdent, reaction)
}
