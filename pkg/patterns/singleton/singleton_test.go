package singleton

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Yemanden/Learning/pkg/models"
	"github.com/Yemanden/Learning/pkg/patterns/singleton/buyer"
	"github.com/Yemanden/Learning/pkg/patterns/singleton/shop"
)

const (
	testSellerName = "Test GetSeller"

	testBuyerPass1Name = "Clothes available"
	testBuyerPass2Name = "No clothes available"
)

// TestGetInstance ...
func TestSeller(t *testing.T) {
	t.Run(testSellerName, func(t *testing.T) {
		want := shop.GetSeller()

		got := shop.GetSeller()
		got.Sale(models.Pants)

		assert.EqualValues(t, want, got)
	})
}

func TestBuyer(t *testing.T) {
	t.Run(testBuyerPass1Name, func(t *testing.T) {
		error := buyer.NewBuyer().Buy(models.Pants, shop.GetSeller())
		assert.NoError(t, error)
	})

	var err error
	for err == nil {
		err = buyer.NewBuyer().Buy(models.Pants, shop.GetSeller())
	}

	t.Run(testBuyerPass2Name, func(t *testing.T) {
		error := buyer.NewBuyer().Buy(models.Pants, shop.GetSeller())
		assert.Error(t, error)
	})
}
