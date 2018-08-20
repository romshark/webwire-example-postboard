package main

import (
	"flag"
	"log"

	"github.com/qbeon/webwire-messenger/server/apisrv"
	"github.com/qbeon/webwire-messenger/server/apisrv/config"
)

var addr = flag.String(
	"addr",
	"127.0.0.1:",
	"Server address (port is chosen automatically if not specified)",
)
var metricsAddr = flag.String(
	"metaddr",
	"127.0.0.1:6000",
	"Metrics server address (port is chosen automatically if not specified)",
)

func main() {
	// Parse command line arguments
	flag.Parse()

	// NOTE: load server configuration and do everything you need
	// before the server is initialized and launched

	// Create a new API server instance
	srv, err := apisrv.NewApiServer(config.Config{
		ServerAddress:        *addr,
		MetricsServerAddress: *metricsAddr,
	})
	if err != nil {
		log.Fatalf("server initialization failed: %s", err)
	}

	// Launch the server
	log.Print("launching server...")
	if err := srv.Launch(); err != nil {
		log.Fatalf("server launch failed: %s", err)
	}
	log.Printf("server listening on: %s", srv.Address())

	// Setup termination signal listener to gracefully shutdown the server
	// in case of an OS termination signal
	onProcessInterrupt(func() {
		srv.Shutdown()
	})

	// Block main goroutine until the server is shut down
	if err := srv.AwaitShutdown(); err != nil {
		log.Fatalf("server shutdown failure: %s", err)
	}
	log.Print("server stopped")
}
