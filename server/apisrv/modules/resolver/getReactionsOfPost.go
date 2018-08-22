package resolver

import (
	"github.com/qbeon/webwire-messenger/server/apisrv/api"
	"github.com/qbeon/webwire-messenger/server/apisrv/sessinfo"
)

// GetReactionsOfPost implements the Resolver interface
func (rsv *resolver) GetReactionsOfPost(
	session *sessinfo.SessionInfo,
	params *api.GetReactionsOfPostParams,
) (interface{}, error) {
	// Instruct the engine to get n number of posts after the specified one
	result, err := rsv.engine.GetReactionsOfPost(params.PostIdent)
	if err != nil {
		return nil, err
	}

	// Posts successfully retrieved
	return result, nil
}
