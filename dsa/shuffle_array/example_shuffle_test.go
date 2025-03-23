package shuffle_array

import "fmt"

func ExampleShuffle() {
	nums, _ := Shuffle([]int{1, 2, 3, 4, 5})
	fmt.Print(nums)
	// Output:
	// [1 2 3 4 5]
}
