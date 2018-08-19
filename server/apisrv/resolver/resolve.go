package resolver

import (
	"fmt"

	wwr "github.com/qbeon/webwire-go"
	engiface "github.com/qbeon/webwire-messenger/server/apisrv/engine"
)

// Resolve implements the Resolver interface
func (rsv *resolver) Resolve(
	connection wwr.Connection,
	sessionInfo *engiface.SessionInfo,
	message wwr.Message,
) (wwr.Payload, error) {
	switch message.Name() {
	case "login":
		// Handle login request
		var params *LoginParams
		if err := parseParameters(message.Payload(), params); err != nil {
			return nil, err
		}
		return rsv.Login(sessionInfo, connection, params)

	case "logout":
		// Handle logout request
		return rsv.Logout(sessionInfo, connection)

	case "getMessages":
		// Handle messages request
		var params *GetMessagesParams
		if err := parseParameters(message.Payload(), params); err != nil {
			return nil, err
		}
		return rsv.GetMessages(sessionInfo, params)

	case "postMessage":
		// Handle message creation request
		var params *PostMessageParams
		if err := parseParameters(message.Payload(), params); err != nil {
			return nil, err
		}
		return rsv.PostMessage(sessionInfo, params)

	case "editMessage":
		// Handle message edit request
		var params *EditMessageParams
		if err := parseParameters(message.Payload(), params); err != nil {
			return nil, err
		}
		return rsv.EditMessage(sessionInfo, params)

	case "removeMessage":
		// Handle message deletion request
		var params *RemoveMessageParams
		if err := parseParameters(message.Payload(), params); err != nil {
			return nil, err
		}
		return rsv.RemoveMessage(sessionInfo, params)

	case "postMessageReaction":
		// Handle message reaction creation request
		var params *PostMessageReactionParams
		if err := parseParameters(message.Payload(), params); err != nil {
			return nil, err
		}
		return rsv.PostMessageReaction(sessionInfo, params)

	default:
		// Fail requests asking for an invalid API action with a typed error
		return nil, wwr.ReqErr{
			Code: "UNSUPPORTED_REQ_TYPE",
			Message: fmt.Sprintf(
				"Unsupported request type: '%s'",
				message.Name(),
			),
		}
	}
}
