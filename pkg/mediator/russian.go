package mediator

// russian is concrete Performer
type russian struct {
	mediator Mediator
}

// Send sends message to Mediator
func (p *russian) Send(msg string) string {
	return p.mediator.send(msg)
}

// SetMediator ...
func (r *russian) SetMediator(m Mediator) {
	r.mediator = m
}

func (p *russian) receive(s string) string {
	return s + " Успешно!"
}

// NewRussian returns a new russian
func NewRussian() Performer {
	return &russian{}
}
