package mediator

// englishman is concrete Performer
type englishman struct {
	mediator Mediator
}

// Send sends message to Mediator
func (e *englishman) Send(msg string) string {
	return e.mediator.send(msg)
}

// SetMediator ...
func (e *englishman) SetMediator(m Mediator) {
	e.mediator = m
}

func (e *englishman) receive(s string) string {
	return s + " Done!"
}

// NewEnglishman returns a new englishman
func NewEnglishman() Performer {
	return &englishman{}
}
