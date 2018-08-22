package resolver

import (
	"github.com/qbeon/webwire-messenger/server/apisrv/api"
	"github.com/qbeon/webwire-messenger/server/apisrv/sessinfo"
)

// GetMessages implements the Resolver interface
func (rsv *resolver) GetMessages(
	session *sessinfo.SessionInfo,
	params *api.GetMessagesParams,
) (interface{}, error) {
	// Validate query limit
	if err := rsv.validator.MessagesQueryLimit(params.Limit); err != nil {
		return nil, err
	}

	// Instruct the engine to get n number of messages after the specified one
	result, err := rsv.engine.GetMessagesAfter(
		params.After,
		params.Limit,
	)
	if err != nil {
		return nil, err
	}

	// Messages successfully retrieved
	return result, nil
}
