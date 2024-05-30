package main

/*
#include <stdlib.h>
*/
import "C"
import (
	"fmt"
	"sync"
)

type Gateway struct {
	Counter int
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
	g := &Gateway{Counter: 0}
	gateways[nextID] = g
	id := C.int(nextID)
	nextID++
	return id
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
