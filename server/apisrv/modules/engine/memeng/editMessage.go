package memeng

import (
	"time"

	"github.com/qbeon/webwire-messenger/server/apisrv/api"
)

// editMessage changes the contents of the message identified by 'ident'
// to 'newContents' and updates its last edit time to 'editTime'
func (eng *engine) editMessage(
	ident api.Identifier,
	editTime time.Time,
	newContents string,
) error {
	message, err := eng.findMessageByIdent(ident)
	if err != nil {
		return err
	}

	// Update contents and lastEdit fields
	message.Message.Contents = newContents
	message.Message.LastEdit = &editTime

	return nil
}

// EditMessage implements the Engine interface
func (eng *engine) EditMessage(
	ident api.Identifier,
	editTime time.Time,
	newContents string,
) error {
	// Lock the engine store to execute the following operation transactionally
	eng.lock.Lock()
	defer eng.lock.Unlock()

	return eng.editMessage(ident, editTime, newContents)
}
