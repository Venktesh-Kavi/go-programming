package main

import (
	"fmt"
	"reflect"
)

func main() {

	// Zero value for float is 0 and it is considered as float64 type by  default during type inference.
	var zf float32
	f := 23.123 // float64

	tf := reflect.TypeOf(f)
	fmt.Println(tf)
	fmt.Printf("zero value of zf: %f\n", zf)
	fmt.Printf("Type of f is: %T", f)
}
