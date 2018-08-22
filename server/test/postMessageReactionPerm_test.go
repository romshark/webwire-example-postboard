package test

import (
	"context"
	"testing"

	"github.com/qbeon/webwire-messenger/server/apisrv/api"
	"github.com/qbeon/webwire-messenger/server/client"
	"github.com/qbeon/webwire-messenger/server/test/setup"
	"github.com/stretchr/testify/require"
)

// TestPostMessageReactionPerm tests posting message reactions
// without proper permission
func TestPostMessageReactionPerm(t *testing.T) {
	tryPostReaction := func(
		t *testing.T,
		ts *setup.TestSetup,
		messageAuthor,
		reactionAuthor client.ApiClient,
	) {
		// Post message
		messageBefore := ts.Helper.PostMessage(messageAuthor, api.PostMessageParams{
			Contents: "now look at thaaat!",
		})

		// Try posting message reaction
		ident, err := reactionAuthor.PostMessageReaction(
			context.Background(),
			api.PostMessageReactionParams{
				MessageIdent: messageBefore.Identifier,
				Type:         api.Shock,
				Description:  "this will not be posted due to lack of perm.",
			},
		)
		ts.Helper.VerifyUnauthErr(err)
		require.True(t, ident.IsNull())

		// Ensure the reaction wasn't actually posted
		messageAfter, err := messageAuthor.GetMessage(
			context.Background(),
			api.GetMessageParams{
				Ident: messageBefore.Identifier,
			},
		)
		require.NoError(t, err)
		require.Len(t, messageAfter.Reactions, 0)
	}

	t.Run("AsGuest_ToRootMessage", func(t *testing.T) {
		t.Parallel()
		ts := setup.New(t, setupConf)
		defer ts.Teardown()

		root := ts.NewAdminClient("root", "root")

		// Create random test users
		guest := ts.NewGuestClient()

		tryPostReaction(
			t,
			ts,
			root,
			guest,
		)
	})

	t.Run("AsGuest_ToUserMessage", func(t *testing.T) {
		t.Parallel()
		ts := setup.New(t, setupConf)
		defer ts.Teardown()

		root := ts.NewAdminClient("root", "root")

		// Create random test users
		guest := ts.NewGuestClient()
		_, _, user := ts.Helper.CreateUserRand(root, api.UtUser)

		tryPostReaction(
			t,
			ts,
			user,
			guest,
		)
	})
}
