package resolver

import (
	"github.com/qbeon/webwire-messenger/server/apisrv/api"
	"github.com/qbeon/webwire-messenger/server/apisrv/sessinfo"
)

// GetMessage implements the Resolver interface
func (rsv *resolver) GetMessage(
	session *sessinfo.SessionInfo,
	params *api.GetMessageParams,
) (interface{}, error) {
	// Instruct the engine to find the message
	messages, err := rsv.engine.GetMessages([]api.Identifier{params.Ident})
	if err != nil {
		return nil, err
	}

	// Message successfully retrieved
	return messages[0], nil
}
