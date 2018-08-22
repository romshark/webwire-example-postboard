package resolver

import (
	"time"

	"github.com/qbeon/webwire-example-postboard/server/apisrv/api"
	"github.com/qbeon/webwire-example-postboard/server/apisrv/modules/authorizer"
	"github.com/qbeon/webwire-example-postboard/server/apisrv/sessinfo"
)

// EditPost implements the Resolver interface
func (rsv *resolver) EditPost(
	session *sessinfo.SessionInfo,
	params *api.EditPostParams,
) (interface{}, error) {
	// Validate new contents before proceeding
	if err := rsv.validator.PostContents(params.NewContents); err != nil {
		return nil, err
	}

	// Check authorization
	if err := rsv.authorizer.MeetsAll(
		session,
		authorizer.IsAuthenticated("guest clients are allows to read only"),
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

	// Check authorization
	if err := rsv.authorizer.MeetsAll(
		session,
		authorizer.IsResourceOwner(
			retrieved[0].Author,
			"only authors can edit their posts",
		),
	); err != nil {
		return nil, err
	}

	// Instruct the engine to edit the post
	if err := rsv.engine.EditPost(
		params.PostIdent,
		time.Now().UTC(),
		params.NewContents,
	); err != nil {
		return nil, err
	}

	// Post successfully edited
	return nil, nil
}
