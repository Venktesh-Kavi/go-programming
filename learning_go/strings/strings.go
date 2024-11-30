package main

import "fmt"

func main() {
	s := "foo" // shorthand declaration + initialization
	ns := "bar"

	s = ns          // s point to the backing array of ns
	ns += "k"       // ns is bark now, since strings are immutable a new backing array is created.
	fmt.Println(s)  // s is still pointing to bar, the foo's backing array will be garbage collecting.
	fmt.Println(ns) // ns is pointing to the new backing array bark

	c := s[:2] // slice s, c points to the same memory address as s
	fmt.Printf("s points to %v\n", &c)
	fmt.Printf("s is pointing to the backing array at: %v", &s)
}
