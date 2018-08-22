package setup

import (
	"context"

	"github.com/qbeon/webwire-example-postboard/server/apisrv/api"
	"github.com/qbeon/webwire-example-postboard/server/client"
	"github.com/stretchr/testify/require"
)

// RemovePosts removes several posts using the given client
// and verifies whether the posts can't be retrieved individually
func (h *Helper) RemovePosts(
	clt client.ApiClient,
	postsIdents []api.Identifier,
) {
	for _, ident := range postsIdents {
		require.NoError(h.t, clt.RemovePost(
			context.Background(),
			api.RemovePostParams{
				PostIdent: ident,
			},
		))
		retrieved, err := clt.GetPost(
			context.Background(),
			api.GetPostParams{
				Ident: ident,
			},
		)
		require.NoError(h.t, err)
		require.Nil(h.t, retrieved)
	}
}
