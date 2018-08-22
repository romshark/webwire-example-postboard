package resolver

import (
	wwr "github.com/qbeon/webwire-go"
	"github.com/qbeon/webwire-example-postboard/server/apisrv/api"
	"github.com/qbeon/webwire-example-postboard/server/apisrv/sessinfo"
)

// Logout implements the Resolver interface
func (rsv *resolver) Logout(
	session *sessinfo.SessionInfo,
	connection wwr.Connection,
) (interface{}, error) {
	// Ensure the user is not already an unauthenticated guest,
	// which is easily done by checking the current session's user type
	if session.UserType == api.UtGuest {
		return nil, wwr.ReqErr{
			Code:    "INVALID_REQ",
			Message: "user is not authenticated",
		}
	}

	// Drop the session for this connection
	if err := connection.CloseSession(); err != nil {
		return nil, rsv.logInternalError(err)
	}

	return nil, nil
}
