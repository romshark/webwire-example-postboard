package test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/qbeon/webwire-messenger/server/apisrv/api"
	"github.com/qbeon/webwire-messenger/server/test/setup"
)

// TestLoginRoot tests login to root account
func TestLoginRoot(t *testing.T) {
	ts := setup.New(t, stats)
	defer ts.Teardown()

	clt := ts.NewGuestClient()
	require.NoError(t, clt.Login(context.Background(), api.LoginParams{
		Username: "root",
		Password: "root",
	}))
}
