package main

import (
	"dsa/stack"
	"fmt"
)

func main() {
	// stack ops
	s := stack.New[int]()
	s.Push(10)
	v, err := s.Pop()
	if err != nil {
		panic(err)
	}
	fmt.Printf("size of stack: %d, head: %d", s.Len(), v)
}
