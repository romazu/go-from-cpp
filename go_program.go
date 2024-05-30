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

type Gateway struct {
	Counter  int
	Running  bool
	StopChan chan bool
}

var (
	mu       sync.Mutex
	gateways = make(map[int]*Gateway)
	nextID   = 0
)

//export NewGateway
func NewGateway() C.int {
	mu.Lock()
	defer mu.Unlock()
	stopChan := make(chan bool)
	g := &Gateway{Counter: 0, Running: true, StopChan: stopChan}
	gateways[nextID] = g
	id := C.int(nextID)
	nextID++
	go runGateway(g) // Start the goroutine for this gateway
	return id
}

func runGateway(g *Gateway) {
	for {
		select {
		case <-g.StopChan:
			g.Running = false
			return
		default:
			time.Sleep(3 * time.Second)
			fmt.Printf("Gateway active: now = %v\n", time.Now().UTC().Format(time.RFC3339Nano))
		}
	}
}

//export StopGateway
func StopGateway(id C.int) {
	mu.Lock()
	g, ok := gateways[int(id)]
	mu.Unlock()

	if ok && g.Running {
		g.StopChan <- true
	}
}

//export SayHello
func SayHello(id C.int, name *C.char) {
	mu.Lock()
	g, ok := gateways[int(id)]
	mu.Unlock()

	if !ok {
		fmt.Println("Gateway not found")
		return
	}

	g.Counter++
	fmt.Printf("Hello, %s, counter: %v\n", C.GoString(name), g.Counter)
}

func main() {
	select {} // Keep the application running
}
