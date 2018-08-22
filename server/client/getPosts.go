package client

import (
	"context"

	"github.com/qbeon/webwire-messenger/server/apisrv/api"
)

// GetPosts implements the ApiClient interface
func (c *apiClient) GetPosts(
	ctx context.Context,
	params api.GetPostsParams,
) ([]*api.Post, error) {
	result := make([]*api.Post, 0)
	if err := c.Query(ctx, api.GetPosts, params, &result); err != nil {
		return nil, err
	}
	return result, nil
}
