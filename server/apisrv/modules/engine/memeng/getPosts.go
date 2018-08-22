package memeng

import (
	"github.com/qbeon/webwire-example-postboard/server/apisrv/api"
)

// GetPosts implements the Engine interface
func (eng *engine) GetPosts(
	identifiers []api.Identifier,
) ([]*api.Post, error) {
	// Lock the engine store to execute the following operation transactionally
	eng.lock.Lock()
	defer eng.lock.Unlock()

	// Preallocate buffer
	result := make([]*api.Post, len(identifiers))

	// Search by the post.identifier index for each requested identifier
	for i, searchedIdent := range identifiers {
		postIndex, exists := eng.postsByIdent[searchedIdent]
		if exists {
			post := eng.posts[postIndex]
			// Skip archived posts
			if !post.Archived {
				result[i] = &post.Post
			}
		} else {
			// Reset inexistent identifier to nil post
			result[i] = nil
		}
	}

	return result, nil
}
