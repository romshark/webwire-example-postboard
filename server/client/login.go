package client

import (
	"context"

	"github.com/qbeon/webwire-messenger/server/apisrv/api"
)

// Login implements the ApiClient interface
func (c *apiClient) Login(
	ctx context.Context,
	params api.LoginParams,
) error {
	return c.Query(ctx, api.Login, params, nil)
}
