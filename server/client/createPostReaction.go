package client

import (
	"context"

	"github.com/qbeon/webwire-messenger/server/apisrv/api"
)

// CreatePostReaction implements the ApiClient interface
func (c *apiClient) CreatePostReaction(
	ctx context.Context,
	params api.CreatePostReactionParams,
) (api.Identifier, error) {
	var result api.CreatePostReactionReturn
	if err := c.Query(
		ctx,
		api.CreatePostReaction,
		params,
		&result,
	); err != nil {
		return api.Identifier{}, err
	}
	return result.PostReactionIdent, nil
}
