package memeng

import "github.com/qbeon/webwire-example-postboard/server/apisrv/api"

// PostReaction defines the post reaction in-memory storage representations
type PostReaction struct {
	// Reaction stores the actual post reaction data
	Reaction api.PostReaction

	// Archived indicates whether this post reaction is archived or not
	Archived bool
}
