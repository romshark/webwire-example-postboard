package api

import "time"

// Post defines the structure of a post entity
type Post struct {
	Identifier  Identifier `json:"id"`
	Author      Identifier `json:"author"`
	Publication time.Time  `json:"publication"`
	Contents    string     `json:"contents"`
	LastEdit    *time.Time `json:"lastEdit"`

	// Reactions represents a list of reactions to this post
	// indexed by the reaction authors
	Reactions []PostReaction `json:"reaction"`
}
