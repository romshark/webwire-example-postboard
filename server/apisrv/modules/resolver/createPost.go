package resolver

import (
	"time"

	"github.com/qbeon/webwire-example-postboard/server/apisrv/api"
	"github.com/qbeon/webwire-example-postboard/server/apisrv/modules/authorizer"
	"github.com/qbeon/webwire-example-postboard/server/apisrv/sessinfo"
)

// CreatePost implements the Resolver interface
func (rsv *resolver) CreatePost(
	session *sessinfo.SessionInfo,
	params *api.CreatePostParams,
) (interface{}, error) {
	// Check authorization
	if err := rsv.authorizer.MeetsAll(
		session,
		authorizer.IsAuthenticated(
			"only authenticated clients are allows to create posts",
		),
	); err != nil {
		return nil, err
	}

	newIdent := api.NewIdentifier()

	// Instruct the engine to create a new post
	if err := rsv.engine.CreatePost(&api.Post{
		Ident:       newIdent,
		Author:      session.UserIdentifier,
		Publication: time.Now().UTC(),
		Contents:    params.Contents,
	}); err != nil {
		return nil, err
	}

	// Post successfully created
	return api.CreatePostReturn{
		PostIdent: newIdent,
	}, nil
}
