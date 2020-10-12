package shop

import (
	"sync"

	"github.com/Yemanden/Learning/pkg/models"
)

type clothes = string

// Global variable for singleton
var singleSeller *seller

// Seller ...
type Seller interface {
	Sale(clothes) (clothes, error)
}

// singleton struct
type seller struct {
	sync.Mutex
	wareHouse wareHouse
}

func (s *seller) Sale(garmentName clothes) (garment clothes, error error) {
	s.Lock()
	defer s.Unlock()
	switch garmentName {
	case models.Shirt:
		{
			if s.wareHouse.shirtCount > 0 {
				s.wareHouse.shirtCount--
				garment = garmentName
				return
			}
		}
	case models.TShirt:
		{
			if s.wareHouse.tShirtCount > 0 {
				s.wareHouse.tShirtCount--
				garment = garmentName
				return
			}
		}
	case models.Pants:
		{
			if s.wareHouse.pantsCount > 0 {
				s.wareHouse.pantsCount--
				garment = garmentName
				return
			}
		}
	}
	error = models.SellerMissingProductError
	return
}

// GetSeller returns ours seller or creates his and returns him
func GetSeller() Seller {
	if singleSeller == nil {
		singleSeller = &seller{
			wareHouse: wareHouse{
				shirtCount:  10,
				tShirtCount: 10,
				pantsCount:  10,
			},
		}
	}
	return singleSeller
}
