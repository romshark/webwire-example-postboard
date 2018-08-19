package apisrv

import (
	"net"
	"net/http"
	"time"

	"github.com/pkg/errors"
)

// tcpKeepAliveListener resembles the standard HTTP TCP Keep-Alive listener.
// It allows the server to pick any system provided port
// when no port is explicitly specified
type tcpKeepAliveListener struct {
	*net.TCPListener
}

// Accept accepts incoming TCP connections setting the keep-alive period
func (ln tcpKeepAliveListener) Accept() (net.Conn, error) {
	tc, err := ln.AcceptTCP()
	if err != nil {
		return nil, err
	}
	tc.SetKeepAlive(true)
	tc.SetKeepAlivePeriod(3 * time.Minute)
	return tc, nil
}

// listenMetrics launches the local metrics export HTTP server
// and triggers the ready channel with nil when the listener is operational.
// And error is passed to the ready channel in case of a failed launch
func (srv *apiServer) listenMetrics(ready chan error) {
	// Set address if not yet set
	srv.lock.Lock()
	addr := srv.conf.MetricsServerAddress
	srv.lock.Unlock()

	if addr == "" {
		addr = ":http"
	}

	// Setup TCP listener
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		// Fail launch and shutdown the server
		// if the TCP listener failed to set up properly
		err = errors.Wrap(err, "TCP listener failed")
		ready <- err
		srv.shutdown(err)
		return
	}
	defer listener.Close()

	// Remember listener address
	srv.lock.Lock()
	srv.metricsAddr = listener.Addr()
	addr = srv.metricsAddr.String()
	srv.lock.Unlock()

	// Serve
	srv.log.Printf("metrics server listening on http://%s", addr)
	ready <- nil

	// Serve HTTP
	if err := srv.metricsSrv.Serve(
		tcpKeepAliveListener{
			listener.(*net.TCPListener),
		},
	); err != http.ErrServerClosed {
		srv.shutdown(errors.Wrap(
			err,
			"metrics server failed",
		))
	}
}

// listen launches the main HTTP(S) server
// and triggers the ready channel with nil when the listener is operational.
// And error is passed to the ready channel in case of a failed launch
func (srv *apiServer) listen(ready chan error) {
	// Set address if not yet set
	srv.lock.Lock()
	addr := srv.conf.ServerAddress
	srv.lock.Unlock()

	if addr == "" {
		addr = ":http"
	}

	// Setup TCP listener
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		// Fail launch and shutdown the server
		// if the TCP listener failed to set up properly
		err = errors.Wrap(err, "TCP Listener failed: %s")
		ready <- err
		srv.shutdown(err)
		return
	}
	defer listener.Close()

	// Remember listener address
	srv.lock.Lock()
	srv.addr = listener.Addr()
	addr = srv.addr.String()
	srv.lock.Unlock()

	// Serve
	if srv.conf.TLS != nil {
		srv.log.Printf("listening on https://%s", addr)
		ready <- nil

		// Serve HTTPS
		if err := srv.httpSrv.ServeTLS(
			tcpKeepAliveListener{
				listener.(*net.TCPListener),
			},
			srv.conf.TLS.FullChainCertFilePath,
			srv.conf.TLS.KeyFilePath,
		); err != http.ErrServerClosed {
			srv.shutdown(errors.Wrap(err, "HTTPS server failed"))
		}
	} else {
		srv.log.Printf("listening on http://%s", addr)
		ready <- nil

		// Serve HTTP
		if err := srv.httpSrv.Serve(
			tcpKeepAliveListener{
				listener.(*net.TCPListener),
			},
		); err != http.ErrServerClosed {
			srv.shutdown(errors.Wrap(err, "HTTP server failed"))
		}
	}
}
