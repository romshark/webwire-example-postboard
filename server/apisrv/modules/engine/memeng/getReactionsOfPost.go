package memeng

import (
	"github.com/qbeon/webwire-messenger/server/apisrv/api"
)

// GetReactionsOfPost implements the Engine interface
func (eng *engine) GetReactionsOfPost(
	postIdent api.Identifier,
) ([]*api.PostReaction, error) {
	// Lock the engine store to execute the following operation transactionally
	eng.lock.Lock()
	defer eng.lock.Unlock()

	// Find the actual post
	post, err := eng.findPostByIdent(postIdent)
	if err != nil {
		return nil, err
	}

	// Preallocate buffer
	result := make([]*api.PostReaction, 0, len(post.Reactions))

	// Filter out all archived posts
	for _, postReaction := range post.Reactions {
		if !postReaction.Archived {
			result = append(result, &postReaction.Reaction)
		}
	}

	return result, nil
}
