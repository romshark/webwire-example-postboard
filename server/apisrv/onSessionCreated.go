package apisrv

import (
	wwr "github.com/qbeon/webwire-go"
	"github.com/qbeon/webwire-messenger/server/apisrv/metrics"
)

// OnSessionCreated implements the wwr.SessionManager interface
func (srv *apiServer) OnSessionCreated(conn wwr.Connection) error {
	// Perform engine call
	err := srv.engine.SaveSession(conn.Session())
	if err != nil {
		return err
	}

	// Record session creation
	metrics.SessionCreated()

	return nil
}
