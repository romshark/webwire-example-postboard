package resolver

import (
	"encoding/json"

	wwr "github.com/qbeon/webwire-go"
	engiface "github.com/qbeon/webwire-messenger/server/apisrv/engine"
)

// GetMessages implements the Resolver interface
func (rsv *resolver) GetMessages(
	session *engiface.SessionInfo,
	params *GetMessagesParams,
) (wwr.Payload, error) {
	// Validate query limit
	if err := rsv.validator.MessagesQueryLimit(params.Limit); err != nil {
		return nil, err
	}

	// Instruct the engine to get n number of messages after the specified one
	result, err := rsv.engine.GetMessagesAfter(
		params.After,
		params.Limit,
	)
	if err := rsv.handleError(err); err != nil {
		return nil, err
	}

	// Marshal the resulting slice of messages into JSON
	res, err := json.Marshal(result)
	if err != nil {
		return nil, rsv.logInternalError(err)
	}

	return wwr.NewPayload(wwr.EncodingUtf8, res), nil
}
