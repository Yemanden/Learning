package visitor

// Bank ...
type Bank interface {
	SeeBalance() string
	Place
}

type bank struct{}

// SeeBalance ...
func (b *bank) SeeBalance() string {
	return "10 dollars "
}

// Accept ...
func (b *bank) Accept(v Visitor) string {
	return v.VisitBank(b)
}

// NewBank ...
func NewBank() Bank {
	return &bank{}
}
