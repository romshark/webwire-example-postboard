package resolver

import (
	"time"

	wwr "github.com/qbeon/webwire-go"
	"github.com/qbeon/webwire-messenger/server/apisrv/authorizer"
	engiface "github.com/qbeon/webwire-messenger/server/apisrv/engine"
)

// PostMessage implements the Resolver interface
func (rsv *resolver) PostMessage(
	session *engiface.SessionInfo,
	params *PostMessageParams,
) (wwr.Payload, error) {
	// Check authorization, ensure the user is authenticated,
	// because guests are allowed to read only
	if err := rsv.authorizer.MeetsAll(
		session,
		authorizer.IsAuthenticated{},
	); err != nil {
		return nil, err
	}

	// Instruct the engine to create a new message
	err := rsv.engine.PostMessage(&engiface.Message{
		Identifier:  engiface.NewIdentifier(),
		Author:      session.UserIdentifier,
		Publication: time.Now().UTC(),
		Contents:    params.Contents,
	})
	if err := rsv.handleError(err); err != nil {
		return nil, err
	}

	// Message successfully posted
	return nil, nil
}
