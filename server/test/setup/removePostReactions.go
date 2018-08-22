package setup

import (
	"context"

	"github.com/qbeon/webwire-example-postboard/server/apisrv/api"
	"github.com/qbeon/webwire-example-postboard/server/client"
	"github.com/stretchr/testify/require"
)

// RemovePostReactions removes several post reactions using the given client
// and verifies whether the post reactions can't be retrieved individually
func (h *Helper) RemovePostReactions(
	clt client.ApiClient,
	postReactionsIdents []api.Identifier,
) {
	for _, ident := range postReactionsIdents {
		// Remove a post reaction
		require.NoError(h.t, clt.RemovePostReaction(
			context.Background(),
			api.RemovePostReactionParams{
				ReactionIdent: ident,
			},
		))

		// Ensure the post reaction is removed
		retrieved, err := clt.GetPostReaction(
			context.Background(),
			api.GetPostReactionParams{
				ReactionIdent: ident,
			},
		)
		require.NoError(h.t, err)
		require.Nil(h.t, retrieved)
	}
}
