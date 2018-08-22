package resolver

import (
	"time"

	"github.com/qbeon/webwire-messenger/server/apisrv/api"
	"github.com/qbeon/webwire-messenger/server/apisrv/modules/authorizer"
	"github.com/qbeon/webwire-messenger/server/apisrv/sessinfo"
)

// CreatePostReaction implements the Resolver interface
func (rsv *resolver) CreatePostReaction(
	session *sessinfo.SessionInfo,
	params *api.CreatePostReactionParams,
) (interface{}, error) {
	if err := rsv.authorizer.MeetsAll(
		session,
		authorizer.IsAuthenticated(
			"only administrators and authenticated users are allowed "+
				"to create post reactions",
		),
	); err != nil {
		return nil, err
	}

	// Perform an engine call to retrieve the targeted post
	// to verify authorization
	retrieved, err := rsv.engine.GetPosts(
		[]api.Identifier{params.PostIdent},
	)
	if err != nil {
		return nil, err
	}

	// Check authorization, ensure the user is not the owner of this
	// post because users are not allowed to react on their on posts
	if err := rsv.authorizer.MeetsAll(
		session,
		authorizer.IsNotResourceOwner(
			retrieved[0].Author,
			"users are not allowed to post reactions on their own posts",
		),
	); err != nil {
		return nil, err
	}

	// Validate reaction description
	if err := rsv.validator.ReactionDescription(
		params.Description,
	); err != nil {
		return nil, err
	}

	newIdent := api.NewIdentifier()

	// Instruct the engine to add the new post reaction
	if err := rsv.engine.CreatePostReaction(
		params.PostIdent,
		&api.PostReaction{
			Ident:       newIdent,
			Author:      session.UserIdentifier,
			Type:        params.Type,
			Description: params.Description,
			Creation:    time.Now().UTC(),
		},
	); err != nil {
		return nil, err
	}

	return api.CreatePostReactionReturn{
		PostReactionIdent: newIdent,
	}, nil
}
