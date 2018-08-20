package client

import wwr "github.com/qbeon/webwire-go"

// OnSignal implements the wwrclt.Implementation interface
func (c *apiClient) OnSignal(message wwr.Payload) {
	if c.hooks.OnSignal != nil {
		c.hooks.OnSignal(message)
	}
}
