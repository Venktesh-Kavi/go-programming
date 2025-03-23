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

func RemoveElement(nums []int, idx int) error {
	if !(idx < len(nums)) {
		return fmt.Errorf("index out of range")
	}

	// removing the element, approach 1
	// In the below approach it results in creation of new slice and attaching it to nums. Also the length is modified.
	nums = append(nums[:idx], nums[idx+1:]...)
	return nil
}

func RemoveElementWithoutModification(nums []int, idx int) []int {
	new_arr := make([]int, (len(nums) - 1))
	k := 0
	for i := 0; i < (len(nums) - 1); {
		if i != idx {
			new_arr[i] = nums[k]
			k++
		} else {
			k++
		}
		i++
	}

	return new_arr
}

func RemoveElementWithoutMod(nums []int, idx int) []int {
	ret := make([]int, 0)
	ret = append(ret, nums[:idx]...)
	return append(ret, nums[idx+1:]...)
}
