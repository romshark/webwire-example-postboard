package setup

import (
	wwr "github.com/qbeon/webwire-go"
	"github.com/stretchr/testify/require"
)

// VerifyWrongCredsErr helps ensuring that the given error is
// a wrong-credentials error with an error message included
func (h *Helper) VerifyWrongCredsErr(err error) {
	require.Error(h.t, err)
	require.IsType(h.t, wwr.ReqErr{}, err)

	reqErr := err.(wwr.ReqErr)
	require.Equal(h.t, "WRONG_CREDENTIALS", reqErr.Code)
	require.True(h.t, len(reqErr.Message) > 1)
}
