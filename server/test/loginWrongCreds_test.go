package test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/qbeon/webwire-messenger/server/apisrv/api"
	"github.com/qbeon/webwire-messenger/server/test/setup"
)

// TestLoginWrongCreds tests login with wrong credentials
func TestLoginWrongCreds(t *testing.T) {
	t.Run("WrongPassword", func(t *testing.T) {
		t.Parallel()
		ts := setup.New(t, setupConf)
		defer ts.Teardown()

		root := ts.NewAdminClient("root", "root")
		userProf, userPass, _ := ts.Helper.CreateUserRand(root, api.UtUser)
		adminProf, adminPass, _ := ts.Helper.CreateUserRand(root, api.UtAdmin)

		clt1 := ts.NewGuestClient()
		clt2 := ts.NewGuestClient()
		ts.Helper.VerifyWrongCredsErr(
			clt1.Login(context.Background(), api.LoginParams{
				Username: userProf.Username,
				Password: adminPass,
			}),
		)
		require.Nil(t, clt1.Session())
		require.NoError(t, clt1.Login(context.Background(), api.LoginParams{
			Username: userProf.Username,
			Password: userPass,
		}))

		ts.Helper.VerifyWrongCredsErr(
			clt2.Login(context.Background(), api.LoginParams{
				Username: adminProf.Username,
				Password: userPass,
			}),
		)
		require.Nil(t, clt2.Session())
		require.NoError(t, clt2.Login(context.Background(), api.LoginParams{
			Username: adminProf.Username,
			Password: adminPass,
		}))
	})
}
