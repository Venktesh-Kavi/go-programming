package main

import "fmt"

// RemoveElement removes an element from a slice
func ExampleRemoveElement() {
	nums := []int{1, 2, 3, 4, 5}
	RemoveElement(nums, 3)
	fmt.Println(nums)
	// Output:
	// [1 2 3 5 5]
}
