package apisrv

import "net"

type ApiServer interface {
	// Launch initializes and launches the server
	Launch() error

	// AwaitShutdown blocks until the server is shut down
	AwaitShutdown() error

	// Address returns the address the server is listening on
	Address() net.Addr

	// MetricsAddress returns the address the metrics server is listening on
	MetricsAddress() net.Addr

	// Shutdown appoints a server shutdown and blocks until the server
	// is gracefully stopped
	Shutdown() error
}
