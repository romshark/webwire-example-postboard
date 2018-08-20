package client

import (
	"context"

	"github.com/qbeon/webwire-messenger/server/apisrv/api"
)

// GetMessages implements the ApiClient interface
func (c *apiClient) GetMessages(
	ctx context.Context,
	params api.GetMessagesParams,
) ([]*api.Message, error) {
	var result api.GetMessagesReturn
	if err := c.Query(ctx, api.GetMessages, params, result); err != nil {
		return nil, err
	}
	return result.Messages, nil
}
