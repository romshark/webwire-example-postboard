package resolver

import (
	engiface "github.com/qbeon/webwire-messenger/server/apisrv/modules/engine"
)

// logInternalError logs an internal error to the error log
// and returns just a shallow error indicating that an internal error happened.
// This prevents internal errors from leaking through the API
// causing security issues
func (rsv *resolver) logInternalError(err error) error {
	rsv.log.Errf(err.Error())

	return engiface.NewError(engiface.ErrInternal)
}
