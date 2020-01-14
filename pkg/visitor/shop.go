package visitor

// Shop ...
type Shop interface {
	BuyProduct() string
	Place
}

type shop struct {
	product string
}

// BuyProduct ...
func (s *shop) BuyProduct() string {
	return "Buy " + s.product + " "
}

// Accept ...
func (s *shop) Accept(v Visitor) string {
	return v.VisitShop(s)
}

// NewShop ...
func NewShop(product string) Shop {
	return &shop{product}
}
