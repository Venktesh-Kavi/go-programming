package islice

import "unsafe"

// Understanding how slice is implemented internally.

// Dynamically Resizing a slice in Go.

type Element struct {
	Value any
}

type ISlice[T any] interface{}
type SliceHeader struct {
	data unsafe.Pointer
	cap  int
	len  int
}
