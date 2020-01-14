package visitor

// Place ...
type Place interface {
	Accept(Visitor) string
}
