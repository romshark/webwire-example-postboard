package resolver

import (
	"github.com/qbeon/webwire-messenger/server/apisrv/api"
	"github.com/qbeon/webwire-messenger/server/apisrv/modules/authorizer"
	"github.com/qbeon/webwire-messenger/server/apisrv/sessinfo"
)

// RemovePost implements the Resolver interface
func (rsv *resolver) RemovePost(
	session *sessinfo.SessionInfo,
	params *api.RemovePostParams,
) (interface{}, error) {
	// Check authorization, ensure the user is not a guest,
	// because guests are allowed to read only
	if err := rsv.authorizer.MeetsAll(
		session,
		authorizer.IsAuthenticated(
			"only authenticated clients are allowed to remove posts",
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

	// Check authorization (ownership)
	if err := rsv.authorizer.MeetsEitherOf(
		session,
		authorizer.IsAdmin(
			"only administrators and the author of the post "+
				"are allowed to remove a post",
		),
		authorizer.IsResourceOwner(
			retrieved[0].Author,
			"only administrators and the author of the post "+
				"are allowed to remove a post",
		),
	); err != nil {
		return nil, err
	}

	// Instruct the engine to remove the post
	if err := rsv.engine.RemovePost(params.PostIdent); err != nil {
		return nil, err
	}

	// Post successfully removed
	return nil, nil
}
