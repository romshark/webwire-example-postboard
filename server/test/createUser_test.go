package test

import (
	"testing"

	"github.com/qbeon/webwire-messenger/server/apisrv/api"
	"github.com/qbeon/webwire-messenger/server/test/setup"
)

// TestCreateUser tests user account creation
// by a root administrator client
func TestCreateUser(t *testing.T) {
	t.Parallel()
	ts := setup.New(t, stats)
	defer ts.Teardown()

	root := ts.NewAdminClient("root", "root")

	ts.Helper.CreateUser(
		root,
		api.CreateUserParams{
			FirstName: "testFirstName",
			LastName:  "testLastName",
			Username:  "testUsername",
			Password:  "testPassword",
			Type:      api.UtUser,
		},
	)
}
