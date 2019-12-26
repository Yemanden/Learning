package visitor

// Place ...
type Place interface {
	Accept(v Visitor) string
}
