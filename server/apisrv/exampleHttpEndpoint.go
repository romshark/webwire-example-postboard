package apisrv

import (
	"fmt"
	"net/http"
	"time"

	"github.com/qbeon/webwire-example-postboard/server/apisrv/metrics"
)

// ExampleHTTPEndpoint represents an HTTP handler example
// used alongside the webwire endpoint
type ExampleHTTPEndpoint struct {
	// apiServer refers to the API server hosting this handler
	apiServer ApiServer
}

// NewExampleHTTPEndpoint creates a new HTTP handler
func NewExampleHTTPEndpoint(apiServer ApiServer) http.Handler {
	return &ExampleHTTPEndpoint{
		apiServer: apiServer,
	}
}

// setCORSHeaders sets the CORS headers
func (hand *ExampleHTTPEndpoint) setCORSHeaders(resp http.ResponseWriter) {
	resp.Header().Set(
		"Access-Control-Allow-Headers",
		"Accept, Content-Type, Content-Length, "+
			"Accept-Encoding, Authorization,X-CSRF-Token",
	)
	resp.Header().Set("Access-Control-Expose-Headers", "Authorization")
	resp.Header().Set("Access-control-Allow-Origin", "*")
	resp.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
}

// ServeHTTP implements the http.Handler interface
func (hand *ExampleHTTPEndpoint) ServeHTTP(
	resp http.ResponseWriter,
	req *http.Request,
) {
	start := time.Now()
	hand.setCORSHeaders(resp)

	// Record a request
	metrics.Request()

	fmt.Fprint(resp, "I am an API server!")

	// Record request completion
	metrics.RequestCompleted(time.Since(start))

	hand.setCORSHeaders(resp)
	resp.Header().Set("Content-Type", "application/json")
}
