package apisrv

import "github.com/qbeon/webwire-messenger/server/apisrv/metrics"

// OnSessionClosed implements the wwr.SessionManager interface
func (srv *apiServer) OnSessionClosed(sessionKey string) error {
	// Perform engine call
	srv.engine.CloseSession(sessionKey)

	// Record session closure
	metrics.SessionClosed()

	return nil
}
