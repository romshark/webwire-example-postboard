package memeng

import "github.com/qbeon/webwire-messenger/server/apisrv/api"

// UserAccount defines a user account in-memory storage representation
type UserAccount struct {
	Profile api.User

	// Password stores this user's password in a secure hashed form
	Password string
}
