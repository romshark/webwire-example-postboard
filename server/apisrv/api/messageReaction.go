package api

import "time"

// MessageReaction represents a reaction to a message
type MessageReaction struct {
	Ident       Identifier          `json:"id"`
	Author      Identifier          `json:"author"`
	Type        MessageReactionType `json:"type"`
	Description string              `json:"description"`
	Creation    time.Time           `json:"creation"`
}
