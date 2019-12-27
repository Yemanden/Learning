package abstractfactory

type productB struct {
	Price int
}

// GetPrice ...
func (c *productB) GetPrice() int {
	return c.Price
}

// NewProductB ...
func NewProductB(p int) Producter {
	return &productB{p}
}
