package shuffle_array

import (
	"errors"
	"math/rand"
)

func Shuffle(nums []int) ([]int, error) {
	if nums == nil {
		return nil, errors.New("nums is nil")
	}
	aux := make([]int, len(nums))
	copy(aux, nums)

	for i := 0; i < len(nums); i++ {
		ri := rand.Intn(len(nums))
		nums[i] = aux[ri]
		aux = append(aux[:ri], aux[ri+1:]...) // till ri-1 and ri+1 till end
	}
	return nums, nil
}
