package main

import (
	"fmt"
	"testing"
)

func TestRemoveElement(t *testing.T) {

	t.Run("provided nums & returned should have different slice descriptors", func(t *testing.T) {
		nums := []int{1, 2, 3, 4, 5}
		pp := &nums
		idx := 2
		RemoveElement(nums, idx)
		cp := &nums
		if pp == cp {
			t.Errorf("expected different pointers, got %v", cp)
		}
	})
}

func TestRemoveElementWithoutModification(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		nums := []int{1, 2, 3, 4, 5}
		idx := 2
		nn := RemoveElementWithoutModification(nums, idx)
		if len(nn) == len(nums) {
			t.Errorf("expected length to be %d, got %d", len(nums)-1, len(nn))
		}
		fmt.Println(nn)
		fmt.Println(nums)
	})
}

func TestRemoveElementWithoutMod(t *testing.T) {
	t.Run("happy path better approach", func(t *testing.T) {
		nums := []int{1, 2, 3, 4, 5}
		idx := 2
		nn := RemoveElementWithoutMod(nums, idx)
		if len(nn) == len(nums) {
			t.Errorf("expected length to be %d, got %d", len(nums)-1, len(nn))
		}
		fmt.Println(nn)
		fmt.Println(nums)
	})
}
