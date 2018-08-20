package setup

import (
	"testing"

	"github.com/qbeon/webwire-messenger/server/apisrv/config"
)

// New creates a new test setup
func New(t *testing.T, statsRec *StatisticsRecorder) *TestSetup {
	// Load & rewrite configuration here

	// Create a new test setup with the following server configuration
	return newTestSetup(t, statsRec, config.Config{
		// Make the server automatically pick ports in testing mode
		ServerAddress:        "127.0.0.1:",
		MetricsServerAddress: "127.0.0.1:",

		// Disable TLS encryption in testing mode
		TLS: nil,

		// Log to console in testing mode
		Log: &config.LogConfig{
			// Disable debug logging during testing
			DebugEnabled: false,
		},
	})
}
