package client

import (
	"context"

	"github.com/qbeon/webwire-messenger/server/apisrv/api"
)

// PostMessageReaction implements the ApiClient interface
func (c *apiClient) PostMessageReaction(
	ctx context.Context,
	params api.PostMessageReactionParams,
) (api.Identifier, error) {
	var result api.PostMessageReactionReturn
	if err := c.Query(
		ctx,
		api.PostMessageReaction,
		params,
		&result,
	); err != nil {
		return api.Identifier{}, err
	}
	return result.MessageReactionIdent, nil
}
