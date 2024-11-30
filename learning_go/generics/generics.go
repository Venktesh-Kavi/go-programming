package main

import (
	"fmt"
	"golang.org/x/exp/constraints"
	"math/rand"
)

func main() {
	s := []int{}
	ss := []string{"foo", "bar", "ada"}
	const n = 10
	for i := 0; i < 10; i++ {
		s = append(s, rand.Intn(n))
	}

	fmt.Printf("Unsorted s: %v\n", s)
	sr := GSliceSort(s) // sort an int slice
	fmt.Println(sr)

	fmt.Printf("Unsorted sf: %v\n", ss)
	sfr := GSliceSort(ss)
	fmt.Println(sfr)
}

// brute force sorting for a slice, internally sorting is based on pdq quicksort
func GSliceSort[T constraints.Ordered](slice []T) []T {
	res := make([]T, len(slice))
	copy(res, slice)
	for i := 0; i < len(slice); i++ {
		for j := i + 1; j < len(slice); j++ {
			if res[i] > res[j] {
				res[i], res[j] = res[j], res[i]
			}
		}
	}
	return res
}
