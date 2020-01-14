package visitor

// Visitor ...
type Visitor interface {
	VisitShop(Shop) string
	VisitBank(Bank) string
	VisitCinema(Cinema) string
}

type visitor struct{}

// VisitShop ...
func (v *visitor) VisitShop(s Shop) string {
	return s.BuyProduct()
}

// VisitBank ...
func (v *visitor) VisitBank(b Bank) string {
	return b.SeeBalance()
}

// VisitCinema ...
func (v *visitor) VisitCinema(c Cinema) string {
	return c.SeeMovie()
}

// NewVisitor ...
func NewVisitor() Visitor {
	return &visitor{}
}
