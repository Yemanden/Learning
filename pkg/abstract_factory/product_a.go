package abstractfactory

type productA struct {
	Price int
}

// GetPrice ...
func (c *productA) GetPrice() int {
	return c.Price
}

// NewProductA ...
func NewProductA(p int) Producter {
	return &productA{p}
}
