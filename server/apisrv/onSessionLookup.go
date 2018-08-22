package apisrv

import (
	wwr "github.com/qbeon/webwire-go"
	"github.com/qbeon/webwire-example-postboard/server/apisrv/metrics"
)

// OnSessionLookup implements the wwr.SessionManager interface
func (srv *apiServer) OnSessionLookup(key string) (
	wwr.SessionLookupResult,
	error,
) {
	// Perform engine call
	result, err := srv.engine.FindSession(key)
	if err != nil {
		return wwr.SessionLookupResult{}, err
	}

	// Record session lookup
	metrics.SessionLookup()
	return result, err
}
