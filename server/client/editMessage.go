package client

import (
	"context"

	"github.com/qbeon/webwire-messenger/server/apisrv/api"
)

// EditMessage implements the ApiClient
func (c *apiClient) EditMessage(
	ctx context.Context,
	params api.EditMessageParams,
) error {
	return c.Query(ctx, api.EditMessage, params, nil)
}
