package api

import "time"

// User defines the structure of a user profile
type User struct {
	Identifier Identifier `json:"id"`
	FirstName  string     `json:"firstName"`
	LastName   string     `json:"lastName"`

	// Username uniquely identifies the user in the system,
	// it's also used as the login username
	Username string `json:"Username"`

	// Registration defines when this account was created
	Registration time.Time `json:"registration"`

	// Reputation defines the amount of upvotes
	// on posts and reactions by this user
	Reputation float64 `json:"reputation"`

	// Type defines whether this user is an admin, a regular user or a guest
	Type UserType `json:"type"`
}
