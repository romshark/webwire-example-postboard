package setup

import (
	"log"
	"os"
	"time"

	wwrclt "github.com/qbeon/webwire-go/client"
	"github.com/qbeon/webwire-messenger/server/client"
	"github.com/stretchr/testify/require"
)

// newClient creates and connects a new client
// verifying whether the connection was successfully established
func (ts *TestSetup) newClient() client.ApiClient {
	// Initialize webwire API client
	clt := client.New(
		ts.ApiServer.Address().String(),
		wwrclt.Options{
			DefaultRequestTimeout: 5 * time.Second,
			ErrorLog: log.New(
				os.Stderr,
				"WWR_ERR: ",
				log.Ldate|log.Ltime|log.Lshortfile,
			),
			WarnLog: log.New(
				os.Stdout,
				"WWR_WARN: ",
				log.Ldate|log.Ltime|log.Lshortfile,
			),
		},
		client.Hooks{},
	)

	// Manually call connect to ensure the server is available
	require.NoError(ts.t, clt.Connect())

	// Remember client to properly close it in the end of the test
	ts.clients = append(ts.clients, clt)
	return clt
}
