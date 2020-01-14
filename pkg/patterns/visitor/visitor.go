package visitor

// Visitor ...
type Visitor interface {
	VisitShop(Shoper) string
	VisitBank(Banker) string
	VisitCinema(Cinemer) string
}

type visitor struct{}

// VisitShop ...
func (v *visitor) VisitShop(s Shoper) string {
	return s.BuyProduct()
}

// VisitBank ...
func (v *visitor) VisitBank(b Banker) string {
	return b.SeeBalance()
}

// VisitCinema ...
func (v *visitor) VisitCinema(c Cinemer) string {
	return c.SeeMovie()
}

// NewVisitor ...
func NewVisitor() Visitor {
	return &visitor{}
}
