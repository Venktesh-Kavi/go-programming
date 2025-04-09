package main

import "testing"

func TestSubsetEqualSum(t *testing.T) {
	tc := []struct {
		nums []int
		want bool
	}{
		{
			nums: []int{3, 2, 1, 4},
			want: true,
		},
		{
			nums: []int{2, 3},
			want: false,
		},
	}

	for _, tt := range tc {
		if SubsetEqualSum(tt.nums) != tt.want {
			t.Errorf("SubsetEqualSum(%v) = %v, want %v", tt.nums, SubsetEqualSum(tt.nums), tt.want)
		}
	}
}
