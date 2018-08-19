package dam

import "sync"

// Dam represents a goroutine dam
type Dam interface {
	// Flush flushes all channels and return the number of flushed channels
	Flush(err error) uint32

	// Wait blocks until the dam is flushed
	Wait() error

	// Chan returns a channel that's written to when the dam is flushed
	Chan() chan error
}

// New constructs a new dam with the given preallocated capacity
func New(cap int) Dam {
	return &dam{
		lock:     sync.Mutex{},
		channels: make([]chan error, 0, cap),
	}
}

// dam represents an implementation of the Dam interface
type dam struct {
	lock     sync.Mutex
	channels []chan error
}

// add adds a new channel to the dam and returns it
func (d *dam) add() chan error {
	ch := make(chan error, 1)
	d.lock.Lock()
	d.channels = append(d.channels, ch)
	d.lock.Unlock()
	return ch
}

// Flush implements the Dam interface
func (d *dam) Flush(err error) uint32 {
	d.lock.Lock()
	flushed := uint32(len(d.channels))
	for _, ch := range d.channels {
		ch <- err
	}
	d.lock.Unlock()
	return flushed
}

// Wait implements the Dam interface
func (d *dam) Wait() error {
	return <-d.add()
}

// Chan implements the Dam interface
func (d *dam) Chan() chan error {
	return d.add()
}
