package memeng

import (
	"errors"

	"github.com/qbeon/webwire-messenger/server/apisrv/api"
)

// createPostReaction sets a new reaction to the post identified by
// `postIdent`
// A not-found error is returned if it either doesn't exist
// or if it was already deleted
func (eng *engine) createPostReaction(
	postIdent api.Identifier,
	reaction *api.PostReaction,
) error {
	if reaction == nil {
		return errors.New("can't post a nil post reaction")
	}

	post, err := eng.findPostByIdent(postIdent)
	if err != nil {
		return err
	}

	// Create a new post reaction
	newPostReaction := &PostReaction{
		Reaction: *reaction,
	}

	// Set the reaction
	post.Reactions = append(post.Reactions, newPostReaction)

	// Update reaction.identifier index
	eng.postReactionsByIdent[reaction.Ident] = newPostReaction

	return nil
}

// CreatePostReaction implements the Engine interface
func (eng *engine) CreatePostReaction(
	postIdent api.Identifier,
	reaction *api.PostReaction,
) error {
	// Lock the engine store to execute the following operation transactionally
	eng.lock.Lock()
	defer eng.lock.Unlock()

	return eng.createPostReaction(postIdent, reaction)
}
