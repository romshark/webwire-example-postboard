package test

import (
	"flag"
	"os"
	"testing"
	"time"

	"github.com/qbeon/webwire-messenger/server/apisrv/config"
	"github.com/qbeon/webwire-messenger/server/client"
	"github.com/qbeon/webwire-messenger/server/test/setup"
)

var printStats = flag.Bool(
	"stats",
	false,
	"Toggles the printing of statistics at the end of the test",
)

// maxCreationTimeDeviation defines the
var maxCreationTimeDeviation = flag.Int64(
	"conf-max-crt-time-dev",
	4000,
	"Maximum entity creation-time deviation in milliseconds",
)

// setupConf represents the default global test setup configuration
var setupConf *setup.Config

// TestMain runs the API tests and computes & prints the statistics
func TestMain(m *testing.M) {
	// Initialize test statistics recorder
	statisticsRecorder := setup.NewStatisticsRecorder()

	// Initialize test setup configuration
	setupConf = setup.NewConfig(
		// Test statistics recorder
		statisticsRecorder,

		// Default maximum entity creation time deviation
		time.Duration(*maxCreationTimeDeviation)*time.Millisecond,

		// Server configuration
		config.Config{
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
		},
	)

	// Run the tests
	exitCode := m.Run()

	// Compute statistics
	computedStats := statisticsRecorder.Compute()

	// Print statistics
	if *printStats {
		computedStats.Print()
	}

	os.Exit(exitCode)
}

// Readers represents a list of API clients
type Readers []client.ApiClient
