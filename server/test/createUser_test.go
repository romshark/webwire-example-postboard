package test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/qbeon/webwire-messenger/server/apisrv/api"
	"github.com/qbeon/webwire-messenger/server/test/setup"
)

// TestCreateUser tests user account creation
func TestCreateUser(t *testing.T) {
	ts := setup.New(t, stats)
	defer ts.Teardown()

	admin := ts.NewAdminClient("root", "root")

	newUserIdent, err := admin.CreateUser(
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

	user := ts.NewUserClient("testUsername", "testPassword")

	// Verify returned identifier
	session := user.Session()
	require.NotNil(t, session)
	require.Equal(t, newUserIdent, session.Info.Value("id").(api.Identifier))
}
