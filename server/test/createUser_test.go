package test

import (
	"testing"

	"github.com/qbeon/webwire-messenger/server/apisrv/api"
	"github.com/qbeon/webwire-messenger/server/test/setup"
)

// TestCreateUser tests user account creation
func TestCreateUser(t *testing.T) {
	t.Parallel()
	ts := setup.New(t, stats)
	defer ts.Teardown()

	admin := ts.NewAdminClient("root", "root")

	ts.Helper.CreateUser(
		admin,
		api.CreateUserParams{
			FirstName: "testFirstName",
			LastName:  "testLastName",
			Username:  "testUsername",
			Password:  "testPassword",
			Type:      api.UtUser,
		},
	)
}
