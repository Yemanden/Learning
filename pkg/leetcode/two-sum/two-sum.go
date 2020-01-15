package twosum

import "github.com/Yemanden/Learning/pkg/models"

// TwoSummator ...
type TwoSummator interface {
	TwoSum([]int, int) ([]int, error)
}

type twoSummator struct{}

// TwoSum ...
func (l *twoSummator) TwoSum(nums []int, target int) ([]int, error) {
	m := make(map[int]int, len(nums))
	for i := 0; i < len(nums); i++ {
		m[nums[i]] = i + 1
		if m[target-nums[i]] != i+1 && m[target-nums[i]] != 0 {
			return []int{i - 1, m[target-nums[i]]}, nil
		}
	}
	return nil, models.InvalidNumbersError
}

// NewTwoSummator ...
func NewTwoSummator() TwoSummator {
	return &twoSummator{}
}
