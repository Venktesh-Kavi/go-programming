package tstructs

import (
	"fmt"
	"time"
)

// Pointer & Value Method Receiver Examples

type Counter struct {
	total       int
	lastUpdated time.Time
}

// example to indicate that when c is passed as a value type, pointer receiver based method mutations only happen on the copy of the value type c.
func DoUpdateWrong(c Counter) {
	c.Increment()
}

func DoUpdateRight(c *Counter) {
	c.Increment()
}

func (c *Counter) Increment() *Counter {
	c.total += 1
	return c
}
func (c Counter) String() string {
	return fmt.Sprintf("Total: %d, Last Updated: %s", c.total, c.lastUpdated)
}
