package main

import "fmt"

func main() {

	// converting a rune to a string
	// rune is an alias to int32, byte is an alias to uint32.
	// converting a signed or unsigned integer to a string, yields a string containing an UTF-8 representation of an integer.

	r := '{'
	rl := 'a'
	fmt.Printf("rune representation: %v, string representation: %s or while using format statement use: %c\n", r, string(r), r)
	fmt.Println(rl)
}
