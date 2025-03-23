package shuffle_array

import (
	"errors"
	"reflect"
	"testing"
)

func TestShuffle(t *testing.T) {
	subtests := []struct {
		name        string
		nums        []int
		expectedErr error
	}{
		{
			name: "happy path shuffling",
			nums: []int{1, 2, 3, 4, 5},
		},
		{
			name: "empty array",
			nums: []int{},
		},
		{
			name: "single element array",
			nums: []int{1},
		},
		{
			name:        "nil array",
			nums:        nil,
			expectedErr: errors.New("nums is nil"),
		},
	}

	for _, subtest := range subtests {
		t.Run(subtest.name, func(t *testing.T) {
			nums, err := Shuffle(subtest.nums)
			// compare two slices
			// https://stackoverflow.com/questions/15311969/checking-the-equality-of-two-slices
			if reflect.DeepEqual(subtest.nums, nums) {
				t.Errorf("Shuffle() error = %v, wantErr %v", err, subtest.expectedErr)
			}
			if !errors.Is(err, subtest.expectedErr) {
				t.Errorf("Shuffle() error = %v, wantErr %v", err, subtest.expectedErr)
			}
		})
	}
}
