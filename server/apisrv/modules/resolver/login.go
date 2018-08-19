package resolver

import (
	"fmt"

	wwr "github.com/qbeon/webwire-go"
	"github.com/qbeon/webwire-messenger/server/apisrv/api"
	engiface "github.com/qbeon/webwire-messenger/server/apisrv/modules/engine"
	"github.com/qbeon/webwire-messenger/server/apisrv/sessinfo"
)

// Login implements the Resolver interface
func (rsv *resolver) Login(
	session *sessinfo.SessionInfo,
	connection wwr.Connection,
	params *api.LoginParams,
) (wwr.Payload, error) {
	// Verify the user isn't already logged in,
	// which is easily done by checking the current session's user type
	if session.UserType != api.UtGuest {
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
	user, actualPassword, err := rsv.engine.AuthenticateUser(params.Username)
	if err := rsv.handleError(err); err != nil {
		return nil, err
	}

	// Verify password
	if !rsv.passwordHasher.Compare(params.Password, actualPassword) {
		return nil, engiface.NewError(engiface.ErrWrongCredentials)
	}

	// Successfully create the session which will automatically be synchronized
	// to the API client
	if err := connection.CreateSession(&sessinfo.SessionInfo{
		UserIdentifier: user.Identifier,
		UserType:       user.Type,
	}); err != nil {
		return nil, rsv.logInternalError(err)
	}

	return nil, nil
}
