package chain_of_resp

import (
	"fmt"
	"testing"
)

func TestChainLink_Append(t *testing.T) {
	var handler HandlerA
	var chain ChainLink
	chain.Append(handler)

	got := *chain.current
	want := handler

	if got != want {
		fmt.Printf("%s: got [%v] want [%v]", "Append: ", got, want)
	}
}

func TestChainLink_Request(t *testing.T) {
	var a HandlerA
	var chain ChainLink

	t.Run("Without handler", func(t *testing.T) {
		chain.Request("Without handler")
	})

	chain = ChainLink{}
	t.Run("One handler", func(t *testing.T) {
		chain.Append(a)
		chain.Request("One handler")
	})

	chain = ChainLink{}
	t.Run("Two handler", func(t *testing.T) {
		chain.Append(a)
		chain.Append(a)
		chain.Request("Two handler")
	})

	chain = ChainLink{}
	t.Run("Three handler", func(t *testing.T) {
		chain.Append(a)
		chain.Append(a)
		chain.Append(a)
		chain.Request("Three handler")
	})
	chain = ChainLink{}
	t.Run("One handler + one chain", func(t *testing.T) {
		chain.Append(a)
		var chain2 ChainLink
		chain2.Append(a)
		chain.Append(chain2)
		chain.Request("One handler + one chain")
	})

}
