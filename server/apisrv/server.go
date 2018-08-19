package apisrv

import (
	"net"
	"net/http"
	"sync"

	"github.com/pkg/errors"
	wwr "github.com/qbeon/webwire-go"
	"github.com/qbeon/webwire-messenger/server/apisrv/config"
	"github.com/qbeon/webwire-messenger/server/apisrv/dam"
	engiface "github.com/qbeon/webwire-messenger/server/apisrv/modules/engine"
	"github.com/qbeon/webwire-messenger/server/apisrv/modules/logger"
	"github.com/qbeon/webwire-messenger/server/apisrv/modules/resolver"
)

type apiServer struct {
	// conf stores the initial server configuration
	conf config.Config

	// lock protects the API server internals from concurrent access
	lock *sync.RWMutex

	// addr represents the HTTP(S) server address
	addr net.Addr

	// metricsAddr represents the HTTP metrics server address
	metricsAddr net.Addr

	// stop represents a goroutine dam that's flushed
	// when the server is stopped.
	// It's used to make AwaitShutdown calls block
	stop dam.Dam

	// log represents the debug and error logger
	log logger.Logger

	// resolver represents the module resolving requests to engine calls
	resolver resolver.Resolver

	// engine represents the module implementing the core functionality
	// of the API server
	engine engiface.Engine

	// exampleHTTPEndpoint represents an example of an HTTP endpoint
	// used alongside the webwire websockets endpoint
	exampleHTTPEndpoint http.Handler

	// httpSrv represents the internal HTTP(S) server instance.
	// This instance will host the underlying webwire server
	httpSrv *http.Server

	// metricsSrv represents the internal HTTP metrics server instance.
	// It's used to export metrics to the outside
	metricsSrv *http.Server

	// wwrSrv represents the internal webwire server instance
	wwrSrv wwr.Server
}

// Address implements the Server interface
func (srv *apiServer) Address() net.Addr {
	srv.lock.RLock()
	addr := srv.addr
	srv.lock.RUnlock()
	return addr
}

// MetricsAddress implements the Server interface
func (srv *apiServer) MetricsAddress() net.Addr {
	srv.lock.RLock()
	addr := srv.metricsAddr
	srv.lock.RUnlock()
	return addr
}

// Launch implements the Server interface
func (srv *apiServer) Launch() error {
	metricsServerReady := make(chan error, 1)
	serverReady := make(chan error, 1)

	// Launch metrics server and wait for it
	go srv.listenMetrics(metricsServerReady)
	if err := <-metricsServerReady; err != nil {
		return errors.Wrap(err, "metrics server launch failed")
	}

	// Launch API server and wait for it
	go srv.listen(serverReady)
	if err := <-serverReady; err != nil {
		return errors.Wrap(err, "server launch failed")
	}

	return nil
}

// AwaitShutdown implements the Server interface
func (srv *apiServer) AwaitShutdown() error {
	return srv.stop.Wait()
}
