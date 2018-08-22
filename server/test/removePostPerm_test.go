package test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/qbeon/webwire-example-postboard/server/apisrv/api"
	"github.com/qbeon/webwire-example-postboard/server/client"
	"github.com/qbeon/webwire-example-postboard/server/test/setup"
)

// TestRemovePostPerm tests removing a post without proper permissions
func TestRemovePostPerm(t *testing.T) {
	tryRemove := func(
		ts *setup.TestSetup,
		postAuthor,
		remover client.ApiClient,
	) {
		created := ts.Helper.CreatePost(
			postAuthor,
			api.CreatePostParams{
				Contents: "first post",
			},
		)

		// Try to remove the post without having proper permissions
		err := remover.RemovePost(
			context.Background(),
			api.RemovePostParams{
				PostIdent: created.Ident,
			},
		)
		ts.Helper.VerifyUnauthErr(err)

		// Verify the post wasn't removed
		post, err := postAuthor.GetPost(
			context.Background(),
			api.GetPostParams{
				Ident: created.Ident,
			},
		)
		require.NoError(t, err)
		require.NotNil(t, post)
	}

	t.Run("AsGuest_PostFromRoot", func(t *testing.T) {
		//t.Parallel()
		ts := setup.New(t, setupConf)
		defer ts.Teardown()

		root := ts.NewAdminClient("root", "root")
		guest := ts.NewGuestClient()

		tryRemove(
			ts,
			root,  // Post author
			guest, // Remover
		)
	})

	t.Run("AsGuest_PostFromUser", func(t *testing.T) {
		//t.Parallel()
		ts := setup.New(t, setupConf)
		defer ts.Teardown()

		root := ts.NewAdminClient("root", "root")
		_, _, user := ts.Helper.CreateUserRand(root, api.UtUser)
		guest := ts.NewGuestClient()

		tryRemove(
			ts,
			user,  // Post author
			guest, // Remover
		)
	})

	t.Run("AsUser_PostFromRoot", func(t *testing.T) {
		//t.Parallel()
		ts := setup.New(t, setupConf)
		defer ts.Teardown()

		root := ts.NewAdminClient("root", "root")
		_, _, userAuthor := ts.Helper.CreateUserRand(root, api.UtUser)
		_, _, userRemover := ts.Helper.CreateUserRand(root, api.UtUser)

		tryRemove(
			ts,
			userAuthor,  // Post author
			userRemover, // Remover
		)
	})
}
