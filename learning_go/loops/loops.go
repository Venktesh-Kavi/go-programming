package main

import "fmt"

// Different looping mechanisms in go
func main() {
	s := []int{2, 5, 6, 7}
	// for range loop
	fmt.Println("******* FOR RANGE *******")
	for i, e := range s {
		fmt.Printf("idx: %d, value: %d\n", i, e)
	}

	fmt.Println("****** Go's While ******* ")
	i := 5
	for i >= 0 {
		fmt.Printf("val: %d, ", i)
		i -= 1
	}
	fmt.Println()

	fmt.Println("******* Iterate over an INT (>=Go 1.22) **********")
	for i := range 10 {
		fmt.Printf("val: %d, ", i)
	}
	fmt.Println()

	fmt.Println("******** Go's Do While ******")
	m := map[string]int{
		"foo": 1,
	}
	for {
		fmt.Println(m)
		if _, ok := m["foo"]; ok {
			delete(m, "foo")
		}
		if len(m) == 0 {
			fmt.Println("map is empty breaking")
			break
		}
	}
}
