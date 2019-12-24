package question_two

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	TestIsomorphicPass1Name = "isomorphic words"
	TestIsomorphicPass2Name = "not isomorphic words 1"
	TestIsomorphicPass3Name = "not isomorphic words 2"
	TestIsomorphicPass4Name = "long word"
	TestIsomorphicInput1    = "paper"
	TestIsomorphicInput2    = "title"
	TestIsomorphicInput3    = "jeans"
	TestIsomorphicInput4    = "longword"
)

// TestIsomorphic is testing the public function "Isomorphic"
func TestIsomorphic(t *testing.T) {
	Isomorph := NewIsomorpher()

	// This is pass with isomorphic words
	t.Run(TestIsomorphicPass1Name, func(t *testing.T) {
		got := Isomorph.Isomorphic(TestIsomorphicInput1, TestIsomorphicInput2)
		want := true

		assert.EqualValues(t, got, want)
	})

	// Pass with not isomorphic words
	t.Run(TestIsomorphicPass2Name, func(t *testing.T) {
		got := Isomorph.Isomorphic(TestIsomorphicInput1, TestIsomorphicInput3)
		want := false

		assert.EqualValues(t, got, want)
	})

	// Pass with not isomorphic words
	t.Run(TestIsomorphicPass3Name, func(t *testing.T) {
		got := Isomorph.Isomorphic(TestIsomorphicInput3, TestIsomorphicInput2)
		want := false

		assert.EqualValues(t, got, want)
	})

	// Pass with too long a word
	t.Run(TestIsomorphicPass4Name, func(t *testing.T) {
		got := Isomorph.Isomorphic(TestIsomorphicInput1, TestIsomorphicInput4)
		want := false

		assert.EqualValues(t, got, want)
	})
}
