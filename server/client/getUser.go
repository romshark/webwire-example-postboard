package client

import (
	"context"

	"github.com/qbeon/webwire-messenger/server/apisrv/api"
)

// GetUser implements the ApiClient interface
func (c *apiClient) GetUser(
	ctx context.Context,
	params api.GetUserParams,
) (*api.User, error) {
	profile := &api.User{}
	if err := c.Query(ctx, api.GetUser, params, profile); err != nil {
		return nil, err
	}
	return profile, nil
}
