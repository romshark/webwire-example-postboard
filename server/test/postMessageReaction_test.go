package test

import (
	"testing"

	"github.com/qbeon/webwire-messenger/server/apisrv/api"
	"github.com/qbeon/webwire-messenger/server/test/setup"
)

// TestPostMessageReaction tests posting message reactions
func TestPostMessageReaction(t *testing.T) {

	t.Run("AsRoot_ToUsersMessage", func(t *testing.T) {
		t.Parallel()
		ts := setup.New(t, setupConf)
		defer ts.Teardown()

		root := ts.NewAdminClient("root", "root")

		// Create random test users
		_, _, user1 := ts.Helper.CreateUserRand(root, api.UtUser)

		message := ts.Helper.PostMessage(user1, api.PostMessageParams{
			Contents: "now look at thaaat!",
		})

		ts.Helper.PostMessageReaction(
			user1,
			root,
			message.Identifier,
			api.Shock,
			"wooow!",
		)
	})

	t.Run("AsUser_ToRootMessage", func(t *testing.T) {
		t.Parallel()
		ts := setup.New(t, setupConf)
		defer ts.Teardown()

		root := ts.NewAdminClient("root", "root")

		// Create random test users
		_, _, user1 := ts.Helper.CreateUserRand(root, api.UtUser)

		message := ts.Helper.PostMessage(root, api.PostMessageParams{
			Contents: "we did it!",
		})

		ts.Helper.PostMessageReaction(
			root,
			user1,
			message.Identifier,
			api.Shock,
			"Huraah!",
		)
	})
}
