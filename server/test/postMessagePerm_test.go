package test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/qbeon/webwire-messenger/server/apisrv/api"
	"github.com/qbeon/webwire-messenger/server/client"
	"github.com/qbeon/webwire-messenger/server/test/setup"
)

// TestPostMessagePerm tests posting a message without proper permissions
func TestPostMessagePerm(t *testing.T) {
	tryPosting := func(
		t *testing.T,
		ts *setup.TestSetup,
		root,
		author client.ApiClient,
	) {
		// Try posting a new messages
		ident, err := author.PostMessage(
			context.Background(),
			api.PostMessageParams{
				Contents: "this will not be posted due to lack of permissions",
			},
		)
		ts.Helper.VerifyUnauthErr(err)
		require.True(t, ident.IsNull())

		// Ensure the message wasn't posted
		messages, err := root.GetMessages(
			context.Background(),
			api.GetMessagesParams{
				Limit: 10,
			},
		)
		require.NoError(t, err)
		require.Len(t, messages, 0)
	}

	t.Run("AsGuest", func(t *testing.T) {
		t.Parallel()
		ts := setup.New(t, setupConf)
		defer ts.Teardown()

		root := ts.NewAdminClient("root", "root")
		guest := ts.NewGuestClient()

		tryPosting(
			t,
			ts,    // Test setup
			root,  // Root
			guest, // Author
		)
	})
}
