package client

import (
	"context"

	"github.com/qbeon/webwire-messenger/server/apisrv/api"
)

// RemoveMessage implements the ApiClient interface
func (c *apiClient) RemoveMessage(
	ctx context.Context,
	params api.RemoveMessageParams,
) error {
	return c.Query(ctx, api.RemoveMessage, params, nil)
}
