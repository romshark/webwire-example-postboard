package test

import (
	"context"
	"testing"

	"github.com/qbeon/webwire-messenger/server/apisrv/api"
	"github.com/qbeon/webwire-messenger/server/client"
	"github.com/qbeon/webwire-messenger/server/test/setup"
	"github.com/stretchr/testify/require"
)

// TestCreatePostReactionPerm tests creating post reactions
// without proper permission
func TestCreatePostReactionPerm(t *testing.T) {
	tryCreatePostReaction := func(
		t *testing.T,
		ts *setup.TestSetup,
		postAuthor,
		reactionAuthor client.ApiClient,
	) {
		// Post post
		post := ts.Helper.CreatePost(postAuthor, api.CreatePostParams{
			Contents: "now look at thaaat!",
		})

		// Try creating a post reaction
		ident, err := reactionAuthor.CreatePostReaction(
			context.Background(),
			api.CreatePostReactionParams{
				PostIdent:   post.Ident,
				Type:        api.Shock,
				Description: "this will not be created due to lack of perm.",
			},
		)
		ts.Helper.VerifyUnauthErr(err)
		require.True(t, ident.IsNull())

		// Ensure the reaction wasn't actually created
		reactionsAfter, err := postAuthor.GetReactionsOfPost(
			context.Background(),
			api.GetReactionsOfPostParams{
				PostIdent: post.Ident,
			},
		)
		require.NoError(t, err)
		require.Len(t, reactionsAfter, 0)
	}

	t.Run("AsGuest_ToRootPost", func(t *testing.T) {
		t.Parallel()
		ts := setup.New(t, setupConf)
		defer ts.Teardown()

		root := ts.NewAdminClient("root", "root")

		// Create random test users
		guest := ts.NewGuestClient()

		tryCreatePostReaction(
			t,
			ts,
			root,
			guest,
		)
	})

	t.Run("AsGuest_ToUserPost", func(t *testing.T) {
		t.Parallel()
		ts := setup.New(t, setupConf)
		defer ts.Teardown()

		root := ts.NewAdminClient("root", "root")

		// Create random test users
		guest := ts.NewGuestClient()
		_, _, user := ts.Helper.CreateUserRand(root, api.UtUser)

		tryCreatePostReaction(
			t,
			ts,
			user,
			guest,
		)
	})
}
