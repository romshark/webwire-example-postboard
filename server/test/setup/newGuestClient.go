package setup

import (
	"github.com/qbeon/webwire-messenger/server/client"
	"github.com/stretchr/testify/require"
)

// NewGuestClient creates and connects a new guest client
func (ts *TestSetup) NewGuestClient() client.ApiClient {
	clt := ts.newClient()

	// Verify session absence
	require.Nil(ts.t, clt.Session())

	return clt
}
