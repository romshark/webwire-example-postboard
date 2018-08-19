package engine

import (
	"strings"

	"github.com/pkg/errors"
)

// MessageReactionType represents a message reaction type
type MessageReactionType int

const (
	// Celebration
	Celebration MessageReactionType = iota

	// Love
	Love

	// Anger
	Anger

	// Approval
	Approval

	// Confusion
	Confusion

	// Fear
	Fear

	// Thinking
	Thinking

	// Dislike
	Dislike

	// Cry
	Cry

	// Shock
	Shock
)

// String stringifies the value
func (rt MessageReactionType) String() string {
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
	panic(errors.Errorf("Invalid MessageReactionType value: %d", int(rt)))
}

// Utf8Symbol returns the UTF8 symbol corresponding to the reaction
func (rt MessageReactionType) Utf8Symbol() rune {
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
func (rt *MessageReactionType) FromString(str string) error {
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
			"invalid string representation of the MessageReactionType type: %s",
			str,
		)
	}
	return nil
}
