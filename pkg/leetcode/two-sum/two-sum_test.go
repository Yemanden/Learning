package twosum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	testTwoNumbersPass1Name  = "Test with matching numbers"
	testTwoNumbersPass1Input = 3

	testTwoNumbersPass2Name  = "Test without matching numbers"
	testTwoNumbersPass2Input = 10
)

var (
	testTwoNumberPass2WantError = ErrInvalidNumbers
)

// TestTwoSum ...
func TestTwoSum(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}

	t.Run(testTwoNumbersPass1Name, func(t *testing.T) {
		want := []int{0, 1}
		got, err := NewTwoSummator().TwoSum(numbers, testTwoNumbersPass1Input)
		if assert.NoError(t, err) {
			assert.EqualValues(t, want, got)
		}
	})

	t.Run(testTwoNumbersPass2Name, func(t *testing.T) {
		var want []int = nil
		wantErr := testTwoNumberPass2WantError
		lc := NewTwoSummator()
		got, err := lc.TwoSum(numbers, testTwoNumbersPass2Input)
		assert.EqualValues(t, want, got)
		assert.EqualValues(t, wantErr, err)
	})
}
