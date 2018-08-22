package dam_test

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/qbeon/webwire-example-postboard/server/apisrv/dam"
)

// Tests the Wait method
func TestWait(t *testing.T) {
	numWaiters := 4
	d := dam.New(numWaiters)
	wg := sync.WaitGroup{}
	wg.Add(numWaiters)

	for i := 0; i < numWaiters; i++ {
		go func() {
			if err := d.Wait(); err != nil {
				t.Errorf("unexpected wait error: %s", err)
			}
			wg.Done()
		}()
	}
	time.Sleep(10 * time.Millisecond)

	if numFlushed := d.Flush(nil); numFlushed != uint32(numWaiters) {
		t.Fatalf("unexpected flush num: %d / %d", numFlushed, numWaiters)
	}
	wg.Wait()
}

// Tests the Chan method
func TestChan(t *testing.T) {
	numWaiters := 4
	d := dam.New(numWaiters)
	wg := sync.WaitGroup{}
	wg.Add(numWaiters)

	for i := 0; i < numWaiters; i++ {
		go func() {
			if err := <-d.Chan(); err != nil {
				t.Errorf("unexpected error: %s", err)
			}
			wg.Done()
		}()
	}
	time.Sleep(10 * time.Millisecond)

	if numFlushed := d.Flush(nil); numFlushed != uint32(numWaiters) {
		t.Fatalf("unexpected flush num: %d / %d", numFlushed, numWaiters)
	}
	wg.Wait()
}

// TestWaitError tests the Wait method when flushed with an error
func TestWaitError(t *testing.T) {
	numWaiters := 4
	d := dam.New(numWaiters)
	wg := sync.WaitGroup{}
	wg.Add(numWaiters)

	for i := 0; i < numWaiters; i++ {
		go func() {
			err := d.Wait()
			if err == nil {
				t.Errorf("missing wait error")
			}
			if err.Error() != "sample error" {
				t.Errorf("unexpected wait error: %s", err)
			}
			wg.Done()
		}()
	}
	time.Sleep(10 * time.Millisecond)

	if numFlushed := d.Flush(
		fmt.Errorf("sample error"),
	); numFlushed != uint32(numWaiters) {
		t.Fatalf("unexpected flush num: %d / %d", numFlushed, numWaiters)
	}
	wg.Wait()
}

// TestChanError tests the Chan method when flushed with an error
func TestChanError(t *testing.T) {
	numWaiters := 4
	d := dam.New(numWaiters)
	wg := sync.WaitGroup{}
	wg.Add(numWaiters)

	for i := 0; i < numWaiters; i++ {
		go func() {
			err := <-d.Chan()
			if err == nil {
				t.Errorf("missing error")
			}
			if err.Error() != "sample error" {
				t.Errorf("unexpected error: %s", err)
			}
			wg.Done()
		}()
	}
	time.Sleep(10 * time.Millisecond)

	if numFlushed := d.Flush(
		fmt.Errorf("sample error"),
	); numFlushed != uint32(numWaiters) {
		t.Fatalf("unexpected flush num: %d / %d", numFlushed, numWaiters)
	}
	wg.Wait()
}
