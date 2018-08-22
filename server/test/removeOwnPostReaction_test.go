package test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/qbeon/webwire-messenger/server/apisrv/api"
	"github.com/qbeon/webwire-messenger/server/client"
	"github.com/qbeon/webwire-messenger/server/test/setup"
)

// TestRemoveOwnPostReaction tests removing post reactions
// on behalf of their authors
func TestRemoveOwnPostReaction(t *testing.T) {
	createRemoveAndVerifyHistory := func(
		ts *setup.TestSetup,
		postAuthor,
		reactionAuthor client.ApiClient,
		reactionType api.PostReactionType,
		description string,
	) {
		post := ts.Helper.CreatePost(postAuthor, api.CreatePostParams{
			Contents: "sample post",
		})

		created := ts.Helper.CreatePostReaction(
			postAuthor,
			reactionAuthor,
			post.Ident,
			reactionType,
			description,
		)

		// Remove reaction
		ts.Helper.RemovePostReactions(
			reactionAuthor,
			[]api.Identifier{created.Ident},
		)

		// Verify removed post reactions
		retrieved, err := reactionAuthor.GetPostReaction(
			context.Background(),
			api.GetPostReactionParams{
				ReactionIdent: created.Ident,
			},
		)
		require.NoError(t, err)
		require.Nil(t, retrieved)
	}

	t.Run("AsRoot", func(t *testing.T) {
		t.Parallel()
		ts := setup.New(t, setupConf)
		defer ts.Teardown()

		root := ts.NewAdminClient("root", "root")
		_, _, user := ts.Helper.CreateUserRand(root, api.UtUser)

		createRemoveAndVerifyHistory(
			ts,
			user,            // Post author
			root,            // Reaction author
			api.Celebration, // Reaction type
			"hurah!",        // Description
		)
	})

	t.Run("AsAdmin", func(t *testing.T) {
		t.Parallel()
		ts := setup.New(t, setupConf)
		defer ts.Teardown()

		root := ts.NewAdminClient("root", "root")
		_, _, admin := ts.Helper.CreateUserRand(root, api.UtAdmin)

		createRemoveAndVerifyHistory(
			ts,
			root,            // Post author
			admin,           // Reaction author
			api.Celebration, // Reaction type
			"hurah!",        // Description
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
			root,            // Post author
			user,            // Reaction author
			api.Celebration, // Reaction type
			"hurah!",        // Description
		)
	})
}
