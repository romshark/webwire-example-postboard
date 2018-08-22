package setup

import (
	"testing"
	"time"

	"github.com/qbeon/webwire-example-postboard/server/apisrv"
	"github.com/stretchr/testify/assert"
)

// Helper represents a test helper helping perform routine testing tasks
type Helper struct {
	t         *testing.T
	ts        *TestSetup
	apiServer apisrv.ApiServer
}

// Time returns a time.Time object parsed from the given RFC3339 string
// like: "2014-11-12T11:45:26.371Z"
func (h *Helper) Time(str string) time.Time {
	tm, err := time.Parse(time.RFC3339, str)
	assert.NoError(h.t, err)
	return tm
}
