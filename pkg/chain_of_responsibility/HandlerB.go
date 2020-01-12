package chainofresp

// handlerB is example Handlers
type handlerB struct {
	next Handler
}

// Handle is example of a Request method for Handlers
func (h *handlerB) Handle(s string) string {
	s = s + " HandlerB"

	if h.next != nil {
		return h.next.Handle(s)
	}
	return s
}

// NewHandlerB is constructor, returns a new object of handlerB structure
func NewHandlerB(handler Handler) Handler {
	return &handlerB{handler}
}
