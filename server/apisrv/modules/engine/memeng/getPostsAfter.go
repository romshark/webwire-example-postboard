package memeng

import (
	"github.com/qbeon/webwire-example-postboard/server/apisrv/api"
	engiface "github.com/qbeon/webwire-example-postboard/server/apisrv/modules/engine"
)

// GetPostsAfter implements the Engine interface
func (eng *engine) GetPostsAfter(
	after *api.Identifier,
	limit uint32,
) ([]*api.Post, error) {
	// Lock the engine store to execute the following operation transactionally
	eng.lock.Lock()
	defer eng.lock.Unlock()

	// Preallocate buffer
	result := make([]*api.Post, 0, limit)

	// Begin at the latest post by default
	startIndex := len(eng.posts) - 1

	if after != nil {
		// Try to start the history traversal at the given cursor
		histIndex, exists := eng.postsByIdent[*after]
		if !exists {
			// Return not-found error if the post identified by the after
			// cursor doesn't exist
			return nil, engiface.NewError(engiface.ErrPostNotFound)
		}
		startIndex = histIndex - 1
	}

	// Traverse the history backwards starting at the selected post
	// until either the limit or the front-end of the slice is reached
	for i := startIndex; i >= 0 && limit > 0; i, limit = i-1, limit-1 {
		post := eng.posts[i]
		// Skip archived posts
		if !post.Archived {
			result = append(result, &post.Post)
		}
	}

	return result, nil
}
