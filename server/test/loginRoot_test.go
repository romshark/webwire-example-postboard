package test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/qbeon/webwire-messenger/server/apisrv/api"
	"github.com/qbeon/webwire-messenger/server/test/setup"
)

// TestLoginRoot tests logging in to the root account created by default
func TestLoginRoot(t *testing.T) {
	t.Parallel()
	ts := setup.New(t, stats)
	defer ts.Teardown()

	clt := ts.NewGuestClient()
	require.NoError(t, clt.Login(context.Background(), api.LoginParams{
		Username: "root",
		Password: "root",
	}))
}
