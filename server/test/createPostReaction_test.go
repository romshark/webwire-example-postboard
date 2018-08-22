package test

import (
	"testing"

	"github.com/qbeon/webwire-example-postboard/server/apisrv/api"
	"github.com/qbeon/webwire-example-postboard/server/test/setup"
)

// TestCreatePostReaction tests creating post reactions
func TestCreatePostReaction(t *testing.T) {

	t.Run("AsRoot_ToUsersPost", func(t *testing.T) {
		//t.Parallel()
		ts := setup.New(t, setupConf)
		defer ts.Teardown()

		root := ts.NewAdminClient("root", "root")

		// Create random test users
		_, _, user1 := ts.Helper.CreateUserRand(root, api.UtUser)

		post := ts.Helper.CreatePost(user1, api.CreatePostParams{
			Contents: "now look at thaaat!",
		})

		ts.Helper.CreatePostReaction(
			user1,
			root,
			post.Ident,
			api.Shock,
			"wooow!",
		)
	})

	t.Run("AsUser_ToRootPost", func(t *testing.T) {
		//t.Parallel()
		ts := setup.New(t, setupConf)
		defer ts.Teardown()

		root := ts.NewAdminClient("root", "root")

		// Create random test users
		_, _, user1 := ts.Helper.CreateUserRand(root, api.UtUser)

		post := ts.Helper.CreatePost(root, api.CreatePostParams{
			Contents: "we did it!",
		})

		ts.Helper.CreatePostReaction(
			root,
			user1,
			post.Ident,
			api.Shock,
			"Huraah!",
		)
	})
}
