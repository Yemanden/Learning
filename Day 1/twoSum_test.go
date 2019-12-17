package Day_1

import (
	"reflect"
	"testing"
)

// TestTwoSum является тестом для функции twoSum
func TestTwoSum(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}

	got := twoSum(numbers, 3)
	want := []int{0, 1}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
