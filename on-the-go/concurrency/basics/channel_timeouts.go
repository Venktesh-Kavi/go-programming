package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	select {
	case msg := <-ch: // read from ch blocks forever.
		fmt.Println(msg)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout, saved from a channel deadlock")
	}
}
