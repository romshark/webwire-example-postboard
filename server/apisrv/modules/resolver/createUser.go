package resolver

import (
	"time"

	"github.com/pkg/errors"
	"github.com/qbeon/webwire-messenger/server/apisrv/api"
	"github.com/qbeon/webwire-messenger/server/apisrv/modules/authorizer"
	"github.com/qbeon/webwire-messenger/server/apisrv/sessinfo"
)

// CreateUser implements the Resolver interface
func (rsv *resolver) CreateUser(
	session *sessinfo.SessionInfo,
	params *api.CreateUserParams,
) (interface{}, error) {
	// Check authorization
	if err := rsv.authorizer.MeetsAll(
		session,
		authorizer.IsAdmin(
			"only administrators are allowed to create new user accounts",
		),
	); err != nil {
		return nil, err
	}

	// Validate parameters: username, first name, last name and password
	if err := rsv.validator.Username(params.Username); err != nil {
		return nil, err
	}
	if err := rsv.validator.Name(params.FirstName); err != nil {
		return nil, err
	}
	if err := rsv.validator.Name(params.LastName); err != nil {
		return nil, err
	}
	if err := rsv.validator.Password(params.Password); err != nil {
		return nil, err
	}

	// Encrypt password
	hashedPassword, err := rsv.passwordHasher.Hash(params.Password)
	if err != nil {
		return nil, rsv.logInternalError(errors.Wrap(
			err,
			"couldn't hash password",
		))
	}

	// Generate a new universally unique account identifier
	newIdent := api.NewIdentifier()

	// Instruct the engine to create a new user account
	err = rsv.engine.CreateUser(&api.User{
		Identifier:   newIdent,
		FirstName:    params.FirstName,
		LastName:     params.LastName,
		Username:     params.Username,
		Registration: time.Now().UTC(),
		Reputation:   0,
		Type:         params.Type,
	}, hashedPassword)
	if err := rsv.handleError(err); err != nil {
		return nil, err
	}

	// Message successfully posted
	return api.CreateUserReturn{
		UserIdent: newIdent,
	}, nil
}
