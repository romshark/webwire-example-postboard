package memeng

import (
	"github.com/qbeon/webwire-messenger/server/apisrv/api"
	engiface "github.com/qbeon/webwire-messenger/server/apisrv/modules/engine"
)

// findPostReactionByIdent finds and returns the post reaction
// identified by `ident`.
// Returns a `ErrPostReactionNotFound` error if the post reaction doesn't exist
// or is archived
func (eng *engine) findPostReactionByIdent(
	ident api.Identifier,
) (*PostReaction, error) {
	postReactionRef, exists := eng.postReactionsByIdent[ident]
	if !exists {
		// Return not-found error if the post reaction doesn't exist
		return nil, engiface.NewError(engiface.ErrPostReactionNotFound)
	}
	if postReactionRef.Archived {
		// Return not-found error if the post reaction is archived
		return nil, engiface.NewError(engiface.ErrPostReactionNotFound)
	}
	return postReactionRef, nil
}
