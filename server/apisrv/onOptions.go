package apisrv

import "net/http"

// OnOptions implements the wwr.ServerImplementation interface.
// Sets HTTP access control headers to satisfy CORS
func (srv *apiServer) OnOptions(resp http.ResponseWriter) {
	resp.Header().Set("Access-Control-Allow-Origin", "*")
	resp.Header().Set("Access-Control-Allow-Methods", "WEBWIRE")
}
