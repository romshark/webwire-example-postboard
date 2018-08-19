package apisrv

import (
	"context"
	"time"

	wwr "github.com/qbeon/webwire-go"
	engiface "github.com/qbeon/webwire-messenger/server/apisrv/engine"
	"github.com/qbeon/webwire-messenger/server/apisrv/metrics"
)

// OnRequest implements the wwr.ServerImplementation interface
func (srv *apiServer) OnRequest(
	_ context.Context,
	connection wwr.Connection,
	message wwr.Message,
) (response wwr.Payload, err error) {
	startTime := time.Now()
	defer func() {
		// Defer the record of a completed request until processing completion
		metrics.RequestCompleted(time.Since(startTime))
	}()

	// Record incoming request
	metrics.Request()

	// Extract session information
	var sessionInfo *engiface.SessionInfo
	s := connection.Session()
	if s != nil {
		sessionInfo = s.Info.(*engiface.SessionInfo)
	} else {
		sessionInfo = &engiface.SessionInfo{
			UserType: engiface.UtGuest,
		}
	}

	// Handle request
	return srv.resolver.Resolve(connection, sessionInfo, message)
}
