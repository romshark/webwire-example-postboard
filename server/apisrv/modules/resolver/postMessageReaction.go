package resolver

import (
	"github.com/qbeon/webwire-messenger/server/apisrv/api"
	"github.com/qbeon/webwire-messenger/server/apisrv/modules/authorizer"
	"github.com/qbeon/webwire-messenger/server/apisrv/sessinfo"
)

// PostMessageReaction implements the Resolver interface
func (rsv *resolver) PostMessageReaction(
	session *sessinfo.SessionInfo,
	params *api.PostMessageReactionParams,
) (interface{}, error) {
	// Check authorization, message reactions can only be added by
	// authenticated users (regular users and administrators)
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

	// Check authorization, ensure the user is not the owner of this
	// message because users are not allowed to react on their on messages
	if err := rsv.authorizer.MeetsAll(
		session,
		authorizer.IsNotResourceOwner{
			ResourceOwner: retrieved[0].Author,
		},
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
			Author:      params.AuthorIdent,
			Type:        params.Type,
			Description: params.Description,
		},
	)
	if err := rsv.handleError(err); err != nil {
		return nil, err
	}

	return api.PostMessageReactionReturn{
		MessageReactionIdent: newIdent,
	}, nil
}
