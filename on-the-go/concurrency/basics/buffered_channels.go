package main

import (
	"fmt"
	"math/rand"
)

/*
Buffered channels contain an internal value queue with fixed capacity. They don't block on sending/receiver till the cap is reached.
*/
func main() {
	ch := make(chan int, 10)

	go func() {
		for range 5 {
			val := rand.Intn(100)
			ch <- val // produce random values 5 times.
		}
		close(ch) // close of the channel is required, otherwise the receiver will get blocked forever resulting in deadlock.
	}()

	for val := range ch {
		fmt.Printf("received: %8d\n", val)
	}
}
