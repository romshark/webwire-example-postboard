package setup

import (
	"log"
	"testing"

	"github.com/qbeon/webwire-example-postboard/server/apisrv"
	"github.com/qbeon/webwire-example-postboard/server/apisrv/config"
)

// launchApiServer launches the API server instance in a separate goroutine
// blocking until the API server is initialized and ready
func launchApiServer(t *testing.T, conf config.Config) apisrv.ApiServer {
	apiServer, err := apisrv.NewApiServer(conf)
	if err != nil {
		log.Fatalf("API Server initialization failure: %s", err)
	}
	apiServer.Launch()

	return apiServer
}
