package main

import "fmt"

/*
Ping Pong example, blocking channels are used to achieve synchronization between two go routines.
Unbuffered channels are by default blocking. The send channel blocks until there is receiver to receive and receiver blocks until there is a sender.

In the below example main thread and the go routine run concurrently, the sequencing can vary. We want pong to be printed after ping.
*/
func main() {
	ch := make(chan string)
	go func() {
		msg := <-ch // blocking receiver, waits till ping is produced, if this go routines runs before main.
		fmt.Println(msg)
		ch <- "pong"
	}()

	// produce pong
	ch <- "ping" // blocks until there is a receiver, the abv go routine setups a receiver unblocking ping.
	fmt.Println(<-ch)
}
