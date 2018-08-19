package metrics

import (
	"expvar"
	"time"
)

// met represents the global metrics state
var met struct {
	requests          *expvar.Int
	requestLatencyAvg *expvar.Int
	requestLatencyMed *expvar.Int
	requestLatencyMax *expvar.Int
	requestLatencyMin *expvar.Int
	connections       *expvar.Int
	sessions          *expvar.Int
	sessionLookups    *expvar.Int
}

// init automatically initializes the global metrics
func init() {
	met.requests = expvar.NewInt("Requests")
	met.requestLatencyAvg = expvar.NewInt("RequestLatencyAvg")
	met.requestLatencyMed = expvar.NewInt("RequestLatencyMed")
	met.requestLatencyMax = expvar.NewInt("RequestLatencyMax")
	met.requestLatencyMin = expvar.NewInt("RequestLatencyMin")
	met.connections = expvar.NewInt("Connection")
	met.sessions = expvar.NewInt("Sessions")
	met.sessionLookups = expvar.NewInt("SessionLookups")
	met.requestLatencyAvg.Set(0)
	met.requestLatencyMed.Set(0)
	met.requestLatencyMax.Set(0)
	met.requestLatencyMin.Set(0)
	met.connections.Set(0)
	met.sessions.Set(0)
	met.sessionLookups.Set(0)
}

// Request logs a request
func Request() {
	met.requests.Add(1)
}

// RequestCompleted logs the maximum, minimum, average and medium
// latency of completed requests
func RequestCompleted(ms time.Duration) {
	// Update average request latency
	met.requestLatencyAvg.Add(int64(ms / time.Millisecond))

	// Update minimum request latency
	minVal := int64(ms / time.Millisecond)
	if minVal < met.requestLatencyMin.Value() {
		met.requestLatencyMin.Set(minVal)
	}

	// Update maximum request latency
	maxVal := int64(ms / time.Millisecond)
	if maxVal > met.requestLatencyMax.Value() {
		met.requestLatencyMax.Set(maxVal)
	}

	// Update median request latency
	met.requestLatencyMed.Set(int64(ms / time.Millisecond))
}

// ConnectionCreated logs a webwire connection creation
func ConnectionCreated() {
	met.connections.Add(1)
}

// ConnectionClosed logs a webwire connection closure
func ConnectionClosed() {
	met.connections.Add(-1)
}

// SessionCreated logs a session creation
func SessionCreated() {
	met.sessions.Add(1)
}

// SessionClosed logs a session closure
func SessionClosed() {
	met.sessions.Add(-1)
}

// SessionLookup logs a session lookup
func SessionLookup() {
	met.sessionLookups.Add(1)
}
