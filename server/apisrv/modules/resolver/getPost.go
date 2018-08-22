package resolver

import (
	"github.com/qbeon/webwire-example-postboard/server/apisrv/api"
	"github.com/qbeon/webwire-example-postboard/server/apisrv/sessinfo"
)

// GetPost implements the Resolver interface
func (rsv *resolver) GetPost(
	session *sessinfo.SessionInfo,
	params *api.GetPostParams,
) (interface{}, error) {
	// Instruct the engine to find the post
	posts, err := rsv.engine.GetPosts([]api.Identifier{params.Ident})
	if err != nil {
		return nil, err
	}

	// Post successfully retrieved
	return posts[0], nil
}
