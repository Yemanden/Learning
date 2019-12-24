package mediator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	TestMediatorPass1Name = "TestMediator Russian to Englishman"
	TestMediatorPass2Name = "TestMediator Englishman to Russian"
	TestMediatorPass3Name = "TestMediator Missing"
	TestMediatorInput1    = "Добрый вечер!"
	TestMediatorWant1     = "Given the error: Zdarova! Done!"
	TestMediatorInput2    = "Hello!"
	TestMediatorWant2     = "С учётом погрешности: Здарова! Успешно!"
	testMediatorInput3    = "Test"
	TestMediatorWant3     = "Ja ja, naturlich!"
)

func TestMediator(t *testing.T) {
	t.Run(TestMediatorPass1Name, func(t *testing.T) {
		medic := NewConcreteMediator()
		a := NewEnglishman(&medic)
		b := NewRussian(&medic)

		medic.SetPerformers(&a, &b)

		got := b.Send(TestMediatorInput1)
		want := TestMediatorWant1

		assert.EqualValues(t, got, want)
	})
	t.Run(TestMediatorPass2Name, func(t *testing.T) {
		medic := NewConcreteMediator()
		a := NewEnglishman(&medic)
		b := NewRussian(&medic)

		medic.SetPerformers(&a, &b)

		got := a.Send(TestMediatorInput2)
		want := TestMediatorWant2

		assert.EqualValues(t, got, want)
	})
	t.Run(TestMediatorPass3Name, func(t *testing.T) {
		medic := NewConcreteMediator()
		a := NewEnglishman(&medic)
		b := NewRussian(&medic)

		medic.SetPerformers(&a, &b)

		got := b.Send(testMediatorInput3)
		want := TestMediatorWant3
		assert.EqualValues(t, got, want)
	})
}
