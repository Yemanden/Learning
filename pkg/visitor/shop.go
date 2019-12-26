package visitor

// Shop ...
type Shop interface {
	BuyProduct() string
	Place
}

type shop struct{}

// BuyProduct ...
func (s *shop) BuyProduct() string {
	return "Buy one chicken "
}

// Accept ...
func (s *shop) Accept(v Visitor) string {
	return v.VisitShop(s)
}

// NewShop ...
func NewShop() Shop {
	return &shop{}
}
