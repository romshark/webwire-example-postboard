package helper

import (
	"testing"
	"time"

	"github.com/qbeon/webwire-messenger/server/apisrv"
	"github.com/stretchr/testify/assert"
)

// Helper represents a test helper
type Helper struct {
	t         *testing.T
	apiServer apisrv.ApiServer
}

// Initialize initializes the helper singleton
func New(t *testing.T, apiServer apisrv.ApiServer) *Helper {
	return &Helper{
		t,
		apiServer,
	}
}

// Time returns a time.Time object parsed from the given RFC3339 string
// like: "2014-11-12T11:45:26.371Z"
func (h *Helper) Time(str string) time.Time {
	tm, err := time.Parse(time.RFC3339, str)
	assert.NoError(h.t, err)
	return tm
}
