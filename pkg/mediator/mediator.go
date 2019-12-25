package mediator

// Sender is interface for Mediator
type Mediator interface {
	send(string) string
	SetPerformers(*Performer, *Performer)
}

type concreteMediator struct {
	Englishman, Russian Performer
}

func (p *concreteMediator) send(s string) string {
	if s == "Hello!" && p.Russian != nil {
		return p.Russian.receive("С учётом погрешности: Здарова!")
	}
	if s == "Добрый вечер!" && p.Englishman != nil {
		return p.Englishman.receive("Given the error: Zdarova!")
	}
	return "Ja ja, naturlich!"
}

// SetPerformers sets new Performers in Mediator
func (p *concreteMediator) SetPerformers(p1, p2 *Performer) {
	p.Englishman = *p1
	p.Russian = *p2
}

// NewConcreteMediator creates and returns a new object of concreteMediator
func NewConcreteMediator() Mediator {
	return &concreteMediator{}
}
