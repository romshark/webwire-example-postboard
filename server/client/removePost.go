package client

import (
	"context"

	"github.com/qbeon/webwire-example-postboard/server/apisrv/api"
)

// RemovePost implements the ApiClient interface
func (c *apiClient) RemovePost(
	ctx context.Context,
	params api.RemovePostParams,
) error {
	return c.Query(ctx, api.RemovePost, params, nil)
}
