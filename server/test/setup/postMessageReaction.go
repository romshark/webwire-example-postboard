package setup

import (
	"context"
	"time"

	"github.com/qbeon/webwire-messenger/server/apisrv/api"
	"github.com/qbeon/webwire-messenger/server/client"
	"github.com/stretchr/testify/require"
)

// PostMessageReaction posts a new reaction to any given message
// using the given client and contents
// and verifies whether the reaction correctly posted
func (h *Helper) PostMessageReaction(
	messageAuthor client.ApiClient,
	reactionAuthor client.ApiClient,
	messageIdent api.Identifier,
	reactionType api.MessageReactionType,
	description string,
) *api.MessageReaction {
	// Get previous reactions
	messageBefore, err := messageAuthor.GetMessage(
		context.Background(),
		api.GetMessageParams{
			Ident: messageIdent,
		},
	)
	require.NoError(h.t, err)
	require.NotNil(h.t, messageBefore)
	reactionsBefore := messageBefore.Reactions
	if reactionsBefore == nil {
		reactionsBefore = make([]api.MessageReaction, 0)
	}

	// Post a new reaction
	reactionIdent, err := reactionAuthor.PostMessageReaction(
		context.Background(),
		api.PostMessageReactionParams{
			MessageIdent: messageIdent,
			Type:         reactionType,
			Description:  description,
		},
	)
	require.NoError(h.t, err)
	require.False(h.t, reactionIdent.IsNull())

	// Get previous reactions
	messageAfter, err := messageAuthor.GetMessage(
		context.Background(),
		api.GetMessageParams{
			Ident: messageIdent,
		},
	)
	require.NoError(h.t, err)
	require.NotNil(h.t, messageAfter)
	reactionsAfter := messageAfter.Reactions

	// Verify the new reaction was actually posted
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
		1*time.Second,
	)

	return &reactionsAfter[len(reactionsAfter)-1]
}
