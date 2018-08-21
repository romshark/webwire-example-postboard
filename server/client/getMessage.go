package client

import (
	"context"

	"github.com/qbeon/webwire-messenger/server/apisrv/api"
)

// GetMessage implements the ApiClient interface
func (c *apiClient) GetMessage(
	ctx context.Context,
	params api.GetMessageParams,
) (*api.Message, error) {
	result := &api.Message{}
	if err := c.Query(ctx, api.GetMessage, params, result); err != nil {
		return nil, err
	}
	return result, nil
}
