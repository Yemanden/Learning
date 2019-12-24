package mediator

// PerformerB is concrete Performer
type russian struct {
	Mediator Mediator
}

// Send sends message to Mediator
func (p *russian) Send(msg string) string {
	return p.Mediator.send(msg)
}

func (p *russian) receive(s string) string {
	return s + " Успешно!"
}

// NewRussian returns a new russian
func NewRussian(m *Mediator) Performer {
	return &russian{*m}
}
