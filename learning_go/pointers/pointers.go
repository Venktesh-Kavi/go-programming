package main

import (
	"fmt"
)

func main() {
	t := []int{2, 4, 5}
	spr := &t // creates a pointer reference to the slice descriptor
	fmt.Printf("spr points to: %p (slice descriptor)\n", spr)
	fmt.Printf("slice descriptor points to: %p (backing array)\n", t)

	// Example of using New
	f := new(Foo) // allocates memory for type foo with zero values of the internal types.
	fmt.Printf("foo values: %v\n", f)
}

type Foo struct {
	name string
	age  int
}
