package test

import (
	"context"
	"testing"

	"github.com/qbeon/webwire-example-postboard/server/apisrv/api"
	"github.com/qbeon/webwire-example-postboard/server/client"
	"github.com/qbeon/webwire-example-postboard/server/test/setup"
	"github.com/stretchr/testify/require"
)

// TestCreateUserPerm tests user account creation without proper permissions
func TestCreateUserPerm(t *testing.T) {
	tryCreateUser := func(
		t *testing.T,
		ts *setup.TestSetup,
		creator client.ApiClient,
	) {
		ident, err := creator.CreateUser(
			context.Background(),
			api.CreateUserParams{
				FirstName: "testFirstName",
				LastName:  "testLastName",
				Username:  "testUsername",
				Password:  "testPassword",
				Type:      api.UtUser,
			},
		)
		ts.Helper.VerifyUnauthErr(err)
		require.True(t, ident.IsNull())
	}

	tryCreateUserSuccess := func(
		t *testing.T,
		ts *setup.TestSetup,
		creator client.ApiClient,
	) {
		ident, err := creator.CreateUser(
			context.Background(),
			api.CreateUserParams{
				FirstName: "testFirstName",
				LastName:  "testLastName",
				Username:  "testUsername",
				Password:  "testPassword",
				Type:      api.UtUser,
			},
		)
		require.NoError(t, err)
		require.False(t, ident.IsNull())
	}

	t.Run("AsUser", func(t *testing.T) {
		t.Parallel()
		ts := setup.New(t, setupConf)
		defer ts.Teardown()

		root := ts.NewAdminClient("root", "root")
		_, _, user := ts.Helper.CreateUserRand(root, api.UtUser)

		tryCreateUser(t, ts, user)
	})

	t.Run("AsGuest", func(t *testing.T) {
		t.Parallel()
		ts := setup.New(t, setupConf)
		defer ts.Teardown()

		guest := ts.NewGuestClient()

		tryCreateUser(t, ts, guest)
	})

	t.Run("AsAdmin", func(t *testing.T) {
		t.Parallel()
		ts := setup.New(t, setupConf)
		defer ts.Teardown()

		root := ts.NewAdminClient("root", "root")
		_, _, admin := ts.Helper.CreateUserRand(root, api.UtAdmin)

		tryCreateUserSuccess(t, ts, admin)
	})
}
