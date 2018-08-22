package api

import (
	"strings"

	"github.com/pkg/errors"
)

// PostReactionType represents a post reaction type
type PostReactionType int

const (
	// Celebration represents a reaction emotion
	Celebration PostReactionType = iota

	// Love represents a reaction emotion
	Love

	// Anger represents a reaction emotion
	Anger

	// Approval represents a reaction emotion
	Approval

	// Confusion represents a reaction emotion
	Confusion

	// Fear represents a reaction emotion
	Fear

	// Thinking represents a reaction emotion
	Thinking

	// Dislike represents a reaction emotion
	Dislike

	// Cry represents a reaction emotion
	Cry

	// Shock represents a reaction emotion
	Shock
)

// String stringifies the value
func (rt PostReactionType) String() string {
	switch rt {
	case Celebration:
		return "celebration"
	case Love:
		return "love"
	case Anger:
		return "anger"
	case Approval:
		return "approval"
	case Confusion:
		return "confusion"
	case Fear:
		return "fear"
	case Thinking:
		return "thinking"
	case Dislike:
		return "dislike"
	case Cry:
		return "cry"
	case Shock:
		return "shock"
	}
	panic(errors.Errorf("Invalid PostReactionType value: %d", int(rt)))
}

// Utf8Symbol returns the UTF8 symbol corresponding to the reaction
func (rt PostReactionType) Utf8Symbol() rune {
	switch rt {
	case Celebration:
		return '🎉'
	case Love:
		return '😍'
	case Anger:
		return '😡'
	case Approval:
		return '👍'
	case Confusion:
		return '😕'
	case Fear:
		return '😱'
	case Thinking:
		return '💭'
	case Dislike:
		return '👎'
	case Cry:
		return '😭'
	case Shock:
		return '😲'
	}
	return '🎁'
}

// FromString parses the value from a string
func (rt *PostReactionType) FromString(str string) error {
	switch strings.ToLower(str) {
	case "celebration":
		*rt = Celebration
	case "love":
		*rt = Love
	case "anger":
		*rt = Anger
	case "approval":
		*rt = Approval
	case "confusion":
		*rt = Confusion
	case "fear":
		*rt = Fear
	case "thinking":
		*rt = Thinking
	case "dislike":
		*rt = Dislike
	case "cry":
		*rt = Cry
	case "shock":
		*rt = Shock
	default:
		return errors.Errorf(
			"invalid string representation of the PostReactionType type: %s",
			str,
		)
	}
	return nil
}
