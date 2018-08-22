package resolver

import (
	"github.com/qbeon/webwire-messenger/server/apisrv/api"
	"github.com/qbeon/webwire-messenger/server/apisrv/sessinfo"
)

// GetPostReaction implements the Resolver interface
func (rsv *resolver) GetPostReaction(
	session *sessinfo.SessionInfo,
	params *api.GetPostReactionParams,
) (interface{}, error) {
	// Instruct the engine to find the post
	posts, err := rsv.engine.GetPostReactions(
		[]api.Identifier{params.ReactionIdent},
	)
	if err != nil {
		return nil, err
	}

	// Post reaction successfully retrieved
	return posts[0], nil
}
