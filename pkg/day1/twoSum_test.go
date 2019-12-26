package question_one

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	twoNumbersTest1Name = "Test 1"
)

// TestTwoSum является тестом для функции twoSum
func TestTwoSum(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}

	t.Run(twoNumbersTest1Name, func(t *testing.T) {
		got := twoSum(numbers, 3)
		want := []int{0, 1}
		if !assert.EqualValues(t, got, want) {
			t.Errorf("%s: got [%v] want [%v]", twoNumbersTest1Name, got, want)
		}
	})
}
