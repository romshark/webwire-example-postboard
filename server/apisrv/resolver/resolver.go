package resolver

import (
	"errors"

	"github.com/qbeon/webwire-messenger/server/apisrv/authorizer"
	engiface "github.com/qbeon/webwire-messenger/server/apisrv/engine"
	"github.com/qbeon/webwire-messenger/server/apisrv/logger"
	"github.com/qbeon/webwire-messenger/server/apisrv/validator"
)

type resolver struct {
	log        logger.Logger
	validator  validator.Validator
	authorizer authorizer.Authorizer
	engine     engiface.Engine
}

// New creates a new API resolver
func New(
	log logger.Logger,
	validator validator.Validator,
	authorizer authorizer.Authorizer,
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
		log:        log,
		validator:  validator,
		authorizer: authorizer,
		engine:     engine,
	}
}
