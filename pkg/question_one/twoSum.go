package question_one

import "errors"

// TwoSummator ...
type TwoSummator interface {
	TwoSum([]int, int) ([]int, error)
}

type leetCodeQuestion struct{}

// TwoSum ...
func (l *leetCodeQuestion) TwoSum(nums []int, target int) ([]int, error) {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		m[nums[i]] = i
	}
	for i, value := range nums {
		if m[target-value] != i && m[target-value] != 0 {
			return []int{i, m[target-value]}, nil
		}
	}
	return nil, errors.New("error: no matching numbers")
}

// NewLeetCodeQuestion ...
func NewLeetCodeQuestion() TwoSummator {
	return &leetCodeQuestion{}
}
