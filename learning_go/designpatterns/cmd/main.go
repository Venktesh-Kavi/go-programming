package main

import (
	"dsp/decorator"
	"dsp/dependency_injection"
	"fmt"
	"os"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		decorator.InitDecoratorServer(&wg)
	}()
	fmt.Printf("started server, lets Gooo!, FYI main thread is unblocked \n")

	dependencyInjectionDriver()
	wg.Wait()
}

func dependencyInjectionDriver() {
	ds := DefaultSleeper{}
	dependency_injection.Counter(os.Stdout, ds)
}

type DefaultSleeper struct{}

func (d DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}
