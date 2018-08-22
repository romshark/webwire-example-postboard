package sessinfo

import (
	wwr "github.com/qbeon/webwire-go"
	"github.com/qbeon/webwire-example-postboard/server/apisrv/api"
)

// SessionInfo represents an API server specific session info object
type SessionInfo struct {
	UserIdentifier api.Identifier
	UserType       api.UserType
}

// Copy implements the wwr.SessionInfo interface
func (sinf *SessionInfo) Copy() wwr.SessionInfo {
	return &SessionInfo{
		UserIdentifier: sinf.UserIdentifier,
		UserType:       sinf.UserType,
	}
}

// Fields implements the wwr.SessionInfo interface
func (sinf *SessionInfo) Fields() []string {
	return []string{"id", "type"}
}

// Value implements the wwr.SessionInfo interface
func (sinf *SessionInfo) Value(fieldName string) interface{} {
	switch fieldName {
	case "id":
		return sinf.UserIdentifier
	case "type":
		return sinf.UserType.String()
	}
	return nil
}
