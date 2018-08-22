package setup

import (
	"context"
	"time"

	"github.com/qbeon/webwire-messenger/server/apisrv/api"
	"github.com/qbeon/webwire-messenger/server/client"
	"github.com/stretchr/testify/require"
)

// CreatePost creates a new post using the given client and parameters
// and verifies whether the post can be retrieved as well as whether
// the individual fields are correct
func (h *Helper) CreatePost(
	author client.ApiClient,
	params api.CreatePostParams,
) *api.Post {
	// Create the post
	msgIdent, err := author.CreatePost(context.Background(), params)
	require.NoError(h.t, err)
	require.False(h.t, msgIdent.IsNull())

	// Retrieve the post
	post, err := author.GetPost(context.Background(), api.GetPostParams{
		Ident: msgIdent,
	})
	require.NoError(h.t, err)
	require.NotNil(h.t, post)

	// Verify post fields
	require.Equal(h.t, msgIdent.String(), post.Ident.String())
	require.Equal(h.t, params.Contents, post.Contents)
	require.WithinDuration(h.t,
		time.Now().UTC(),
		post.Publication,
		h.ts.MaxCreationTimeDeviation(),
	)
	require.Equal(h.t, author.Identifier().String(), post.Author.String())
	require.Nil(h.t,
		post.LastEdit,
		"recently created post cannot have been edited",
	)

	return post
}
