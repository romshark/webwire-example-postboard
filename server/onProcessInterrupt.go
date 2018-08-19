package main

import (
	"os"
	"os/signal"
)

// onProcessInterrupt executes the given callback when an
// OS interruption signal is received by this process
func onProcessInterrupt(callback func()) {
	// Setup termination signal listener
	go func() {
		stop := make(chan os.Signal, 1)
		signal.Notify(stop, os.Interrupt)
		<-stop
		callback()
	}()
}
