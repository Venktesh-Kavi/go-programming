package main

import "fmt"

func main() {
	s := "foo"
	b := []byte(s)
	r := []rune(s)
	fmt.Printf("Byte: %v, len: %d\n", b, len(b))
	fmt.Printf("Rune: %v, len: %d\n", r, len(r))

	ns := "foo 世界"
	nb := []byte(ns)
	nr := []rune(ns)
	fmt.Printf("Byte: %v, len: %d\n", nb, len(nb))
	fmt.Printf("Rune: %v, len: %d\n", nr, len(nr))
}
