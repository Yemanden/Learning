package visitor

// Cinema ...
type Cinema interface {
	SeeMovie() string
	Place
}

type cinema struct{}

// SeeMovie returns string "See Star Wars..."
func (c *cinema) SeeMovie() string {
	return "See \"Star Wars\" "
}

// Accept ...
func (c *cinema) Accept(v Visitor) string {
	return v.VisitCinema(c)
}

// NewCinema ...
func NewCinema() Cinema {
	return &cinema{}
}
