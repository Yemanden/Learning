package Question

import (
	"fmt"
	"testing"
)

func TestIsomorphic(t *testing.T) {
	s1 := "paper"
	s2 := "title"
	s3 := "jeans"
	s4 := "longword"

	t.Run("isomorph words", func(t *testing.T) {
		got := Isomorphic(s1, s2)
		want := true

		if got != true {
			fmt.Printf("%s: got [%t] want [%t] \n", "isomorph words", got, want)
		}
	})

	t.Run("no isomorph words 1", func(t *testing.T) {
		got := Isomorphic(s1, s3)
		want := true

		if got != true {
			fmt.Printf("%s: got [%t] want [%t] \n", "no isomorph words 1", got, want)
		}
	})

	t.Run("no isomorph words 2", func(t *testing.T) {
		got := Isomorphic(s3, s2)
		want := true

		if got != true {
			fmt.Printf("%s: got [%t] want [%t] \n", "no isomorph words 2", got, want)
		}
	})

	t.Run("long word", func(t *testing.T) {
		got := Isomorphic(s1, s4)
		want := true

		if got != true {
			fmt.Printf("%s: got [%t] want [%t] \n", "long word", got, want)
		}
	})
}
