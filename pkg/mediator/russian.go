package mediator

// russian is concrete Performer
type russian struct {
	mediator Mediator
}

// Send sends message to Mediator
func (r *russian) Send(msg string) string {
	return r.mediator.send(msg)
}

// SetMediator ...
func (r *russian) SetMediator(m Mediator) {
	r.mediator = m
}

func (r *russian) receive(s string) string {
	return s + " Успешно!"
}

// NewRussian returns a new russian
func NewRussian() Performer {
	return &russian{}
}
