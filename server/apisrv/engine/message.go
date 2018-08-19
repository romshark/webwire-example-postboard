package engine

import "time"

// Message defines the structure of a message entity
type Message struct {
	Identifier  Identifier `json:"id"`
	Author      Identifier `json:"author"`
	Publication time.Time  `json:"publication"`
	Contents    string     `json:"contents"`
	LastEdit    *time.Time `json:"lastEdit"`
	Archived    bool       `json:"archived"`

	// Reactions represents a list of reactions to this message
	// indexed by the reaction authors
	Reactions map[Identifier]*MessageReaction
}
