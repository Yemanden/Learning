package visitor

import "fmt"

// Banker ...
type Banker interface {
	SeeBalance() string
	Place
}

type bank struct {
	balance int
}

// SeeBalance ...
func (b *bank) SeeBalance() string {
	return fmt.Sprintf("%d %s ", b.balance, "dollars")
}

// Accept ...
func (b *bank) Accept(v Visitor) string {
	return v.VisitBank(b)
}

// NewBank ...
func NewBank(balance int) Banker {
	return &bank{balance}
}
