package resolver

import (
	"errors"

	"github.com/qbeon/webwire-messenger/server/apisrv/modules/authorizer"
	engiface "github.com/qbeon/webwire-messenger/server/apisrv/modules/engine"
	"github.com/qbeon/webwire-messenger/server/apisrv/modules/logger"
	"github.com/qbeon/webwire-messenger/server/apisrv/modules/passhash"
	"github.com/qbeon/webwire-messenger/server/apisrv/modules/validator"
)

type resolver struct {
	log            logger.Logger
	validator      validator.Validator
	authorizer     authorizer.Authorizer
	passwordHasher passhash.PasswordHasher
	engine         engiface.Engine
}

// New creates a new API resolver
func New(
	log logger.Logger,
	validator validator.Validator,
	authorizer authorizer.Authorizer,
	passwordHasher passhash.PasswordHasher,
	engine engiface.Engine,
) Resolver {
	if log == nil {
		panic(errors.New("resolver requires a logger instance"))
	}
	if validator == nil {
		panic(errors.New("resolver requires a validator instance"))
	}
	if engine == nil {
		panic(errors.New("resolver requires a engine instance"))
	}

	return &resolver{
		log:            log,
		validator:      validator,
		authorizer:     authorizer,
		passwordHasher: passwordHasher,
		engine:         engine,
	}
}
