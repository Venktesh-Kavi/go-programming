package main

import (
	"dsp/decorator"
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		decorator.InitDecoratorServer(&wg)
	}()
	fmt.Printf("started server, lets Gooo!, FYI main thread is unblocked \n")
	wg.Wait()
}
