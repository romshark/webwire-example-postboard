package test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/qbeon/webwire-messenger/server/apisrv/api"
	"github.com/qbeon/webwire-messenger/server/client"
	"github.com/qbeon/webwire-messenger/server/test/setup"
)

// TestRemoveOwnMessages tests removing messages on behalf of their authors
func TestRemoveOwnMessages(t *testing.T) {
	postRemoveAndVerifyHistory := func(
		ts *setup.TestSetup,
		clt client.ApiClient,
		messageContents []string,
	) {
		postedMessages := ts.Helper.PostMessages(
			clt,
			[]string{
				"first message",
				"second message",
				"third message",
				"fourth message",
				"fifth message",
			},
			Readers{clt},
		)

		// Collect the identifiers of all posted messages
		identsToBeRemoved := make([]api.Identifier, len(postedMessages))
		for i, posted := range postedMessages {
			identsToBeRemoved[i] = posted.Identifier
		}

		ts.Helper.RemoveMessages(clt, identsToBeRemoved)

		// Verify removed messages don't appear in the history
		messagesFromHistory, err := clt.GetMessages(
			context.Background(),
			api.GetMessagesParams{
				Limit: 10,
			},
		)
		require.NoError(t, err)
		require.Len(t, messagesFromHistory, 0)
	}

	t.Run("AsRoot", func(t *testing.T) {
		t.Parallel()
		ts := setup.New(t, setupConf)
		defer ts.Teardown()

		root := ts.NewAdminClient("root", "root")

		postRemoveAndVerifyHistory(
			ts,
			root,
			[]string{
				"first message",
				"second message",
				"third message",
				"fourth message",
				"fifth message",
			},
		)
	})

	t.Run("AsUser", func(t *testing.T) {
		t.Parallel()
		ts := setup.New(t, setupConf)
		defer ts.Teardown()

		root := ts.NewAdminClient("root", "root")
		_, _, user := ts.Helper.CreateUserRand(root, api.UtUser)

		postRemoveAndVerifyHistory(
			ts,
			user,
			[]string{
				"first message",
				"second message",
				"third message",
				"fourth message",
				"fifth message",
			},
		)
	})
}
