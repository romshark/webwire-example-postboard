package setup

import (
	wwr "github.com/qbeon/webwire-go"
	"github.com/stretchr/testify/require"
)

// VerifyUnauthErr helps ensuring that the given error is an unauth-error
func (h *Helper) VerifyUnauthErr(err error) {
	require.Error(h.t, err)
	require.IsType(h.t, wwr.ReqErr{}, err)

	reqErr := err.(wwr.ReqErr)
	require.Equal(h.t, "UNAUTHORIZED", reqErr.Code)
	require.True(h.t, len(reqErr.Message) > 1)
}
