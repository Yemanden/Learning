package questionone

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	testTwoNumbersPass1Name = "Test with matching numbers"

	testTwoNumbersPass2Name = "Test without matching numbers"
)

var (
	testTwoNumberPass2WantError = errors.New("error: no matching numbers")
)

// TestTwoSum ...
func TestTwoSum(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}

	t.Run(testTwoNumbersPass1Name, func(t *testing.T) {
		lc := NewLeetCodeQuestion()
		got, err := lc.TwoSum(numbers, 3)
		want := []int{0, 1}
		wantErr := error(nil)
		assert.EqualValues(t, want, got)
		assert.EqualValues(t, wantErr, err)
	})
	t.Run(testTwoNumbersPass2Name, func(t *testing.T) {
		lc := NewLeetCodeQuestion()
		got, err := lc.TwoSum(numbers, 10)
		var want []int = nil
		wantErr := testTwoNumberPass2WantError
		assert.EqualValues(t, want, got)
		assert.EqualValues(t, wantErr, err)
	})
}
