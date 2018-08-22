package setup

import (
	"context"
	"time"

	"github.com/qbeon/webwire-example-postboard/server/apisrv/api"
	"github.com/qbeon/webwire-example-postboard/server/client"
	"github.com/stretchr/testify/require"
)

// CreatePostReaction posts a new reaction to any given post
// using the given client and contents
// and verifies whether the reaction correctly created
func (h *Helper) CreatePostReaction(
	postAuthor client.ApiClient,
	reactionAuthor client.ApiClient,
	postIdent api.Identifier,
	reactionType api.PostReactionType,
	description string,
) *api.PostReaction {
	// Get previous reactions
	postBefore, err := postAuthor.GetPost(
		context.Background(),
		api.GetPostParams{
			Ident: postIdent,
		},
	)
	require.NoError(h.t, err)
	require.NotNil(h.t, postBefore)

	// Get all reactions
	reactionsBefore, err := postAuthor.GetReactionsOfPost(
		context.Background(),
		api.GetReactionsOfPostParams{
			PostIdent: postBefore.Ident,
		},
	)
	if reactionsBefore == nil {
		reactionsBefore = make([]*api.PostReaction, 0)
	}

	// Post a new reaction
	reactionIdent, err := reactionAuthor.CreatePostReaction(
		context.Background(),
		api.CreatePostReactionParams{
			PostIdent:   postIdent,
			Type:        reactionType,
			Description: description,
		},
	)
	require.NoError(h.t, err)
	require.False(h.t, reactionIdent.IsNull())

	// Refetch the updated reactions
	reactionsAfter, err := postAuthor.GetReactionsOfPost(
		context.Background(),
		api.GetReactionsOfPostParams{
			PostIdent: postBefore.Ident,
		},
	)
	if reactionsAfter == nil {
		reactionsAfter = make([]*api.PostReaction, 0)
	}

	// Verify the new reaction was actually created
	require.Len(h.t, reactionsAfter, len(reactionsBefore)+1)

	// Verify previous reactions stayed the same
	require.Equal(h.t,
		reactionsBefore,
		reactionsAfter[:len(reactionsAfter)-1],
	)

	// Verify the new reaction
	lastReaction := reactionsAfter[len(reactionsAfter)-1]
	require.False(h.t, lastReaction.Ident.IsNull())
	require.Equal(h.t,
		reactionAuthor.Identifier().String(),
		lastReaction.Author.String(),
	)
	require.Equal(h.t, description, lastReaction.Description)
	require.Equal(h.t, reactionType, lastReaction.Type)
	require.WithinDuration(h.t,
		time.Now().UTC(),
		lastReaction.Creation,
		h.ts.MaxCreationTimeDeviation(),
	)

	return reactionsAfter[len(reactionsAfter)-1]
}
