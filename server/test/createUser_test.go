package test

import (
	"testing"

	"github.com/qbeon/webwire-example-postboard/server/apisrv/api"
	"github.com/qbeon/webwire-example-postboard/server/test/setup"
)

// TestCreateUser tests user account creation
// by a root administrator client
func TestCreateUser(t *testing.T) {
	t.Parallel()
	ts := setup.New(t, setupConf)
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
