package helper

import (
	"encoding/json"

	wwr "github.com/qbeon/webwire-go"
	"github.com/stretchr/testify/require"
)

// VerifyUnauthErr helps ensuring that the given error is an unauth-error
func (h *Helper) VerifyUnauthErr(err error) {
	require.Error(h.t, err)
	require.IsType(h.t, wwr.ReqErr{}, err)

	reqErr := err.(wwr.ReqErr)
	require.Equal(h.t, "GRAPH_ERR", reqErr.Code)
	require.True(h.t, len(reqErr.Message) > 1)

	var errors []struct {
		Code    string `json:"c"`
		Message string `json:"m"`
	}
	require.NoError(h.t, json.Unmarshal([]byte(reqErr.Message), &errors))

	require.Len(h.t, errors, 1)
	require.Equal(h.t, "UNAUTHORIZED", errors[0].Code)
	require.True(h.t, len(errors[0].Message) > 0)
}
