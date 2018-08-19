package engine

// MessageReaction represents a reaction to a message
type MessageReaction struct {
	Ident       Identifier          `json:"id"`
	Author      Identifier          `json:"author"`
	Type        MessageReactionType `json:"type"`
	Description string              `json:"description"`
}
