package resolver

import (
	wwr "github.com/qbeon/webwire-go"
	engiface "github.com/qbeon/webwire-messenger/server/apisrv/engine"
)

// Logout implements the Resolver interface
func (rsv *resolver) Logout(
	session *engiface.SessionInfo,
	connection wwr.Connection,
) (wwr.Payload, error) {
	// Ensure the user is not already an unauthenticated guest,
	// which is easily done by checking the current session's user type
	if session.UserType == engiface.UtGuest {
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
