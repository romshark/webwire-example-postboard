package resolver

import (
	"github.com/qbeon/webwire-example-postboard/server/apisrv/api"
	"github.com/qbeon/webwire-example-postboard/server/apisrv/sessinfo"
)

// GetUser implements the Resolver interface
func (rsv *resolver) GetUser(
	session *sessinfo.SessionInfo,
	params *api.GetUserParams,
) (interface{}, error) {
	// Instruct the engine to find the user profile
	profile, err := rsv.engine.GetUser(params.Ident)
	if err != nil {
		return nil, err
	}

	return profile, nil
}
