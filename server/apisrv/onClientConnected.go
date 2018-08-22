package apisrv

import (
	wwr "github.com/qbeon/webwire-go"
	"github.com/qbeon/webwire-example-postboard/server/apisrv/metrics"
)

// OnClientConnected implements the wwr.ServerImplementation interface
func (srv *apiServer) OnClientConnected(client wwr.Connection) {
	metrics.ConnectionCreated()
}
