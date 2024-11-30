package slices

// take an array of numbers and return the total.
func SumSlice(nums []int) int {
	var total int
	for _, num := range nums {
		total += num
	}
	return total
}

func SumAll(nl ...[]int) []int {
	var sums []int
	for _, nums := range nl {
		is := SumSlice(nums)
		sums = append(sums, is)
	}
	return sums
}
