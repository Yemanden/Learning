package memento

import (
	"github.com/Yemanden/Learning/pkg/memento/memento"
	originator2 "github.com/Yemanden/Learning/pkg/memento/originator"
	"testing"

	"github.com/stretchr/testify/assert"
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

// TestMemento ...
func TestMemento(t *testing.T) {
	mem := memento.NewMemento(testMementoInit)
	orig := originator2.NewOriginator(mem)

	t.Run(testMementoPass1Name, func(t *testing.T) {
		got := orig.GetState()
		want := testMementoPass1Want

		assert.EqualValues(t, want, got)
	})

	t.Run(testMementoPass2Name, func(t *testing.T) {
		mem2 := memento.NewMemento(testMementoPass2Input)
		orig.ChangeState(mem2)

		got := orig.GetState()
		want := testMementoPass2Want

		assert.EqualValues(t, want, got)
	})

	t.Run(testMementoPass3Name, func(t *testing.T) {
		orig.Rollback()

		got := orig.GetState()
		want := testMementoPass3Want

		assert.EqualValues(t, want, got)
	})
}
