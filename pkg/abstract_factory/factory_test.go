package abstractfactory

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	testFactoryPass1Name = "Get Price productA"
	testFactoryPass2Name = "Get Price productB"
)

var (
	testFactoryPass1InputAndWant = 10
	testFactoryPass2InputAndWant = 14
)

func TestFactory(t *testing.T) {
	t.Run(testFactoryPass1Name, func(t *testing.T) {
		p := NewProductA(testFactoryPass1InputAndWant)

		got := p.GetPrice()
		want := testFactoryPass1InputAndWant

		assert.EqualValues(t, got, want)
	})

	t.Run(testFactoryPass2Name, func(t *testing.T) {
		p := NewProductB(testFactoryPass2InputAndWant)

		got := p.GetPrice()
		want := testFactoryPass2InputAndWant

		assert.EqualValues(t, got, want)
	})
}
