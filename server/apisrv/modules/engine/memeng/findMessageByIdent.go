package memeng

import (
	"github.com/qbeon/webwire-messenger/server/apisrv/api"
	engiface "github.com/qbeon/webwire-messenger/server/apisrv/modules/engine"
)

// findMessageByIdent finds and returns the message identified by 'ident'.
// Returns a not-found error if the message doesn't exist or is archived
func (eng *engine) findMessageByIdent(
	ident api.Identifier,
) (*Message, error) {
	messageIndex, exists := eng.messagesByIdent[ident]
	if !exists {
		// Return not-found error if the message doesn't exist
		return nil, engiface.NewError(engiface.ErrMessageNotFound)
	}
	message := eng.messages[messageIndex]
	if message.Archived {
		// Return not-found error if the message is archived
		return nil, engiface.NewError(engiface.ErrMessageNotFound)
	}
	return message, nil
}
