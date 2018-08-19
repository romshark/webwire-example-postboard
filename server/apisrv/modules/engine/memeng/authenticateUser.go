package memeng

import (
	"github.com/qbeon/webwire-messenger/server/apisrv/api"
	engiface "github.com/qbeon/webwire-messenger/server/apisrv/modules/engine"
)

// findUserByUsername returns the user account identified by the given
// username. Returns a user-not-found error if no user was identified
func (eng *engine) findUserByUsername(
	username string,
) (*UserAccount, error) {
	user, exists := eng.usersByUsername[username]
	if !exists {
		return nil, engiface.NewError(engiface.ErrUserNotFound)
	}
	return user, nil
}

// AuthenticateUser implements the Engine interface
func (eng *engine) AuthenticateUser(
	username string,
) (
	*api.User,
	string,
	error,
) {
	// Lock the engine store to execute the following operation transactionally
	eng.lock.Lock()
	defer eng.lock.Unlock()

	user, err := eng.findUserByUsername(username)
	if err != nil {
		return nil, "", err
	}

	return &user.Profile, user.Password, nil
}
