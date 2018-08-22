package client

import (
	"context"

	"github.com/qbeon/webwire-example-postboard/server/apisrv/api"
)

// GetReactionsOfPost implements the ApiClient interface
func (c *apiClient) GetReactionsOfPost(
	ctx context.Context,
	params api.GetReactionsOfPostParams,
) ([]*api.PostReaction, error) {
	result := make([]*api.PostReaction, 0)
	if err := c.Query(
		ctx,
		api.GetReactionsOfPost,
		params,
		&result,
	); err != nil {
		return nil, err
	}
	return result, nil
}
