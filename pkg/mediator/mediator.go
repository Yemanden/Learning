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
func (p *concreteMediator) SetPerformers(p1, p2 Performer) {
	p.englishman = p1
	p.russian = p2
}

func (p *concreteMediator) send(s string) string {
	if s == "Hello!" && p.russian != nil {
		return p.russian.receive("С учётом погрешности: Здарова!")
	}
	if s == "Добрый вечер!" && p.englishman != nil {
		return p.englishman.receive("Given the error: Zdarova!")
	}
	return "Ja ja, naturlich!"
}

// NewConcreteMediator creates and returns a new object of concreteMediator
func NewConcreteMediator() Mediator {
	return &concreteMediator{}
}
