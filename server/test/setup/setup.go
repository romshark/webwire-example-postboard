package setup

import (
	"testing"
	"time"

	"github.com/qbeon/webwire-example-postboard/server/apisrv/config"
)

// Config represents the test setup configuration
type Config struct {
	statisticsRecorder              *StatisticsRecorder
	defaultMaxCreationTimeDeviation time.Duration
	serverConfig                    config.Config
}

// NewConfig creates a new test setup configuration
func NewConfig(
	statisticsRecorder *StatisticsRecorder,
	defaultMaxCreationTimeDeviation time.Duration,
	serverConfig config.Config,
) *Config {
	return &Config{
		statisticsRecorder:              statisticsRecorder,
		defaultMaxCreationTimeDeviation: defaultMaxCreationTimeDeviation,
		serverConfig:                    serverConfig,
	}
}

// New creates a new test setup
func New(t *testing.T, conf *Config) *TestSetup {
	// Load & rewrite configuration here

	// Create a new test setup with the following server configuration
	return newTestSetup(t, conf)
}
