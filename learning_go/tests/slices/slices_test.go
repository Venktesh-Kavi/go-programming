package slices

import (
	"reflect"
	"testing"
)

func Test_SliceUsage(t *testing.T) {
	// summation of array
	t.Run("test success slice usage", func(t *testing.T) {
		nums := []int{5, 3, 4, 1}
		got := SumSlice(nums)
		expected := 13

		if got != expected {
			t.Errorf("got %d want %d", got, expected)
		}
	})

	t.Run("failure slice usage", func(t *testing.T) {
		nums := []int{3, 4, 2, 1}
		got := SumSlice(nums)
		expected := 9
		if got == expected {
			t.Errorf("got %d want %d", got, expected)
		}
	})
}

func Test_SumAll(t *testing.T) {

	t.Run("success sum all function", func(t *testing.T) {
		nums := []int{1, 2, 3}
		nums1 := []int{4, 5, 6}
		got := SumAll(nums, nums1)
		expected := []int{6, 15}

		if !reflect.DeepEqual(got, expected) {
			t.Errorf("got %v want %v", got, expected)
		}
	})
}

func TestSumTails(t *testing.T) {
	input := [][]int{{4, 5, 6}, {1, 2, 4}, {6, 8, 9}}
	expected := []int{11, 6, 17}
	got := SumTails(input)

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("got %v want %v", got, expected)
	}
}
