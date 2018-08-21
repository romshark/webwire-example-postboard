package api

// Method represents an API method name
type Method string

const (
	Login               Method = "Login"
	Logout              Method = "Logout"
	CreateUser          Method = "CreateUser"
	GetUser             Method = "GetUser"
	GetMessage          Method = "GetMessage"
	GetMessages         Method = "GetMessages"
	PostMessage         Method = "PostMessage"
	EditMessage         Method = "EditMessage"
	RemoveMessage       Method = "RemoveMessage"
	PostMessageReaction Method = "PostMessageReaction"
)
