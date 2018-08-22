package client

import (
	"context"

	"github.com/qbeon/webwire-example-postboard/server/apisrv/api"
)

// RemovePostReaction implements the ApiClient interface
func (c *apiClient) RemovePostReaction(
	ctx context.Context,
	params api.RemovePostReactionParams,
) error {
	return c.Query(ctx, api.RemovePostReaction, params, nil)
}
