package client

import (
	"context"

	"github.com/qbeon/webwire-example-postboard/server/apisrv/api"
)

// EditPost implements the ApiClient
func (c *apiClient) EditPost(
	ctx context.Context,
	params api.EditPostParams,
) error {
	return c.Query(ctx, api.EditPost, params, nil)
}
