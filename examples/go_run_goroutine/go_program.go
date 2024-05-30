package main

/*
#include <stdlib.h>
*/
import "C"
import (
	"fmt"
	"sync"
	"time"
)

var (
	keepRunning bool
	lock        sync.Mutex
)

// Start a long-running goroutine that logs time every second.
//
//export StartGateway
func StartGateway() {
	lock.Lock()
	if keepRunning {
		lock.Unlock()
		return // Goroutine is already running
	}
	keepRunning = true
	lock.Unlock()

	go func() {
		for {
			lock.Lock()
			if !keepRunning {
				lock.Unlock()
				break
			}
			lock.Unlock()

			fmt.Println("Current time in Go:", time.Now().UTC().Format(time.RFC3339Nano))
			time.Sleep(1 * time.Second)
		}
	}()
}

// Stop the goroutine.
//
//export StopGateway
func StopGateway() {
	lock.Lock()
	keepRunning = false
	lock.Unlock()
}

func main() {
	// Main function is kept minimal.
	select {}
}
