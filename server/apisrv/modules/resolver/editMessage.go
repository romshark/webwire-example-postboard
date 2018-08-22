package resolver

import (
	"time"

	"github.com/qbeon/webwire-messenger/server/apisrv/api"
	"github.com/qbeon/webwire-messenger/server/apisrv/modules/authorizer"
	"github.com/qbeon/webwire-messenger/server/apisrv/sessinfo"
)

// EditMessage implements the Resolver interface
func (rsv *resolver) EditMessage(
	session *sessinfo.SessionInfo,
	params *api.EditMessageParams,
) (interface{}, error) {
	// Validate new contents before proceeding
	if err := rsv.validator.MessageContents(params.NewContents); err != nil {
		return nil, err
	}

	// Check authorization
	if err := rsv.authorizer.MeetsAll(
		session,
		authorizer.IsAuthenticated("guest clients are allows to read only"),
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

	// Check authorization
	if err := rsv.authorizer.MeetsAll(
		session,
		authorizer.IsResourceOwner(
			retrieved[0].Author,
			"only authors can edit their message",
		),
	); err != nil {
		return nil, err
	}

	// Instruct the engine to edit the message
	if err := rsv.engine.EditMessage(
		params.MessageIdent,
		time.Now().UTC(),
		params.NewContents,
	); err != nil {
		return nil, err
	}

	// Message successfully edited
	return nil, nil
}
