package memeng

import "github.com/qbeon/webwire-messenger/server/apisrv/api"

// removePost marks the post identified by `ident` as archived.
// A `ErrPostNotFound` error is returned if it either doesn't exist
// or if it was already removed
func (eng *engine) removePost(ident api.Identifier) error {
	post, err := eng.findPostByIdent(ident)
	if err != nil {
		return err
	}

	// Mark the post as archived
	post.Archived = true

	return nil
}

// RemovePost implements the Engine interface
func (eng *engine) RemovePost(
	ident api.Identifier,
) error {
	// Lock the engine store to execute the following operation transactionally
	eng.lock.Lock()
	defer eng.lock.Unlock()

	return eng.removePost(ident)
}
