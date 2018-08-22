package memeng

import (
	"github.com/qbeon/webwire-messenger/server/apisrv/api"
	engiface "github.com/qbeon/webwire-messenger/server/apisrv/modules/engine"
)

// findPostByIdent finds and returns the post identified by `ident`.
// Returns a `ErrPostNotFound` error if the post doesn't exist
// or is archived
func (eng *engine) findPostByIdent(
	ident api.Identifier,
) (*Post, error) {
	postIndex, exists := eng.postsByIdent[ident]
	if !exists {
		// Return not-found error if the post doesn't exist
		return nil, engiface.NewError(engiface.ErrPostNotFound)
	}
	post := eng.posts[postIndex]
	if post.Archived {
		// Return not-found error if the post is archived
		return nil, engiface.NewError(engiface.ErrPostNotFound)
	}
	return post, nil
}
