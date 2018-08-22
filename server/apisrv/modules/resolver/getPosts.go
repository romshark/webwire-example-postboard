package resolver

import (
	"github.com/qbeon/webwire-example-postboard/server/apisrv/api"
	"github.com/qbeon/webwire-example-postboard/server/apisrv/sessinfo"
)

// GetPosts implements the Resolver interface
func (rsv *resolver) GetPosts(
	session *sessinfo.SessionInfo,
	params *api.GetPostsParams,
) (interface{}, error) {
	// Validate query limit
	if err := rsv.validator.PostsQueryLimit(params.Limit); err != nil {
		return nil, err
	}

	// Instruct the engine to get n number of posts after the specified one
	result, err := rsv.engine.GetPostsAfter(
		params.After,
		params.Limit,
	)
	if err != nil {
		return nil, err
	}

	// Posts successfully retrieved
	return result, nil
}
