package memento

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	testMementoInit = "Initiated"

	testMementoPass1Name = "Initial Originator"
	testMementoPass1Want = "Initiated"

	testMementoPass2Name  = "Change state"
	testMementoPass2Input = "Change"
	testMementoPass2Want  = "Change"

	testMementoPass3Name = "Rollback state"
	testMementoPass3Want = "Initiated"
)

func TestMemento(t *testing.T) {
	originator := NewOriginator(testMementoInit)

	t.Run(testMementoPass1Name, func(t *testing.T) {
		got := originator.GetState()
		want := testMementoPass1Want

		assert.EqualValues(t, got, want)
	})

	t.Run(testMementoPass2Name, func(t *testing.T) {
		originator.ChangeState(testMementoPass2Input)

		got := originator.GetState()
		want := testMementoPass2Want

		assert.EqualValues(t, got, want)
	})

	t.Run(testMementoPass3Name, func(t *testing.T) {
		originator.Rollback()

		got := originator.GetState()
		want := testMementoPass3Want

		assert.EqualValues(t, got, want)
	})
}
