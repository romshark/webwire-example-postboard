package memeng

import (
	"errors"

	"github.com/qbeon/webwire-example-postboard/server/apisrv/api"
	engiface "github.com/qbeon/webwire-example-postboard/server/apisrv/modules/engine"
)

// createPost posts the new post `newPost` to the store
// updating the identifier index.
// Returns an error if a post with a similar identifier already exists
func (eng *engine) createPost(newPost *api.Post) error {
	if newPost == nil {
		return errors.New("can't create nil post")
	}

	// Return error if a post with a similar identifier is already defined,
	// ignore the returned error because it'll definitely be a not-found error
	post, _ := eng.findPostByIdent(newPost.Ident)
	if post != nil {
		return engiface.NewError(engiface.ErrPostAlreadyExists)
	}

	// Commit the post to the chronological history
	eng.posts = append(eng.posts, &Post{
		Post:     *newPost,
		Archived: false,
	})

	// Update post identifier index
	eng.postsByIdent[newPost.Ident] = len(eng.posts) - 1

	return nil
}

// CreatePost implements the Engine interface
func (eng *engine) CreatePost(newPost *api.Post) error {
	// Lock the engine store to execute the following operation transactionally
	eng.lock.Lock()
	defer eng.lock.Unlock()

	return eng.createPost(newPost)
}
