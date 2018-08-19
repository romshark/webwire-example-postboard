package apisrv

import (
	"context"

	"github.com/pkg/errors"
)

// shutdown initiates the server shutdown
func (srv *apiServer) shutdown(shutdownReasonErr error) error {
	if shutdownReasonErr != nil {
		srv.log.Printf(
			"attempting server shutdown due to runtime error: %s",
			shutdownReasonErr,
		)
	} else {
		srv.log.Print("attempting server shutdown")
	}

	defer func() {
		// Finalize the shutdown process
		srv.log.Print("server stopped")

		// Close loggers at the very end of the shutdown process
		srv.log.Close()

		// Notify the initiator about the shutdown process completion
		srv.stop.Flush(shutdownReasonErr)
	}()

	// Shutdown main HTTP server
	if err := srv.httpSrv.Shutdown(context.Background()); err != nil {
		return errors.Wrap(err, "http server shutdown failed")
	}

	// Shutdown metrics server
	if err := srv.metricsSrv.Shutdown(context.Background()); err != nil {
		return errors.Wrap(err, "metrics server shutdown failed")
	}

	return shutdownReasonErr
}

// Shutdown implements the apiServer interface
func (srv *apiServer) Shutdown() error {
	return srv.shutdown(nil)
}
