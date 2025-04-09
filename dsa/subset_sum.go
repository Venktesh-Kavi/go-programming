package main

import "fmt"

// SubsetEqualSum can be used to determine if a given array can be partitioned into two equal subsets.
// nums is the provided slice and isEqual is the result.
func SubsetEqualSum(nums []int) (isEqual bool) {

	p1 := make([]int, 1)
	p2 := make([]int, len(nums)-1)

	for i := 0; i < len(nums); i++ {
		p1 = append(p1, nums[i])
		p2 = append(p2, nums[i+1:]...)
	}
	fmt.Println(len(p1), len(p2))
	fmt.Println(p1)
	fmt.Println(p2)
	return true
}

/*




 */
