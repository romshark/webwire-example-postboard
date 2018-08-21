package setup

import (
	"context"
	"time"

	"github.com/qbeon/webwire-messenger/server/apisrv/api"
	"github.com/qbeon/webwire-messenger/server/client"
	"github.com/stretchr/testify/require"
)

// PostMessage posts a new message using the given client and parameters
// and verifies whether the message can be retrieved as well as whether
// the individual fields are correct
func (h *Helper) PostMessage(
	author client.ApiClient,
	params api.PostMessageParams,
) *api.Message {
	// Post message
	msgIdent, err := author.PostMessage(context.Background(), params)
	require.NoError(h.t, err)
	require.False(h.t, msgIdent.IsNull())

	// Retrieve message
	message, err := author.GetMessage(context.Background(), api.GetMessageParams{
		Ident: msgIdent,
	})
	require.NoError(h.t, err)
	require.NotNil(h.t, message)

	// Verify message fields
	require.Equal(h.t, msgIdent.String(), message.Identifier.String())
	require.Equal(h.t, params.Contents, message.Contents)
	require.WithinDuration(h.t,
		time.Now().UTC(),
		message.Publication,
		1*time.Second,
	)
	require.Equal(h.t, author.Identifier().String(), message.Author.String())
	require.Nil(h.t,
		message.LastEdit,
		"recently created message cannot have been edited",
	)
	require.Len(h.t,
		message.Reactions,
		0,
		"recently created message cannot have any reactions",
	)

	return message
}
