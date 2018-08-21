package setup

import (
	"context"
	"time"

	"github.com/qbeon/webwire-messenger/server/apisrv/api"
	"github.com/qbeon/webwire-messenger/server/client"
	"github.com/stretchr/testify/require"
)

// PostMessages posts several new messages using the given client and contents
// and verifies whether the messages are correct and can be read by the given
// reader clients
func (h *Helper) PostMessages(
	author client.ApiClient,
	messageContents []string,
	readers []client.ApiClient,
) []*api.Message {
	postedMessages := make([]*api.Message, len(messageContents))
	for i, contents := range messageContents {
		postedMessages[i] = h.PostMessage(author, api.PostMessageParams{
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
		require.NoError(h.t, err)
		require.Len(h.t, messages, len(postedMessages))

		// Verify the messages
		for i := range postedMessages {
			retrievedMessage := messages[i]

			// Iterate through the posted messages in reversed order
			postedItr := len(postedMessages) - 1 - i
			postedMessage := postedMessages[postedItr]

			require.Equal(h.t,
				postedMessage.Identifier.String(),
				retrievedMessage.Identifier.String(),
			)
			require.Equal(h.t,
				*author.Identifier(),
				retrievedMessage.Author,
			)
			require.Equal(h.t,
				messageContents[postedItr],
				retrievedMessage.Contents,
			)
			require.Nil(h.t, postedMessage.LastEdit)
			require.Equal(h.t,
				postedMessage.Publication,
				retrievedMessage.Publication,
			)
			require.WithinDuration(h.t,
				time.Now().UTC(),
				retrievedMessage.Publication,
				3*time.Second,
			)
			require.Nil(h.t, retrievedMessage.Reactions)
		}
	}

	return postedMessages
}
