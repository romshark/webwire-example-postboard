package memeng

import (
	"github.com/pkg/errors"
	"github.com/qbeon/webwire-example-postboard/server/apisrv/api"
	engiface "github.com/qbeon/webwire-example-postboard/server/apisrv/modules/engine"
)

// createUser creates a new user account.
// Returns a `ErrUsernameReserved` error if the username is already reserved
func (eng *engine) createUser(newAccount *UserAccount) error {
	// Ensure users with a similar username don't exist
	_, err := eng.findUserByUsername(newAccount.Profile.Username)
	if err == nil {
		return engiface.NewError(engiface.ErrUsernameReserved)
	}

	// Ensure users with a similar identifier don't exist
	_, similarIdentExists := eng.users[newAccount.Profile.Identifier]
	if similarIdentExists {
		return errors.Errorf(
			"user identifier collision (%s)",
			newAccount.Profile.Identifier.String(),
		)
	}

	// Successfully register the new account
	eng.users[newAccount.Profile.Identifier] = newAccount
	eng.usersByUsername[newAccount.Profile.Username] = newAccount

	return nil
}

// CreateUser implements the Engine interface
func (eng *engine) CreateUser(newAccount *api.User, passwordHash string) error {
	// Lock the engine store to execute the following operation transactionally
	eng.lock.Lock()
	defer eng.lock.Unlock()

	return eng.createUser(&UserAccount{
		Profile:  *newAccount,
		Password: passwordHash,
	})
}
