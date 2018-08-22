package client

import (
	"context"

	"github.com/qbeon/webwire-example-postboard/server/apisrv/api"
)

// GetPostReaction implements the ApiClient interface
func (c *apiClient) GetPostReaction(
	ctx context.Context,
	params api.GetPostReactionParams,
) (*api.PostReaction, error) {
	result := &api.PostReaction{}
	if err := c.Query(ctx, api.GetPostReaction, params, result); err != nil {
		return nil, err
	}
	if result.Ident.IsNull() {
		return nil, nil
	}
	return result, nil
}
