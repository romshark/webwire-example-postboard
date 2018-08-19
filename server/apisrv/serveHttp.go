package apisrv

import (
	"net/http"
)

// ServeHTTP implements the http.Handler interface.
func (srv *apiServer) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	if req.Method == "OPTIONS" {
		switch req.URL.Path {
		case "/":
			srv.wwrSrv.ServeHTTP(resp, req)
		case "/example":
			srv.exampleHTTPEndpoint.ServeHTTP(resp, req)
		default:
			srv.respondNotFound(resp)
		}
		return
	}

	switch req.URL.Path {
	// The main endpoint is occupied by the webwire server
	case "/":
		srv.wwrSrv.ServeHTTP(resp, req)
	// The following endpoint represents an example of an HTTP endpoint
	// used alongside the webwire endpoint
	case "/example":
		srv.exampleHTTPEndpoint.ServeHTTP(resp, req)
	default:
		srv.respondNotFound(resp)
	}
}

func (srv *apiServer) respondNotFound(resp http.ResponseWriter) {
	resp.WriteHeader(http.StatusNotFound)
	resp.Write([]byte("Resource not available"))
}
