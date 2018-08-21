package api

// Method represents an API method name
type Method string

const (
	// Login represents an API method
	Login Method = "Login"

	// Logout represents an API method
	Logout Method = "Logout"

	// CreateUser represents an API method
	CreateUser Method = "CreateUser"

	// GetUser represents an API method
	GetUser Method = "GetUser"

	// GetMessage represents an API method
	GetMessage Method = "GetMessage"

	// GetMessages represents an API method
	GetMessages Method = "GetMessages"

	// PostMessage represents an API method
	PostMessage Method = "PostMessage"

	// EditMessage represents an API method
	EditMessage Method = "EditMessage"

	// RemoveMessage represents an API method
	RemoveMessage Method = "RemoveMessage"

	// PostMessageReaction represents an API method
	PostMessageReaction Method = "PostMessageReaction"
)
