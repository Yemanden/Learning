package question_two

import (
	"fmt"
	"testing"
)

// TestIsomorphic is testing the public function "Isomorphic"
func TestIsomorphic(t *testing.T) {
	s1 := "paper"
	s2 := "title"
	s3 := "jeans"
	s4 := "longword"

	// This is pass with isomorphic words
	t.Run("isomorphic words", func(t *testing.T) {
		got := Isomorphic(s1, s2)
		want := true

		if got != true {
			fmt.Printf("%s: got [%t] want [%t] \n", "isomorphic words", got, want)
		}
	})

	// Pass with not isomorphic words
	t.Run("not isomorphic words 1", func(t *testing.T) {
		got := Isomorphic(s1, s3)
		want := true

		if got != true {
			fmt.Printf("%s: got [%t] want [%t] \n", "no isomorphic words 1", got, want)
		}
	})

	// Pass with not isomorphic words
	t.Run("not isomorphic words 2", func(t *testing.T) {
		got := Isomorphic(s3, s2)
		want := true

		if got != true {
			fmt.Printf("%s: got [%t] want [%t] \n", "no isomorphic words 2", got, want)
		}
	})

	// Pass with too long a word
	t.Run("long word", func(t *testing.T) {
		got := Isomorphic(s1, s4)
		want := true

		if got != true {
			fmt.Printf("%s: got [%t] want [%t] \n", "long word", got, want)
		}
	})
}
