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

	case api.GetPost:
		// Handle post request
		params := &api.GetPostParams{}
		if err := parseParameters(message, params); err != nil {
			return nil, err
		}
		result, err = rsv.GetPost(sessionInfo, params)

	case api.GetPosts:
		// Handle posts request
		params := &api.GetPostsParams{}
		if err := parseParameters(message, params); err != nil {
			return nil, err
		}
		result, err = rsv.GetPosts(sessionInfo, params)

	case api.GetPostReaction:
		// Handle post reaction request
		params := &api.GetPostReactionParams{}
		if err := parseParameters(message, params); err != nil {
			return nil, err
		}
		result, err = rsv.GetPostReaction(sessionInfo, params)

	case api.GetReactionsOfPost:
		// Handle request for reactions of a post
		params := &api.GetReactionsOfPostParams{}
		if err := parseParameters(message, params); err != nil {
			return nil, err
		}
		result, err = rsv.GetReactionsOfPost(sessionInfo, params)

	case api.CreatePost:
		// Handle post creation request
		params := &api.CreatePostParams{}
		if err := parseParameters(message, params); err != nil {
			return nil, err
		}
		result, err = rsv.CreatePost(sessionInfo, params)

	case api.EditPost:
		// Handle post edit request
		params := &api.EditPostParams{}
		if err := parseParameters(message, params); err != nil {
			return nil, err
		}
		result, err = rsv.EditPost(sessionInfo, params)

	case api.RemovePost:
		// Handle post deletion request
		params := &api.RemovePostParams{}
		if err := parseParameters(message, params); err != nil {
			return nil, err
		}
		result, err = rsv.RemovePost(sessionInfo, params)

	case api.RemovePostReaction:
		// Handle post reaction removal request
		params := &api.RemovePostReactionParams{}
		if err := parseParameters(message, params); err != nil {
			return nil, err
		}
		result, err = rsv.RemovePostReaction(sessionInfo, params)

	case api.CreatePostReaction:
		// Handle post reaction creation request
		params := &api.CreatePostReactionParams{}
		if err := parseParameters(message, params); err != nil {
			return nil, err
		}
		result, err = rsv.CreatePostReaction(sessionInfo, params)

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
		} else if engErr := engiface.ToEngineError(err); engErr != nil {
			// Wrap engine-errors
			err = wwr.ReqErr{
				Code:    engErr.Code.String(),
				Message: engErr.Error(),
			}
		} else if err != nil {
			// Log an internal error
			rsv.logInternalError(err)
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
