package setup

import (
	"context"

	"github.com/qbeon/webwire-messenger/server/apisrv/api"
	"github.com/qbeon/webwire-messenger/server/client"
	"github.com/stretchr/testify/require"
)

// PostMessage posts a new message using the given client and parameters
func (h *Helper) PostMessage(
	clt client.ApiClient,
	params api.PostMessageParams,
) *api.Message {
	// Post message
	msgIdent, err := clt.PostMessage(context.Background(), params)
	require.NoError(h.t, err)

	// Retrieve message
	message, err := clt.GetMessage(context.Background(), api.GetMessageParams{
		Ident: msgIdent,
	})
	require.NoError(h.t, err)
	require.NotNil(h.t, message)

	// Verify identifier
	require.Equal(h.t, msgIdent.String(), message.Identifier.String())

	return message
}
