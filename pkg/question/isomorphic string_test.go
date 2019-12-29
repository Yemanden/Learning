package question_two

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	testIsomorphicPass1Name  = "isomorphic words"
	testIsomorphicPass1Input = "paper"
	testIsomorphicPass1Want  = true

	testIsomorphicPass2Name  = "not isomorphic words 1"
	testIsomorphicPass2Input = "title"
	testIsomorphicPass2Want  = false

	testIsomorphicPass3Name  = "not isomorphic words 2"
	testIsomorphicPass3Input = "jeans"
	testIsomorphicPass3Want  = false

	testIsomorphicPass4Name  = "long word"
	testIsomorphicPass4Input = "longword"
	testIsomorphicPass4Want  = false
)

// TestIsomorphic is testing the public function "Isomorphic"
func TestIsomorphic(t *testing.T) {
	isomorph := NewIsomorpher()

	// This is pass with isomorphic words
	t.Run(testIsomorphicPass1Name, func(t *testing.T) {
		got := isomorph.IsIsomorphic(testIsomorphicPass1Input, testIsomorphicPass2Input)
		want := testIsomorphicPass1Want

		assert.EqualValues(t, got, want)
	})

	// Pass with not isomorphic words
	t.Run(testIsomorphicPass2Name, func(t *testing.T) {
		got := isomorph.IsIsomorphic(testIsomorphicPass1Input, testIsomorphicPass3Input)
		want := testIsomorphicPass2Want

		assert.EqualValues(t, got, want)
	})

	// Pass with not isomorphic words
	t.Run(testIsomorphicPass3Name, func(t *testing.T) {
		got := isomorph.IsIsomorphic(testIsomorphicPass3Input, testIsomorphicPass2Input)
		want := testIsomorphicPass3Want

		assert.EqualValues(t, got, want)
	})

	// Pass with too long a word
	t.Run(testIsomorphicPass4Name, func(t *testing.T) {
		got := isomorph.IsIsomorphic(testIsomorphicPass1Input, testIsomorphicPass4Input)
		want := testIsomorphicPass4Want

		assert.EqualValues(t, got, want)
	})
}
