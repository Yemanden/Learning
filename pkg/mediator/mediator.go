package mediator

// Mediator ...
type Mediator interface {
	send(string) string
	SetPerformers(Performer, Performer)
}

type concreteMediator struct {
	englishman, russian Performer
}

// SetPerformers sets new Performers in Mediator
func (c *concreteMediator) SetPerformers(p1, p2 Performer) {
	c.englishman = p1
	c.russian = p2
}

func (c *concreteMediator) send(s string) string {
	if s == "Hello!" && c.russian != nil {
		return c.russian.receive("С учётом погрешности: Здарова!")
	}
	if s == "Добрый вечер!" && c.englishman != nil {
		return c.englishman.receive("Given the error: Zdarova!")
	}
	return "Ja ja, naturlich!"
}

// NewConcreteMediator creates and returns a new object of concreteMediator
func NewConcreteMediator() Mediator {
	return &concreteMediator{}
}
