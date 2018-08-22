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
	// Check authorization
	if err := rsv.authorizer.MeetsEitherOf(
		session,
		authorizer.IsAuthenticated(
			"only authenticated clients are allowed to retrieve a user profile",
		),
	); err != nil {
		return nil, err
	}

	// Instruct the engine to find the user profile
	profile, err := rsv.engine.GetUser(params.Ident)
	if err != nil {
		return nil, err
	}

	return profile, nil
}
