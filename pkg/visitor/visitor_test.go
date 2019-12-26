package visitor

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	testVisitorPass1Name = "Visit Shop"
	testVisitorPass1Want = "Buy one chicken "

	testVisitorPass2Name = "Visit Shop, Bank and Cinema"
	testVisitorPass2Want = "Buy one chicken 10 dollars See \"Star Wars\" "
)

func TestVisitor(t *testing.T) {
	t.Run(testVisitorPass1Name, func(t *testing.T) {
		visitor := NewVisitor()
		shop := NewShop()

		got := shop.Accept(visitor)
		want := testVisitorPass1Want

		assert.EqualValues(t, got, want)
	})

	t.Run(testVisitorPass2Name, func(t *testing.T) {
		visitor := NewVisitor()
		places := []Place{NewShop(), NewBank(), NewCinema()}

		var got string
		for _, item := range places {
			got += item.Accept(visitor)
		}
		want := testVisitorPass2Want

		assert.EqualValues(t, got, want)
	})
}
