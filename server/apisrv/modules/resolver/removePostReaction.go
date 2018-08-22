package resolver

import (
	"github.com/qbeon/webwire-example-postboard/server/apisrv/api"
	"github.com/qbeon/webwire-example-postboard/server/apisrv/modules/authorizer"
	engiface "github.com/qbeon/webwire-example-postboard/server/apisrv/modules/engine"
	"github.com/qbeon/webwire-example-postboard/server/apisrv/sessinfo"
)

// RemovePostReaction implements the Resolver interface
func (rsv *resolver) RemovePostReaction(
	session *sessinfo.SessionInfo,
	params *api.RemovePostReactionParams,
) (interface{}, error) {
	// Check authorization, ensure the user is not a guest,
	// because guests are allowed to read only
	if err := rsv.authorizer.MeetsAll(
		session,
		authorizer.IsAuthenticated(
			"only authenticated clients are allowed to remove post reactions",
		),
	); err != nil {
		return nil, err
	}

	// Perform an engine call to retrieve the targeted reaction
	// to verify authorization
	retrieved, err := rsv.engine.GetPostReactions(
		[]api.Identifier{params.ReactionIdent},
	)
	if err != nil {
		return nil, err
	}
	if retrieved == nil {
		return nil, engiface.NewError(engiface.ErrPostReactionNotFound)
	}

	// Check authorization (ownership)
	if err := rsv.authorizer.MeetsEitherOf(
		session,
		authorizer.IsAdmin(
			"only administrators and the author of the post reaction "+
				"are allowed to remove a post reaction",
		),
		authorizer.IsResourceOwner(
			retrieved[0].Author,
			"only administrators and the author of the post reaction "+
				"are allowed to remove a post reaction",
		),
	); err != nil {
		return nil, err
	}

	// Instruct the engine to remove the post reaction
	if err := rsv.engine.RemovePostReaction(params.ReactionIdent); err != nil {
		return nil, err
	}

	// Post reaction successfully removed
	return nil, nil
}
