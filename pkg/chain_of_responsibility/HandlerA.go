package chain_of_resp

// HandlerA is example Handlers
type handlerA struct {
	next Handler
}

// Request is example of a Request method for Handlers
func (h *handlerA) Handle(s string) string {
	s = s + " HandlerA"

	if h.next != nil {
		return h.next.Handle(s)
	}
	return s
}

func NewHandlerA(handler Handler) Handler {
	return &handlerA{handler}
}
