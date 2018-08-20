package client

import (
	"context"

	"github.com/qbeon/webwire-messenger/server/apisrv/api"
)

// PostMessage implements ApiClient interface
func (c *apiClient) PostMessage(
	ctx context.Context,
	params api.PostMessageParams,
) (api.Identifier, error) {
	var result api.PostMessageReturn
	if err := c.Query(ctx, api.PostMessage, params, &result); err != nil {
		return api.Identifier{}, err
	}
	return result.MessageIdent, nil
}
