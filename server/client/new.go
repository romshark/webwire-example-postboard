package client

import (
	wwrclt "github.com/qbeon/webwire-go/client"
	"github.com/qbeon/webwire-messenger/server/apisrv/sessinfo"
)

// New constructs a new API client instance
func New(
	srvAddr string,
	opts wwrclt.Options,
	hooks Hooks,
) ApiClient {
	newClient := &apiClient{
		hooks: hooks,
	}

	// Initialize connection
	opts.SessionInfoParser = sessinfo.ParseSessionInfo
	newClient.connection = wwrclt.NewClient(srvAddr, newClient, opts)

	return newClient
}
