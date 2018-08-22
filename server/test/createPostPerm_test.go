package test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/qbeon/webwire-example-postboard/server/apisrv/api"
	"github.com/qbeon/webwire-example-postboard/server/client"
	"github.com/qbeon/webwire-example-postboard/server/test/setup"
)

// TestCreatePostPerm tests creating a post without proper permissions
func TestCreatePostPerm(t *testing.T) {
	tryCreating := func(
		t *testing.T,
		ts *setup.TestSetup,
		root,
		author client.ApiClient,
	) {
		// Try creating a new post
		ident, err := author.CreatePost(
			context.Background(),
			api.CreatePostParams{
				Contents: "this will not be created due to lack of permissions",
			},
		)
		ts.Helper.VerifyUnauthErr(err)
		require.True(t, ident.IsNull())

		// Ensure the post wasn't created
		posts, err := root.GetPosts(
			context.Background(),
			api.GetPostsParams{
				Limit: 10,
			},
		)
		require.NoError(t, err)
		require.Len(t, posts, 0)
	}

	t.Run("AsGuest", func(t *testing.T) {
		//t.Parallel()
		ts := setup.New(t, setupConf)
		defer ts.Teardown()

		root := ts.NewAdminClient("root", "root")
		guest := ts.NewGuestClient()

		tryCreating(
			t,
			ts,    // Test setup
			root,  // Root
			guest, // Author
		)
	})
}
