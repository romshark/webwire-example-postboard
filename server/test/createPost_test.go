package test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/qbeon/webwire-example-postboard/server/apisrv/api"
	"github.com/qbeon/webwire-example-postboard/server/client"
	"github.com/qbeon/webwire-example-postboard/server/test/setup"
)

// TestCreatePost tests creating a post and retrieving it
func TestCreatePost(t *testing.T) {
	testCreatingAndReading := func(
		t *testing.T,
		ts *setup.TestSetup,
		author client.ApiClient,
		postContents []string,
		readers []client.ApiClient,
	) {
		// Post new posts
		created := make([]*api.Post, len(postContents))
		for i, contents := range postContents {
			created[i] = ts.Helper.CreatePost(author, api.CreatePostParams{
				Contents: contents,
			})
		}

		// Ensure all readers get the correct posts
		for _, reader := range readers {
			posts, err := reader.GetPosts(
				context.Background(),
				api.GetPostsParams{
					After: nil,
					Limit: 10,
				},
			)
			require.NoError(t, err)
			require.Len(t, posts, len(postContents))

			// Verify the posts
			for i := range created {
				retrievedPost := posts[i]

				// Iterate through the created posts in reversed order
				createdItr := len(created) - 1 - i
				createdPost := created[createdItr]

				require.Equal(t,
					createdPost.Ident.String(),
					retrievedPost.Ident.String(),
				)
				require.Equal(t,
					*author.Identifier(),
					retrievedPost.Author,
				)
				require.Equal(t,
					postContents[createdItr],
					retrievedPost.Contents,
				)
				require.Nil(t, createdPost.LastEdit)
				require.Equal(t,
					createdPost.Publication,
					retrievedPost.Publication,
				)
				require.WithinDuration(t,
					time.Now().UTC(),
					retrievedPost.Publication,
					ts.MaxCreationTimeDeviation(),
				)
			}
		}
	}

	t.Run("AsRoot", func(t *testing.T) {
		//t.Parallel()
		ts := setup.New(t, setupConf)
		defer ts.Teardown()

		root := ts.NewAdminClient("root", "root")

		// Create random test users
		_, _, user1 := ts.Helper.CreateUserRand(root, api.UtUser)
		_, _, user2 := ts.Helper.CreateUserRand(root, api.UtUser)
		guest := ts.NewGuestClient()

		testCreatingAndReading(
			t,
			ts,   // Test setup
			root, // Author
			[]string{"test post"},              // Post contents
			Readers{root, guest, user1, user2}, // Readers
		)
	})

	t.Run("AsUser", func(t *testing.T) {
		//t.Parallel()
		ts := setup.New(t, setupConf)
		defer ts.Teardown()

		root := ts.NewAdminClient("root", "root")

		// Create random test users
		_, _, user1 := ts.Helper.CreateUserRand(root, api.UtUser)
		_, _, user2 := ts.Helper.CreateUserRand(root, api.UtUser)
		guest := ts.NewGuestClient()

		// Ensure all users can read the post including the author himself
		// and the post is as expected
		testCreatingAndReading(
			t,
			ts,    // Test setup
			user1, // Author
			[]string{"test post"},              // Post contents
			Readers{root, guest, user1, user2}, // Readers
		)
	})

	t.Run("AsRoot_Multiple", func(t *testing.T) {
		//t.Parallel()
		ts := setup.New(t, setupConf)
		defer ts.Teardown()

		root := ts.NewAdminClient("root", "root")

		// Create random test users
		_, _, user1 := ts.Helper.CreateUserRand(root, api.UtUser)
		_, _, user2 := ts.Helper.CreateUserRand(root, api.UtUser)
		guest := ts.NewGuestClient()

		// Ensure all users can read the post including the author himself
		// and the post is as expected
		testCreatingAndReading(
			t,
			ts,   // Test setup
			root, // Author
			[]string{
				"first test post",
				"second test post",
				"third test post",
				"fourth test post",
				"fifth test post",
			}, // Post contents
			Readers{root, guest, user1, user2}, // Readers
		)
	})

	t.Run("AsUser_Multiple", func(t *testing.T) {
		//t.Parallel()
		ts := setup.New(t, setupConf)
		defer ts.Teardown()

		root := ts.NewAdminClient("root", "root")

		// Create random test users
		_, _, user1 := ts.Helper.CreateUserRand(root, api.UtUser)
		_, _, user2 := ts.Helper.CreateUserRand(root, api.UtUser)
		guest := ts.NewGuestClient()

		// Ensure all users can read the post including the author himself
		// and the post is as expected
		testCreatingAndReading(
			t,
			ts,    // Test setup
			user1, // Author
			[]string{
				"first test post",
				"second test post",
				"third test post",
				"fourth test post",
				"fifth test post",
			}, // Post contents
			Readers{root, guest, user1, user2}, // Readers
		)
	})
}
