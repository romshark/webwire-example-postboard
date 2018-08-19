package resolver

import (
	wwr "github.com/qbeon/webwire-go"
	engiface "github.com/qbeon/webwire-messenger/server/apisrv/engine"
)

// handleError helps handling errors returned by engine calls.
// It'll log unexpected errors as internal errors and return expected ones
func (rsv *resolver) handleError(err error) error {
	if err == nil {
		return nil
	}

	if err := engiface.ToEngineError(err); err != nil {
		return wwr.ReqErr{
			Code:    err.Code.String(),
			Message: err.Error(),
		}
	} else if err != nil {
		// Log an internal error
		rsv.logInternalError(err)
	}

	return nil
}
