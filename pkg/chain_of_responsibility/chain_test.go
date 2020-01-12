package chainofresp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	TestChainPass1Name  = "Three handlers"
	TestChainPass1Input = "input"
	TestChainPass1Want  = "input HandlerA HandlerA HandlerB"

	TestChainPass2Name  = "Two handlers"
	TestChainPass2Input = "input"
	TestChainPass2Want  = "input HandlerA HandlerB"
)

func TestChain(t *testing.T) {
	t.Run(TestChainPass1Name, func(t *testing.T) {
		chain := NewHandlerA(NewHandlerA(NewHandlerB(nil)))

		got := chain.Handle(TestChainPass1Input)
		want := TestChainPass1Want

		assert.EqualValues(t, got, want)
	})

	t.Run(TestChainPass2Name, func(t *testing.T) {
		chain := NewHandlerA(NewHandlerB(nil))

		got := chain.Handle(TestChainPass2Input)
		want := TestChainPass2Want

		assert.EqualValues(t, got, want)
	})
}
