package apisrv

import "net/http"

// BeforeUpgrade implements the wwr.ServerImplementation interface.
// It must return true to accept the incoming connection,
// returning false will reject it
func (srv *apiServer) BeforeUpgrade(
	resp http.ResponseWriter,
	req *http.Request,
) bool {
	// Accept client connection
	return true
}
