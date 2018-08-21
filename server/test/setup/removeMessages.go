package setup

import (
	"context"

	"github.com/qbeon/webwire-messenger/server/apisrv/api"
	"github.com/qbeon/webwire-messenger/server/client"
	"github.com/stretchr/testify/require"
)

// RemoveMessages revemos several messages using the given client
// and verifies whether the messages can't be retrieved individually
func (h *Helper) RemoveMessages(
	clt client.ApiClient,
	messagesIdents []api.Identifier,
) {
	for _, ident := range messagesIdents {
		require.NoError(h.t, clt.RemoveMessage(
			context.Background(),
			api.RemoveMessageParams{
				MessageIdent: ident,
			},
		))
		retrieved, err := clt.GetMessage(
			context.Background(),
			api.GetMessageParams{
				Ident: ident,
			},
		)
		require.NoError(h.t, err)
		require.Nil(h.t, retrieved)
	}
}
