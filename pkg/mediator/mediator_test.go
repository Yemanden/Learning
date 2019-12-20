package mediator

import "testing"

func TestPerformer_Perform(t *testing.T) {
	message := "Hello!"
	message2 := "Добрый вечер!"

	var medic Translator
	a := PerformerA{&medic}
	b := PerformerB{&medic}

	medic.Englishman = &a
	medic.Russian = &b

	a.Perform(message)
	b.Perform(message2)
}
