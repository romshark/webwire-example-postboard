package test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/qbeon/webwire-example-postboard/server/apisrv/api"
	"github.com/qbeon/webwire-example-postboard/server/client"
	"github.com/qbeon/webwire-example-postboard/server/test/setup"
)

// TestRemoveOwnPost tests removing posts on behalf of their authors
func TestRemoveOwnPost(t *testing.T) {
	createRemoveAndVerifyHistory := func(
		ts *setup.TestSetup,
		clt client.ApiClient,
		postContents []string,
	) {
		created := ts.Helper.CreatePosts(
			clt,
			[]string{
				"first post",
				"second post",
				"third post",
				"fourth post",
				"fifth post",
			},
			Readers{clt},
		)

		// Collect the identifiers of all created posts
		identsToBeRemoved := make([]api.Identifier, len(created))
		for i, post := range created {
			identsToBeRemoved[i] = post.Ident
		}

		ts.Helper.RemovePosts(clt, identsToBeRemoved)

		// Verify removed posts don't appear in the history
		postsFromHistory, err := clt.GetPosts(
			context.Background(),
			api.GetPostsParams{
				Limit: 10,
			},
		)
		require.NoError(t, err)
		require.Len(t, postsFromHistory, 0)
	}

	t.Run("AsRoot", func(t *testing.T) {
		t.Parallel()
		ts := setup.New(t, setupConf)
		defer ts.Teardown()

		root := ts.NewAdminClient("root", "root")

		createRemoveAndVerifyHistory(
			ts,
			root,
			[]string{
				"first post",
				"second post",
				"third post",
				"fourth post",
				"fifth post",
			},
		)
	})

	t.Run("AsUser", func(t *testing.T) {
		t.Parallel()
		ts := setup.New(t, setupConf)
		defer ts.Teardown()

		root := ts.NewAdminClient("root", "root")
		_, _, user := ts.Helper.CreateUserRand(root, api.UtUser)

		createRemoveAndVerifyHistory(
			ts,
			user,
			[]string{
				"first post",
				"second post",
				"third post",
				"fourth post",
				"fifth post",
			},
		)
	})
}
