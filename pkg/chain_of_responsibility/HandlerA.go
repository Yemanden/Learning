package chainofresp

// handlerA is example Handlers
type handlerA struct {
	next Handler
}

// Handle is example of a Request method for Handlers
func (h *handlerA) Handle(s string) string {
	s = s + " HandlerA"

	if h.next != nil {
		return h.next.Handle(s)
	}
	return s
}

// NewHandlerA is constructor, returns a new object of handlerA structure
func NewHandlerA(handler Handler) Handler {
	return &handlerA{handler}
}
