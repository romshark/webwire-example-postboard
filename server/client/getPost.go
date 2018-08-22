package client

import (
	"context"

	"github.com/qbeon/webwire-messenger/server/apisrv/api"
)

// GetPost implements the ApiClient interface
func (c *apiClient) GetPost(
	ctx context.Context,
	params api.GetPostParams,
) (*api.Post, error) {
	result := &api.Post{}
	if err := c.Query(ctx, api.GetPost, params, result); err != nil {
		return nil, err
	}
	if result.Identifier.IsNull() {
		return nil, nil
	}
	return result, nil
}
