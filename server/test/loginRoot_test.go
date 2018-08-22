package test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/qbeon/webwire-example-postboard/server/apisrv/api"
	"github.com/qbeon/webwire-example-postboard/server/test/setup"
)

// TestLogin tests login
func TestLogin(t *testing.T) {
	t.Run("AsRoot", func(t *testing.T) {
		t.Parallel()
		ts := setup.New(t, setupConf)
		defer ts.Teardown()

		clt := ts.NewGuestClient()
		require.NoError(t, clt.Login(context.Background(), api.LoginParams{
			Username: "root",
			Password: "root",
		}))
	})
}
