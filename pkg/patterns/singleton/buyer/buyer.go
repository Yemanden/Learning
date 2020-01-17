package buyer

type (
	clothes  = string
	wardrobe = []clothes
)

type seller interface {
	Sale(clothes) (clothes, error)
}

type Buyer interface {
	Buy(clothes, seller) error
}

type buyer struct {
	wardrobe wardrobe
}

func (b *buyer) Buy(garmentName clothes, seller seller) (error error) {
	garment, error := seller.Sale(garmentName)
	if error != nil {
		return
	}
	b.wardrobe = append(b.wardrobe, garment)
	return
}

func NewBuyer() Buyer {
	return &buyer{}
}
