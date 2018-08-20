package resolver

import (
	"time"

	"github.com/qbeon/webwire-messenger/server/apisrv/api"
	"github.com/qbeon/webwire-messenger/server/apisrv/modules/authorizer"
	"github.com/qbeon/webwire-messenger/server/apisrv/sessinfo"
)

// PostMessage implements the Resolver interface
func (rsv *resolver) PostMessage(
	session *sessinfo.SessionInfo,
	params *api.PostMessageParams,
) (interface{}, error) {
	// Check authorization, ensure the user is authenticated,
	// because guests are allowed to read only
	if err := rsv.authorizer.MeetsAll(
		session,
		authorizer.IsAuthenticated{},
	); err != nil {
		return nil, err
	}

	newIdent := api.NewIdentifier()

	// Instruct the engine to create a new message
	err := rsv.engine.PostMessage(&api.Message{
		Identifier:  newIdent,
		Author:      session.UserIdentifier,
		Publication: time.Now().UTC(),
		Contents:    params.Contents,
	})
	if err := rsv.handleError(err); err != nil {
		return nil, err
	}

	// Message successfully posted
	return api.PostMessageReturn{
		MessageIdent: newIdent,
	}, nil
}
