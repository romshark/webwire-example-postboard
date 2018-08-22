package memeng

import (
	"time"

	"github.com/qbeon/webwire-messenger/server/apisrv/api"
)

// editPost changes the contents of the post identified by `ident`
// to `newContents` and updates its last edit time to 'editTime'
func (eng *engine) editPost(
	ident api.Identifier,
	editTime time.Time,
	newContents string,
) error {
	post, err := eng.findPostByIdent(ident)
	if err != nil {
		return err
	}

	// Update contents and lastEdit fields
	post.Post.Contents = newContents
	post.Post.LastEdit = &editTime

	return nil
}

// EditPost implements the Engine interface
func (eng *engine) EditPost(
	ident api.Identifier,
	editTime time.Time,
	newContents string,
) error {
	// Lock the engine store to execute the following operation transactionally
	eng.lock.Lock()
	defer eng.lock.Unlock()

	return eng.editPost(ident, editTime, newContents)
}
