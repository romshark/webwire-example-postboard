package setup

import (
	"context"
	"time"

	"github.com/qbeon/webwire-messenger/server/apisrv/api"
	"github.com/qbeon/webwire-messenger/server/client"
	"github.com/stretchr/testify/require"
)

// CreatePosts creates several new posts using the given client and contents
// and verifies whether the posts are correct and can be read by the given
// reader clients
func (h *Helper) CreatePosts(
	author client.ApiClient,
	postContents []string,
	readers []client.ApiClient,
) []*api.Post {
	created := make([]*api.Post, len(postContents))
	for i, contents := range postContents {
		created[i] = h.CreatePost(author, api.CreatePostParams{
			Contents: contents,
		})
	}

	// Ensure all readers get the correct posts
	for _, reader := range readers {
		posts, err := reader.GetPosts(
			context.Background(),
			api.GetPostsParams{
				After: nil,
				Limit: 10,
			},
		)
		require.NoError(h.t, err)
		require.Len(h.t, posts, len(created))

		// Verify the posts
		for i := range created {
			retrievedPost := posts[i]

			// Iterate through the created posts in reversed order
			createdItr := len(created) - 1 - i
			createdPost := created[createdItr]

			require.Equal(h.t,
				createdPost.Ident.String(),
				retrievedPost.Ident.String(),
			)
			require.Equal(h.t,
				*author.Identifier(),
				retrievedPost.Author,
			)
			require.Equal(h.t,
				postContents[createdItr],
				retrievedPost.Contents,
			)
			require.Nil(h.t, createdPost.LastEdit)
			require.Equal(h.t,
				createdPost.Publication,
				retrievedPost.Publication,
			)
			require.WithinDuration(h.t,
				time.Now().UTC(),
				retrievedPost.Publication,
				h.ts.MaxCreationTimeDeviation(),
			)
		}
	}

	return created
}
