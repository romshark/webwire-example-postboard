package client

import (
	"context"

	"github.com/qbeon/webwire-example-postboard/server/apisrv/api"
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
	if profile.Identifier.IsNull() {
		return nil, nil
	}
	return profile, nil
}
