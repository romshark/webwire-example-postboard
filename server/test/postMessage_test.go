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
	testPostingAndReading := func(
		t *testing.T,
		ts *setup.TestSetup,
		author client.ApiClient,
		messageContents []string,
		readers []client.ApiClient,
	) {
		// Post new messages
		posted := make([]*api.Message, len(messageContents))
		for i, contents := range messageContents {
			posted[i] = ts.Helper.PostMessage(author, api.PostMessageParams{
				Contents: contents,
			})
		}

		// Ensure all readers get the correct messages
		for _, reader := range readers {
			messages, err := reader.GetMessages(
				context.Background(),
				api.GetMessagesParams{
					After: nil,
					Limit: 10,
				},
			)
			require.NoError(t, err)
			require.Len(t, messages, len(messageContents))

			// Verify the messages
			for i := range posted {
				retrievedMessage := messages[i]

				// Iterate through the posted messages in reversed order
				postedItr := len(posted) - 1 - i
				postedMessage := posted[postedItr]

				require.Equal(t,
					postedMessage.Identifier.String(),
					retrievedMessage.Identifier.String(),
				)
				require.Equal(t,
					*author.Identifier(),
					retrievedMessage.Author,
				)
				require.Equal(t,
					messageContents[postedItr],
					retrievedMessage.Contents,
				)
				require.Nil(t, postedMessage.LastEdit)
				require.Equal(t,
					postedMessage.Publication,
					retrievedMessage.Publication,
				)
				require.WithinDuration(t,
					time.Now().UTC(),
					retrievedMessage.Publication,
					3*time.Second,
				)
				require.Nil(t, retrievedMessage.Reactions)
			}
		}
	}

	t.Run("FromRoot", func(t *testing.T) {
		t.Parallel()
		ts := setup.New(t, stats)
		defer ts.Teardown()

		admin := ts.NewAdminClient("root", "root")

		// Create random test users
		_, _, user1 := ts.Helper.CreateUserRand(admin, api.UtUser)
		_, _, user2 := ts.Helper.CreateUserRand(admin, api.UtUser)
		guest := ts.NewGuestClient()

		testPostingAndReading(
			t,
			ts,    // Test setup
			admin, // Author
			[]string{"test message"}, // Message contents
			[]client.ApiClient{
				guest, admin, user1, user2,
			}, // Readers
		)
	})

	t.Run("FromUser", func(t *testing.T) {
		t.Parallel()
		ts := setup.New(t, stats)
		defer ts.Teardown()

		admin := ts.NewAdminClient("root", "root")

		// Create random test users
		_, _, user1 := ts.Helper.CreateUserRand(admin, api.UtUser)
		_, _, user2 := ts.Helper.CreateUserRand(admin, api.UtUser)
		guest := ts.NewGuestClient()

		// Ensure all users can read the message including the author himself
		// and the message is as expected
		testPostingAndReading(
			t,
			ts,    // Test setup
			user1, // Author
			[]string{"test message"}, // Message contents
			[]client.ApiClient{
				guest, admin, user1, user2,
			}, // Readers
		)
	})

	t.Run("MultipleFromRoot", func(t *testing.T) {
		t.Parallel()
		ts := setup.New(t, stats)
		defer ts.Teardown()

		admin := ts.NewAdminClient("root", "root")

		// Create random test users
		_, _, user1 := ts.Helper.CreateUserRand(admin, api.UtUser)
		_, _, user2 := ts.Helper.CreateUserRand(admin, api.UtUser)
		guest := ts.NewGuestClient()

		// Ensure all users can read the message including the author himself
		// and the message is as expected
		testPostingAndReading(
			t,
			ts,    // Test setup
			admin, // Author
			[]string{
				"first test message",
				"second test message",
				"third test message",
				"fourth test message",
				"fifth test message",
			}, // Message contents
			[]client.ApiClient{
				guest, admin, user1, user2,
			}, // Readers
		)
	})

	t.Run("MultipleFromUser", func(t *testing.T) {
		t.Parallel()
		ts := setup.New(t, stats)
		defer ts.Teardown()

		admin := ts.NewAdminClient("root", "root")

		// Create random test users
		_, _, user1 := ts.Helper.CreateUserRand(admin, api.UtUser)
		_, _, user2 := ts.Helper.CreateUserRand(admin, api.UtUser)
		guest := ts.NewGuestClient()

		// Ensure all users can read the message including the author himself
		// and the message is as expected
		testPostingAndReading(
			t,
			ts,    // Test setup
			user1, // Author
			[]string{
				"first test message",
				"second test message",
				"third test message",
				"fourth test message",
				"fifth test message",
			}, // Message contents
			[]client.ApiClient{
				guest, admin, user1, user2,
			}, // Readers
		)
	})
}
