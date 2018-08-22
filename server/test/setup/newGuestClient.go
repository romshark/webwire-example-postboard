package setup

import (
	"github.com/qbeon/webwire-example-postboard/server/client"
	"github.com/stretchr/testify/require"
)

// NewGuestClient creates and connects a new guest client
// verifying whether the connection was successfully established
func (ts *TestSetup) NewGuestClient() client.ApiClient {
	clt := ts.newClient()

	// Verify session absence
	require.Nil(ts.t, clt.Session())

	return clt
}
