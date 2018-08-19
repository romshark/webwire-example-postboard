package resolver

import (
	"fmt"

	wwr "github.com/qbeon/webwire-go"
	engiface "github.com/qbeon/webwire-messenger/server/apisrv/engine"
)

// Login implements the Resolver interface
func (rsv *resolver) Login(
	session *engiface.SessionInfo,
	connection wwr.Connection,
	params *LoginParams,
) (wwr.Payload, error) {
	// Verify the user isn't already logged in,
	// which is easily done by checking the current session's user type
	if session.UserType != engiface.UtGuest {
		return nil, wwr.ReqErr{
			Code: "INVALID_REQ",
			Message: fmt.Sprintf(
				"user is already logged in as %s",
				session.UserType.String(),
			),
		}
	}

	// Perform an engine call to retrieve the given customers password
	// to verify it's validity
	user, err := rsv.engine.AuthenticateUser(params.Username)
	if err := rsv.handleError(err); err != nil {
		return nil, err
	}

	// Successfully create the session which will automatically be synchronized
	// to the API client
	if err := connection.CreateSession(&engiface.SessionInfo{
		UserIdentifier: user.Identifier,
		UserType:       user.Type,
	}); err != nil {
		return nil, rsv.logInternalError(err)
	}

	return nil, nil
}
