package test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/qbeon/webwire-messenger/server/apisrv/api"
	"github.com/qbeon/webwire-messenger/server/client"
	"github.com/qbeon/webwire-messenger/server/test/setup"
)

// TestRemoveMessagePerm tests removing messages without proper permissions
func TestRemoveMessagePerm(t *testing.T) {
	tryRemove := func(
		ts *setup.TestSetup,
		messageAuthor,
		remover client.ApiClient,
	) {
		postedMessage := ts.Helper.PostMessage(
			messageAuthor,
			api.PostMessageParams{
				Contents: "first message",
			},
		)

		// Try to remove the message without having proper permissions
		err := remover.RemoveMessage(
			context.Background(),
			api.RemoveMessageParams{
				MessageIdent: postedMessage.Identifier,
			},
		)
		ts.Helper.VerifyUnauthErr(err)

		// Verify the message wasn't removed
		message, err := messageAuthor.GetMessage(
			context.Background(),
			api.GetMessageParams{
				Ident: postedMessage.Identifier,
			},
		)
		require.NoError(t, err)
		require.NotNil(t, message)
	}

	t.Run("AsGuest_MessageFromRoot", func(t *testing.T) {
		t.Parallel()
		ts := setup.New(t, setupConf)
		defer ts.Teardown()

		root := ts.NewAdminClient("root", "root")
		guest := ts.NewGuestClient()

		tryRemove(
			ts,
			root,  // Message author
			guest, // Remover
		)
	})

	t.Run("AsGuest_MessageFromUser", func(t *testing.T) {
		t.Parallel()
		ts := setup.New(t, setupConf)
		defer ts.Teardown()

		root := ts.NewAdminClient("root", "root")
		_, _, user := ts.Helper.CreateUserRand(root, api.UtUser)
		guest := ts.NewGuestClient()

		tryRemove(
			ts,
			user,  // Message author
			guest, // Remover
		)
	})

	t.Run("AsUser_MessageFromRoot", func(t *testing.T) {
		t.Parallel()
		ts := setup.New(t, setupConf)
		defer ts.Teardown()

		root := ts.NewAdminClient("root", "root")
		_, _, userAuthor := ts.Helper.CreateUserRand(root, api.UtUser)
		_, _, userRemover := ts.Helper.CreateUserRand(root, api.UtUser)

		tryRemove(
			ts,
			userAuthor,  // Message author
			userRemover, // Remover
		)
	})
}
