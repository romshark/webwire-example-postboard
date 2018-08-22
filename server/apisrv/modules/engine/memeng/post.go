package memeng

import "github.com/qbeon/webwire-messenger/server/apisrv/api"

// Post defines the post in-memory storage representations
type Post struct {
	// Post stores the actual post data
	Post api.Post

	// Reactions links the corresponding reactions in chronological order
	Reactions []*PostReaction

	// Archived indicates whether this post is archived or not
	Archived bool
}
