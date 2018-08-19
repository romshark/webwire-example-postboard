package apisrv

import (
	"context"

	wwr "github.com/qbeon/webwire-go"
)

// OnSignal implements the wwr.ServerImplementation interface
func (srv *apiServer) OnSignal(
	_ context.Context,
	_ wwr.Connection,
	_ wwr.Message,
) {
	// Ignore incoming signals
}
