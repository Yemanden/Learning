package question_nine_test

import (
	"github.com/Yemanden/Learning/pkg/question-nine"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	TestReverseListPass1Name = "Reverse filled list"
	TestReverseListPass2Name = "Reverse empty list"
)

// TestReverseList ...
func TestReverseList(t *testing.T) {
	l1 := questionnine.NewListNode(4, nil)
	l2 := questionnine.NewListNode(3, l1)
	l3 := questionnine.NewListNode(2, l2)
	list1 := questionnine.NewListNode(1, l3)

	var list2 = questionnine.NewListNode(1, nil)

	t.Run(TestReverseListPass1Name, func(t *testing.T) {
		tl1 := questionnine.NewListNode(1, nil)
		tl2 := questionnine.NewListNode(2, tl1)
		tl3 := questionnine.NewListNode(3, tl2)
		tl4 := questionnine.NewListNode(4, tl3)

		got := list1.Reverse().GetData()
		want := tl4.GetData()

		assert.EqualValues(t, got, want)
	})

	t.Run(TestReverseListPass2Name, func(t *testing.T) {
		got := list2.Reverse()
		want := list2

		assert.EqualValues(t, got, want)
	})
}
