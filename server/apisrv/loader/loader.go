package loader

import (
	"time"

	engiface "github.com/qbeon/webwire-messenger/server/apisrv/engine"
)

// batchedLoader represents an implementation of the Loader interface
// using the batching technique to reduce database roundtrips.
// It batches incoming requests until either the batch interval
// triggers or the batch limit is reached
type batchedLoader struct {
	interval time.Duration
	limit    uint32
}

// NewBatchedLoader returns a batched loader
func NewBatchedLoader(
	interval time.Duration,
	limit uint32,
) Loader {
	return &batchedLoader{
		interval: interval,
		limit:    limit,
	}
}

// LoadMessage implements the Loader interface
func (ldr *batchedLoader) LoadMessage(
	identifier engiface.Identifier,
) (*engiface.Message, error) {
	// TODO: implement loader
	return nil, nil
}
