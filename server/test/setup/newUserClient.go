package setup

import (
	"context"

	"github.com/qbeon/webwire-messenger/server/apisrv/api"
	"github.com/qbeon/webwire-messenger/server/apisrv/sessinfo"
	"github.com/qbeon/webwire-messenger/server/client"
	"github.com/stretchr/testify/require"
)

// NewUserClient creates and connects a new regular user client
// verifying whether the connection was successfully established
// and whether the session is correct
func (ts *TestSetup) NewUserClient(
	username,
	password string,
) client.ApiClient {
	clt := ts.newClient()

	// Login
	require.NoError(ts.t, clt.Login(context.Background(), api.LoginParams{
		Username: username,
		Password: password,
	}))

	// Verify session key
	session := clt.Session()
	require.NotNil(ts.t, session, "expected a session, got nil")
	require.True(ts.t, len(session.Key) > 0, "unexpected session key")

	// Verify session info
	require.IsType(ts.t, &sessinfo.SessionInfo{}, session.Info)
	sessionInfo, ok := session.Info.(*sessinfo.SessionInfo)
	require.True(ts.t, ok, "unexpected session info object type")

	// Verify identifier
	require.NotEqual(ts.t, api.Identifier{}, sessionInfo.UserIdentifier)

	// Verify user type
	require.Equal(ts.t, api.UtUser, sessionInfo.UserType)

	return clt
}
