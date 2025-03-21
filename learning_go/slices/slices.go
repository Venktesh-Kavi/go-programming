package main

import (
	"fmt"
	"math/rand"
	"os"
	"sort"
)

func main() {
	s := make([]int, 0, 3)  // make a slice with initial length as 0
	s1 := make([]int, 3, 5) // make a slice with an initial length and capacity. This creates a slice with 3 elements with init value as 0 and a capacity of 5
	s2 := []int{}           // declare and initialise an empty slice
	s3 := [...]int{2, 4, 5} // automatically allocated slice with the provided size.
	fmt.Println(s3)
	const n = 10
	// iterate 10 times
	for i := 0; i < n; i++ {
		s = append(s, i)
		s1 = append(s1, i)
		s2 = append(s2, rand.Intn(n))
	}
	err := SearchASlice(s2)
	if err != nil {
		os.Exit(1)
	}
}

func SearchASlice(slice []int) error {
	fv := sort.Search(len(slice), func(i int) bool {
		return slice[i] >= 2
	})

	if fv == 0 {
		return fmt.Errorf("unable to find value")
	}
	fmt.Println("Found value: ", fv)
	return nil
}
