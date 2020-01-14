package visitor

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	testVisitorPass1Name  = "Visit Shop"
	testVisitorPass1Input = "chicken"
	testVisitorPass1Want  = "Buy chicken "

	testVisitorPass2Name        = "Visit Shop, Bank and Cinema"
	testVisitorPass2InputShop   = "chicken"
	testVisitorPass2InputBank   = 10
	testVisitorPass2InputCinema = "Star Wars"
	testVisitorPass2Want        = "Buy chicken 10 dollars See \"Star Wars\" "
)

// TestVisitor ...
func TestVisitor(t *testing.T) {
	t.Run(testVisitorPass1Name, func(t *testing.T) {
		want := testVisitorPass1Want

		v := NewVisitor()
		s := NewShop(testVisitorPass1Input)
		got := s.Accept(v)

		assert.EqualValues(t, want, got)
	})

	t.Run(testVisitorPass2Name, func(t *testing.T) {
		want := testVisitorPass2Want

		visitor := NewVisitor()
		places := []Place{NewShop(testVisitorPass2InputShop), NewBank(testVisitorPass2InputBank), NewCinema(testVisitorPass2InputCinema)}
		var got string
		for _, item := range places {
			got += item.Accept(visitor)
		}

		assert.EqualValues(t, want, got)
	})
}
