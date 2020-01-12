package questiontwo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	testIsIsomorphicPass1Name  = "isomorphic words"
	testIsIsomorphicPass1Input = "paper"
	testIsIsomorphicPass1Want  = true

	testIsIsomorphicPass2Name  = "not isomorphic words 1"
	testIsIsomorphicPass2Input = "title"
	testIsIsomorphicPass2Want  = false

	testIsIsomorphicPass3Name  = "not isomorphic words 2"
	testIsIsomorphicPass3Input = "jeans"
	testIsIsomorphicPass3Want  = false

	testIsIsomorphicPass4Name  = "long word"
	testIsIsomorphicPass4Input = "longword"
	testIsIsomorphicPass4Want  = false
)

// TestIsomorphic is testing the public function "IsIsomorphic"
func TestIsIsomorphic(t *testing.T) {
	isomorph := NewIsomorpher()

	// This is pass with isomorphic words
	t.Run(testIsIsomorphicPass1Name, func(t *testing.T) {
		got := isomorph.IsIsomorphic(testIsIsomorphicPass1Input, testIsIsomorphicPass2Input)
		want := testIsIsomorphicPass1Want

		assert.EqualValues(t, got, want)
	})

	// Pass with not isomorphic words
	t.Run(testIsIsomorphicPass2Name, func(t *testing.T) {
		got := isomorph.IsIsomorphic(testIsIsomorphicPass1Input, testIsIsomorphicPass3Input)
		want := testIsIsomorphicPass2Want

		assert.EqualValues(t, got, want)
	})

	// Pass with not isomorphic words
	t.Run(testIsIsomorphicPass3Name, func(t *testing.T) {
		got := isomorph.IsIsomorphic(testIsIsomorphicPass3Input, testIsIsomorphicPass2Input)
		want := testIsIsomorphicPass3Want

		assert.EqualValues(t, got, want)
	})

	// Pass with too long a word
	t.Run(testIsIsomorphicPass4Name, func(t *testing.T) {
		got := isomorph.IsIsomorphic(testIsIsomorphicPass1Input, testIsIsomorphicPass4Input)
		want := testIsIsomorphicPass4Want

		assert.EqualValues(t, got, want)
	})
}
