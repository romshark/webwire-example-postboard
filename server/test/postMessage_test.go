package test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/qbeon/webwire-messenger/server/apisrv/api"
	"github.com/qbeon/webwire-messenger/server/client"
	"github.com/qbeon/webwire-messenger/server/test/setup"
)

// TestPostMessage tests posting a message and retrieving it
func TestPostMessage(t *testing.T) {
	t.Run("FromAdminToMultipleUsers", func(t *testing.T) {
		t.Parallel()
		ts := setup.New(t, stats)
		defer ts.Teardown()

		admin := ts.NewAdminClient("root", "root")

		// Create random test users
		_, _, user1 := ts.Helper.CreateUserRand(admin, api.UtUser)
		_, _, user2 := ts.Helper.CreateUserRand(admin, api.UtUser)

		// Post a new message from the root administrator
		posted := ts.Helper.PostMessage(admin, api.PostMessageParams{
			Contents: "test message",
		})

		// Ensure all users can read the message including the author himself
		// and the message is as expected
		readers := []client.ApiClient{
			admin, user1, user2,
		}
		for _, reader := range readers {
			messages, err := reader.GetMessages(
				context.Background(),
				api.GetMessagesParams{
					After: nil,
					Limit: 10,
				},
			)
			require.NoError(t, err)
			require.Len(t, messages, 1)

			// Verify the message
			message := messages[0]
			require.Equal(t, posted.Identifier, message.Identifier)
			require.Equal(t, *admin.Identifier(), message.Author)
			require.Equal(t, "test message", message.Contents)
			require.Nil(t, message.LastEdit)
			require.WithinDuration(t,
				time.Now().UTC(),
				message.Publication,
				3*time.Second,
			)
			require.Nil(t, message.Reactions)
		}

	})
}
