package client

import (
	"context"

	wwr "github.com/qbeon/webwire-go"
	"github.com/qbeon/webwire-messenger/server/apisrv/api"
)

// ApiClient represents an API client
type ApiClient interface {
	// Include the API interface definition
	api.Api

	// Connect wraps wwrclt.Client.Connect
	Connect() error

	// Close wraps wwrclt.Client.Close
	Close()

	// Session wraps wwrclt.Client.Session
	Session() *wwr.Session

	// Identifier returns the identifier of the client.
	// Returns nil if this client is an unauthenticated guest client
	Identifier() *api.Identifier

	Query(
		ctx context.Context,
		method api.Method,
		params interface{},
		result interface{},
	) error
}
