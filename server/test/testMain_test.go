package test

import (
	"flag"
	"os"
	"testing"

	"github.com/qbeon/webwire-messenger/server/client"
	"github.com/qbeon/webwire-messenger/server/test/setup"
)

// stats represents the global statistics recorder the setups must use
var stats = setup.NewStatisticsRecorder()

var printStats = flag.Bool(
	"stats",
	false,
	"Toggles the printing of statistics at the end of the test",
)

// TestMain runs the API tests and computes & prints the statistics
func TestMain(m *testing.M) {
	// Run the tests
	exitCode := m.Run()

	// Compute statistics
	computedStats := stats.Compute()

	// Print statistics
	if *printStats {
		computedStats.Print()
	}

	os.Exit(exitCode)
}

// Readers represents a list of API clients
type Readers []client.ApiClient
