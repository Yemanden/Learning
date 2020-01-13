package questionnine

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	testReverseListPass1Name = "Reverse filled list"
	testReverseListPass2Name = "Reverse empty list"
)

// TestReverseList ...
func TestReverseList(t *testing.T) {
	l1 := NewListNode()
	l2 := NewListNode()
	l3 := NewListNode()
	list1 := NewListNode()

	l1.SetValue(4)
	l2.SetValue(3)
	l2.SetNextNode(l1)
	l3.SetValue(2)
	l3.SetNextNode(l2)
	list1.SetValue(1)
	list1.SetNextNode(l3)

	var list2 = NewListNode()
	list2.SetValue(1)

	t.Run(testReverseListPass1Name, func(t *testing.T) {
		tl1 := NewListNode()
		tl2 := NewListNode()
		tl3 := NewListNode()
		tl4 := NewListNode()

		tl1.SetValue(1)
		tl2.SetValue(2)
		tl2.SetNextNode(tl1)
		tl3.SetValue(3)
		tl3.SetNextNode(tl2)
		tl4.SetValue(4)
		tl4.SetNextNode(tl3)

		got := list1.Reverse().GetData()
		want := tl4.GetData()

		assert.EqualValues(t, got, want)
	})

	t.Run(testReverseListPass2Name, func(t *testing.T) {
		got := list2.Reverse()
		want := list2

		assert.EqualValues(t, got, want)
	})
}
