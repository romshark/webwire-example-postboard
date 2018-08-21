package resolver

import (
	"time"

	"github.com/qbeon/webwire-messenger/server/apisrv/api"
	"github.com/qbeon/webwire-messenger/server/apisrv/modules/authorizer"
	"github.com/qbeon/webwire-messenger/server/apisrv/sessinfo"
)

// PostMessageReaction implements the Resolver interface
func (rsv *resolver) PostMessageReaction(
	session *sessinfo.SessionInfo,
	params *api.PostMessageReactionParams,
) (interface{}, error) {
	if err := rsv.authorizer.MeetsAll(
		session,
		authorizer.IsAuthenticated(
			"only administrators and authenticated users are allowed "+
				"to post message reactions",
		),
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

	// Check authorization, ensure the user is not the owner of this
	// message because users are not allowed to react on their on messages
	if err := rsv.authorizer.MeetsAll(
		session,
		authorizer.IsNotResourceOwner(
			retrieved[0].Author,
			"users are not allowed to post reactions on their own messages",
		),
	); err != nil {
		return nil, err
	}

	// Validate reaction description
	if err := rsv.validator.ReactionDescription(
		params.Description,
	); err != nil {
		return nil, err
	}

	newIdent := api.NewIdentifier()

	// Instruct the engine to add the new message reaction
	err = rsv.engine.PostMessageReaction(
		params.MessageIdent,
		&api.MessageReaction{
			Ident:       newIdent,
			Author:      session.UserIdentifier,
			Type:        params.Type,
			Description: params.Description,
			Creation:    time.Now().UTC(),
		},
	)
	if err := rsv.handleError(err); err != nil {
		return nil, err
	}

	return api.PostMessageReactionReturn{
		MessageReactionIdent: newIdent,
	}, nil
}
