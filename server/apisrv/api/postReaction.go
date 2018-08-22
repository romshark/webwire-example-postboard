package api

import "time"

// PostReaction represents a reaction to a post
type PostReaction struct {
	Ident       Identifier       `json:"id"`
	Author      Identifier       `json:"author"`
	Type        PostReactionType `json:"type"`
	Description string           `json:"description"`
	Creation    time.Time        `json:"creation"`
}
