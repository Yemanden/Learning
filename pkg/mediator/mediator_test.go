package mediator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	testMediatorPass1Name = "TestMediator Russian to Englishman"
	testMediatorInput1    = "Добрый вечер!"
	testMediatorWant1     = "Given the error: Zdarova! Done!"

	testMediatorPass2Name = "TestMediator Englishman to Russian"
	testMediatorInput2    = "Hello!"
	testMediatorWant2     = "С учётом погрешности: Здарова! Успешно!"

	testMediatorPass3Name = "TestMediator Missing"
	testMediatorInput3    = "Test"
	testMediatorWant3     = "Ja ja, naturlich!"
)

// TestMediator ...
func TestMediator(t *testing.T) {
	t.Run(testMediatorPass1Name, func(t *testing.T) {
		medic := NewConcreteMediator()
		a := NewEnglishman()
		a.SetMediator(medic)
		b := NewRussian()
		b.SetMediator(medic)

		medic.SetPerformers(a, b)

		got := b.Send(testMediatorInput1)
		want := testMediatorWant1

		assert.EqualValues(t, want, got)
	})

	t.Run(testMediatorPass2Name, func(t *testing.T) {
		medic := NewConcreteMediator()
		a := NewEnglishman()
		a.SetMediator(medic)
		b := NewRussian()
		b.SetMediator(medic)

		medic.SetPerformers(a, b)

		got := a.Send(testMediatorInput2)
		want := testMediatorWant2

		assert.EqualValues(t, want, got)
	})

	t.Run(testMediatorPass3Name, func(t *testing.T) {
		medic := NewConcreteMediator()
		a := NewEnglishman()
		a.SetMediator(medic)
		b := NewRussian()
		b.SetMediator(medic)

		medic.SetPerformers(a, b)

		got := b.Send(testMediatorInput3)
		want := testMediatorWant3
		assert.EqualValues(t, want, got)
	})
}
