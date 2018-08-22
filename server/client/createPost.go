package client

import (
	"context"

	"github.com/qbeon/webwire-messenger/server/apisrv/api"
)

// CreatePost implements ApiClient interface
func (c *apiClient) CreatePost(
	ctx context.Context,
	params api.CreatePostParams,
) (api.Identifier, error) {
	var result api.CreatePostReturn
	if err := c.Query(ctx, api.CreatePost, params, &result); err != nil {
		return api.Identifier{}, err
	}
	return result.PostIdent, nil
}
