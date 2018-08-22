package memeng

import (
	"github.com/qbeon/webwire-example-postboard/server/apisrv/api"
)

// GetPostReactions implements the Engine interface
func (eng *engine) GetPostReactions(
	identifiers []api.Identifier,
) ([]*api.PostReaction, error) {
	// Lock the engine store to execute the following operation transactionally
	eng.lock.Lock()
	defer eng.lock.Unlock()

	// Preallocate buffer
	result := make([]*api.PostReaction, len(identifiers))

	// Search by the postReaction.identifier index
	// for each requested identifier
	for i, searchedIdent := range identifiers {
		postReaction, exists := eng.postReactionsByIdent[searchedIdent]
		if exists {
			// Skip archived posts
			if !postReaction.Archived {
				result[i] = &postReaction.Reaction
			}
		} else {
			// Reset inexistent identifier to nil post
			result[i] = nil
		}
	}

	return result, nil
}
