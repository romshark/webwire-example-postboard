package resolver

import (
	"github.com/qbeon/webwire-messenger/server/apisrv/api"
	"github.com/qbeon/webwire-messenger/server/apisrv/modules/authorizer"
	"github.com/qbeon/webwire-messenger/server/apisrv/sessinfo"
)

// GetUser implements the Resolver interface
func (rsv *resolver) GetUser(
	session *sessinfo.SessionInfo,
	params *api.GetUserParams,
) (interface{}, error) {
	// Check authorization, ensure the user is either an administrator
	// or the owner of the profile
	if err := rsv.authorizer.MeetsEitherOf(
		session,
		authorizer.IsAdmin{},
		authorizer.IsResourceOwner{
			ResourceOwner: params.Ident,
		},
	); err != nil {
		return nil, err
	}

	// Instruct the engine to find the user profile
	profile, err := rsv.engine.GetUser(params.Ident)
	if err := rsv.handleError(err); err != nil {
		return nil, err
	}

	return profile, nil
}
