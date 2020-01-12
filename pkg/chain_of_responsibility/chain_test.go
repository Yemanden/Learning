package chainofresp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	testChainPass1Name  = "Three handlers"
	testChainPass1Input = "input"
	testChainPass1Want  = "input HandlerA HandlerA HandlerB"

	testChainPass2Name  = "Two handlers"
	testChainPass2Input = "input"
	testChainPass2Want  = "input HandlerA HandlerB"
)

func TestChain(t *testing.T) {
	t.Run(testChainPass1Name, func(t *testing.T) {
		chain := NewHandlerA(NewHandlerA(NewHandlerB(nil)))

		got := chain.Handle(testChainPass1Input)
		want := testChainPass1Want

		assert.EqualValues(t, got, want)
	})

	t.Run(testChainPass2Name, func(t *testing.T) {
		chain := NewHandlerA(NewHandlerB(nil))

		got := chain.Handle(testChainPass2Input)
		want := testChainPass2Want

		assert.EqualValues(t, got, want)
	})
}
