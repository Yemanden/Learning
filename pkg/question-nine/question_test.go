package question_nine

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	TestReverseListPass1Name = "Reverse filled list"
	TestReverseListPass2Name = "Reverse empty list"
)

func TestReverseList(t *testing.T) {
	list1 := &listNode{1, &listNode{2, &listNode{3, &listNode{4, nil}}}}
	var list2 *listNode = nil

	t.Run(TestReverseListPass1Name, func(t *testing.T) {
		got := reverseList(list1)
		want := &listNode{4, &listNode{3, &listNode{2, &listNode{1, nil}}}}

		assert.EqualValues(t, got, want)
	})

	t.Run(TestReverseListPass2Name, func(t *testing.T) {
		got := reverseList(list2)
		want := list2

		assert.EqualValues(t, got, want)
	})
}
