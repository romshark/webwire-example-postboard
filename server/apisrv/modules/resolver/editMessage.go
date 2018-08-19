package resolver

import (
	"time"

	wwr "github.com/qbeon/webwire-go"
	"github.com/qbeon/webwire-messenger/server/apisrv/api"
	"github.com/qbeon/webwire-messenger/server/apisrv/modules/authorizer"
	"github.com/qbeon/webwire-messenger/server/apisrv/sessinfo"
)

// EditMessage implements the Resolver interface
func (rsv *resolver) EditMessage(
	session *sessinfo.SessionInfo,
	params *api.EditMessageParams,
) (wwr.Payload, error) {
	// Validate new contents before proceeding
	if err := rsv.validator.MessageContents(params.NewContents); err != nil {
		return nil, err
	}

	// Check authorization, ensure the user is not a guest,
	// because guests are allowed to read only
	if err := rsv.authorizer.MeetsAll(
		session,
		authorizer.IsAuthenticated{},
	); err != nil {
		return nil, err
	}

	// Perform an engine call to retrieve the targeted message
	// to verify authorization
	retrieved, err := rsv.engine.GetMessages(
		[]api.Identifier{params.MessageIdent},
	)
	if err := rsv.handleError(err); err != nil {
		return nil, err
	}

	// Check authorization, ensure the user is the owner of this message
	// because only the author of a message can edit it
	if err := rsv.authorizer.MeetsAll(
		session,
		authorizer.IsResourceOwner{
			ResourceOwner: retrieved[0].Author,
		},
	); err != nil {
		return nil, err
	}

	// Instruct the engine to edit the message
	err = rsv.engine.EditMessage(
		params.MessageIdent,
		time.Now().UTC(),
		params.NewContents,
	)
	if err := rsv.handleError(err); err != nil {
		return nil, err
	}

	// Message successfully edited
	return nil, nil
}
