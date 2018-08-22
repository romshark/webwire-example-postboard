package resolver

import (
	"github.com/qbeon/webwire-messenger/server/apisrv/api"
	"github.com/qbeon/webwire-messenger/server/apisrv/modules/authorizer"
	"github.com/qbeon/webwire-messenger/server/apisrv/sessinfo"
)

// RemoveMessage implements the Resolver interface
func (rsv *resolver) RemoveMessage(
	session *sessinfo.SessionInfo,
	params *api.RemoveMessageParams,
) (interface{}, error) {
	// Check authorization, ensure the user is not a guest,
	// because guests are allowed to read only
	if err := rsv.authorizer.MeetsAll(
		session,
		authorizer.IsAuthenticated(
			"only authenticated clients are allowed to remove messages",
		),
	); err != nil {
		return nil, err
	}

	// Perform an engine call to retrieve the targeted message
	// to verify authorization
	retrieved, err := rsv.engine.GetMessages(
		[]api.Identifier{params.MessageIdent},
	)
	if err != nil {
		return nil, err
	}

	// Check authorization, ensure the user is either the author of the message
	// or an administrator, because administrators are also allowed to remove
	// messages
	if err := rsv.authorizer.MeetsEitherOf(
		session,
		authorizer.IsAdmin(
			"only administrators and the author the message "+
				"are allowed to remove a message",
		),
		authorizer.IsResourceOwner(
			retrieved[0].Author,
			"only administrators and the author the message "+
				"are allowed to remove a message",
		),
	); err != nil {
		return nil, err
	}

	// Instruct the engine to remove the message
	if err := rsv.engine.RemoveMessage(params.MessageIdent); err != nil {
		return nil, err
	}

	// Successfully removed the message
	return nil, nil
}
