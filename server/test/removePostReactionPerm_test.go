package test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/qbeon/webwire-example-postboard/server/apisrv/api"
	"github.com/qbeon/webwire-example-postboard/server/client"
	"github.com/qbeon/webwire-example-postboard/server/test/setup"
)

// TestRemovePostReactionPerm tests removing a post reaction
// without proper permissions
func TestRemovePostReactionPerm(t *testing.T) {
	createAndTryRemove := func(
		ts *setup.TestSetup,
		postAuthor,
		reactionAuthor,
		remover client.ApiClient,
	) {
		// Create a post to create the reaction on
		created := ts.Helper.CreatePost(
			postAuthor,
			api.CreatePostParams{
				Contents: "first post",
			},
		)

		// Create a reaction
		createdReaction := ts.Helper.CreatePostReaction(
			postAuthor,
			reactionAuthor,
			created.Ident,
			api.Celebration,
			"sample reaction",
		)

		// Try to remove the post reaction without having proper permissions
		err := remover.RemovePost(
			context.Background(),
			api.RemovePostParams{
				PostIdent: created.Ident,
			},
		)
		ts.Helper.VerifyUnauthErr(err)

		// Verify the post reaction wasn't removed
		postReaction, err := postAuthor.GetPostReaction(
			context.Background(),
			api.GetPostReactionParams{
				ReactionIdent: createdReaction.Ident,
			},
		)
		require.NoError(t, err)
		require.NotNil(t, postReaction)
	}

	t.Run("AsGuest_FromUser", func(t *testing.T) {
		//t.Parallel()
		ts := setup.New(t, setupConf)
		defer ts.Teardown()

		root := ts.NewAdminClient("root", "root")
		_, _, user := ts.Helper.CreateUserRand(root, api.UtUser)
		guest := ts.NewGuestClient()

		createAndTryRemove(
			ts,
			root,  // Post author
			user,  // Reaction author
			guest, // Remover
		)
	})

	t.Run("AsUser_FromOtherUser", func(t *testing.T) {
		//t.Parallel()
		ts := setup.New(t, setupConf)
		defer ts.Teardown()

		root := ts.NewAdminClient("root", "root")
		_, _, userAuthor := ts.Helper.CreateUserRand(root, api.UtUser)
		_, _, userRemover := ts.Helper.CreateUserRand(root, api.UtUser)

		createAndTryRemove(
			ts,
			root,        // Post author
			userAuthor,  // Reaction author
			userRemover, // Remover
		)
	})
}
