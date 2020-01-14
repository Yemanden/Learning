package visitor

// Cinemer ...
type Cinemer interface {
	SeeMovie() string
	Place
}

type cinema struct {
	film string
}

// SeeMovie returns string "See " + name of film
func (c *cinema) SeeMovie() string {
	return "See \"" + c.film + "\" "
}

// Accept ...
func (c *cinema) Accept(v Visitor) string {
	return v.VisitCinema(c)
}

// NewCinema ...
func NewCinema(film string) Cinemer {
	return &cinema{film}
}
