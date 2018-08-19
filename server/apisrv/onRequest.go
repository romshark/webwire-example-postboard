package apisrv

import (
	"context"
	"time"

	wwr "github.com/qbeon/webwire-go"
	"github.com/qbeon/webwire-messenger/server/apisrv/api"
	"github.com/qbeon/webwire-messenger/server/apisrv/metrics"
	"github.com/qbeon/webwire-messenger/server/apisrv/sessinfo"
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
	var sessionInfo *sessinfo.SessionInfo
	s := connection.Session()
	if s != nil {
		sessionInfo = s.Info.(*sessinfo.SessionInfo)
	} else {
		sessionInfo = &sessinfo.SessionInfo{
			UserType: api.UtGuest,
		}
	}

	// Handle request
	return srv.resolver.Resolve(connection, sessionInfo, message)
}
