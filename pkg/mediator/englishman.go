package mediator

// PerformerA is concrete Performer
type englishman struct {
	Mediator Mediator
}

// Send sends message to Mediator
func (p *englishman) Send(msg string) string {
	return p.Mediator.send(msg)
}

func (p *englishman) receive(s string) string {
	return s + " Done!"
}

// NewEnglishman returns a new englishman
func NewEnglishman(m *Mediator) Performer {
	return &englishman{*m}
}
