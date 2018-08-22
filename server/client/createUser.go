package client

import (
	"context"

	"github.com/qbeon/webwire-example-postboard/server/apisrv/api"
)

// CreateUser implements the ApiClient interface
func (c *apiClient) CreateUser(
	ctx context.Context,
	params api.CreateUserParams,
) (api.Identifier, error) {
	var result api.CreateUserReturn
	if err := c.Query(ctx, api.CreateUser, params, &result); err != nil {
		return api.Identifier{}, err
	}
	return result.UserIdent, nil
}
