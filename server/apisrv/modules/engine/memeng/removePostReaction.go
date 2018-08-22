package memeng

import (
	"github.com/qbeon/webwire-messenger/server/apisrv/api"
)

// removePostReaction marks the post reaction identified by `ident` as archived.
// A `ErrPostReactionNotFound` error is returned if it either doesn't exist
// or if it was already removed
func (eng *engine) removePostReaction(ident api.Identifier) error {
	postReactionRef, err := eng.findPostReactionByIdent(ident)
	if err != nil {
		return err
	}

	// Mark the post as archived
	postReactionRef.Archived = true

	return nil
}

// RemovePostReaction implements the Engine interface
func (eng *engine) RemovePostReaction(ident api.Identifier) error {
	// Lock the engine store to execute the following operation transactionally
	eng.lock.Lock()
	defer eng.lock.Unlock()

	return eng.removePostReaction(ident)
}
