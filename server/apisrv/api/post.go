package api

import "time"

// Post defines the structure of a post entity
type Post struct {
	Ident       Identifier `json:"id"`
	Author      Identifier `json:"author"`
	Publication time.Time  `json:"publication"`
	Contents    string     `json:"contents"`
	LastEdit    *time.Time `json:"lastEdit"`
}
