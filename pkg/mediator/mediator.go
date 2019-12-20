package mediator

import "fmt"

// Sender is interface for Mediator
type Sender interface {
	send(string)
}

type performer interface {
	Perform(string)
	receive(string)
}

// Translator is structure of concrete Mediator
type Translator struct {
	Englishman, Russian performer
}

func (t *Translator) send(s string) {
	if s == "Hello!" && t.Russian != nil {
		t.Russian.receive("С учётом погрешности: Здарова!")
	}
	if s == "Добрый вечер!" && t.Englishman != nil {
		t.Englishman.receive("Given the error: Zdarova!")
	}
}

// PerformerA is concrete Performer
type PerformerA struct {
	Mediator *Translator
}

// Perform sends message to Mediator
func (p *PerformerA) Perform(msg string) {
	p.Mediator.send(msg)
}

func (p *PerformerA) receive(s string) {
	fmt.Println(s)
}

// PerformerB is concrete Performer
type PerformerB struct {
	Mediator *Translator
}

// Perform sends message to Mediator
func (p *PerformerB) Perform(msg string) {
	p.Mediator.send(msg)
}

func (p *PerformerB) receive(s string) {
	fmt.Println(s)
}
