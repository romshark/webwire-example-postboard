package resolver

import (
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
	wwr "github.com/qbeon/webwire-go"
	"github.com/qbeon/webwire-messenger/server/apisrv/api"
	"github.com/qbeon/webwire-messenger/server/apisrv/modules/authorizer"
	engiface "github.com/qbeon/webwire-messenger/server/apisrv/modules/engine"
	"github.com/qbeon/webwire-messenger/server/apisrv/modules/validator"
	"github.com/qbeon/webwire-messenger/server/apisrv/sessinfo"
)

// Resolve implements the Resolver interface
func (rsv *resolver) Resolve(
	connection wwr.Connection,
	sessionInfo *sessinfo.SessionInfo,
	message wwr.Message,
) (wwr.Payload, error) {
	var result interface{}
	var err error

	switch api.Method(message.Name()) {
	case api.Login:
		// Handle login request
		params := &api.LoginParams{}
		if err := parseParameters(message, params); err != nil {
			return nil, err
		}
		result, err = rsv.Login(sessionInfo, connection, params)

	case api.Logout:
		// Handle logout request
		result, err = rsv.Logout(sessionInfo, connection)

	case api.GetMessage:
		// Handle message request
		params := &api.GetMessageParams{}
		if err := parseParameters(message, params); err != nil {
			return nil, err
		}
		result, err = rsv.GetMessage(sessionInfo, params)

	case api.GetMessages:
		// Handle messages request
		params := &api.GetMessagesParams{}
		if err := parseParameters(message, params); err != nil {
			return nil, err
		}
		result, err = rsv.GetMessages(sessionInfo, params)

	case api.PostMessage:
		// Handle message creation request
		params := &api.PostMessageParams{}
		if err := parseParameters(message, params); err != nil {
			return nil, err
		}
		result, err = rsv.PostMessage(sessionInfo, params)

	case api.EditMessage:
		// Handle message edit request
		params := &api.EditMessageParams{}
		if err := parseParameters(message, params); err != nil {
			return nil, err
		}
		result, err = rsv.EditMessage(sessionInfo, params)

	case api.RemoveMessage:
		// Handle message deletion request
		params := &api.RemoveMessageParams{}
		if err := parseParameters(message, params); err != nil {
			return nil, err
		}
		result, err = rsv.RemoveMessage(sessionInfo, params)

	case api.PostMessageReaction:
		// Handle message reaction creation request
		params := &api.PostMessageReactionParams{}
		if err := parseParameters(message, params); err != nil {
			return nil, err
		}
		result, err = rsv.PostMessageReaction(sessionInfo, params)

	case api.CreateUser:
		// Handle user creation request
		params := &api.CreateUserParams{}
		if err := parseParameters(message, params); err != nil {
			return nil, err
		}
		result, err = rsv.CreateUser(sessionInfo, params)

	case api.GetUser:
		// Handle user profile request
		params := &api.GetUserParams{}
		if err := parseParameters(message, params); err != nil {
			return nil, err
		}
		result, err = rsv.GetUser(sessionInfo, params)

	default:
		// Fail requests asking for an invalid API action with a typed error
		return nil, wwr.ReqErr{
			Code: engiface.ErrInvalidRequest.String(),
			Message: fmt.Sprintf(
				"Unsupported request type: '%s'",
				message.Name(),
			),
		}
	}

	// Handle error if any
	if err != nil {
		if authorizer.IsAuthorizationError(err) {
			// Use unauthorized error code for authorization errors
			err = wwr.ReqErr{
				Code:    engiface.ErrUnauthorized.String(),
				Message: err.Error(),
			}
		} else if validator.IsValidatorError(err) {
			// Use generic invalid-parameter error code for validator errors
			err = wwr.ReqErr{
				Code:    engiface.ErrInvalidRequest.String(),
				Message: err.Error(),
			}
		}
		return nil, err
	}

	// If no results are returned then marshalling can be skipped
	if result == nil {
		return nil, nil
	}

	// Marshall result
	replyData, err := json.Marshal(result)
	if err != nil {
		return nil, rsv.logInternalError(
			errors.Wrap(err, "marshalling failed"),
		)
	}

	return wwr.NewPayload(wwr.EncodingUtf8, replyData), nil
}
