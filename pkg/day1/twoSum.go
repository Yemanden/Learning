package question_one

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for key, elem := range nums {
		m[elem] = key
	}

	for i, value := range nums {
		if m[target-value] != i && m[target-value] != 0 {
			return []int{i, m[target-value]}
		}
	}
	return nil
}
