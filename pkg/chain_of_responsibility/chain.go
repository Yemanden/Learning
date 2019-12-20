package chain_of_resp

import "fmt"

// Handler interface have method Request
type Handler interface {
	Request(string)
}

// ChainLink is LinkList of the Handler
type ChainLink struct {
	next    *ChainLink
	current *Handler
}

// Request begins the execution of the entire chain
func (c ChainLink) Request(s string) {
	cur := c.current
	if cur == nil {
		return
	}
	(*cur).Request(s)

	if c.next == nil {
		return
	}
	c.next.Request(s)
}

// Append appends a new Handler in Chain
func (c *ChainLink) Append(h Handler) {
	if c.current == nil {
		c.current = &h
		return
	}

	tmp := &c.next
	for {
		if *tmp == nil {
			*tmp = &ChainLink{nil, &h}
			return
		}
		tmp = &(*tmp).next
	}
}

// HandlerA is example Handlers
type HandlerA struct {
}

// Request is example of a Request method for Handlers
func (h HandlerA) Request(s string) {
	fmt.Println("HandlerA: " + s)
}
