package client

import (
	wwr "github.com/qbeon/webwire-go"
	wwrclt "github.com/qbeon/webwire-go/client"
	"github.com/qbeon/webwire-messenger/server/apisrv/api"
)

// Hooks represents the hook functions used for hookable API client events
type Hooks struct {
	OnSessionCreated func(*wwr.Session)
	OnSessionClosed  func()
	OnDisconnected   func()
	OnSignal         func(wwr.Payload)
}

// apiClient implements the wwrclt.Implementation interface
type apiClient struct {
	connection wwrclt.Client
	hooks      Hooks
}

// Connect implements the ApiClient interface
func (c *apiClient) Connect() error {
	return c.connection.Connect()
}

// Close implements the ApiClient interface
func (c *apiClient) Close() {
	c.connection.Close()
}

// Session implements the ApiClient interface
func (c *apiClient) Session() *wwr.Session {
	return c.connection.Session()
}

// Identifier implements the ApiClient interface
func (c *apiClient) Identifier() *api.Identifier {
	val := c.connection.SessionInfo("id")
	if val == nil {
		return nil
	}
	ident := val.(api.Identifier)
	return &ident
}

// OnDisconnected implements the wwrclt.Implementation interface
func (c *apiClient) OnDisconnected() {
	if c.hooks.OnDisconnected != nil {
		c.hooks.OnDisconnected()
	}
}

// OnSessionClosed implements the wwrclt.Implementation interface
func (c *apiClient) OnSessionClosed() {
	if c.hooks.OnSessionClosed != nil {
		c.hooks.OnSessionClosed()
	}
}

// OnSessionCreated implements the wwrclt.Implementation interface
func (c *apiClient) OnSessionCreated(newSession *wwr.Session) {
	if c.hooks.OnSessionCreated != nil {
		c.hooks.OnSessionCreated(newSession)
	}
}
