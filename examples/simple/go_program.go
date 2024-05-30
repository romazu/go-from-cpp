package main

import (
	"C"
	"fmt"
)

//export SayHello
func SayHello(name *C.char) {
	fmt.Println("Hello, " + C.GoString(name))
}

func main() {
	// Prevent the program from exiting immediately
	select {}
}
