package chain_of_resp

// HandlerA is example Handlers
type handlerB struct {
	next Handler
}

// Request is example of a Request method for Handlers
func (h *handlerB) Handle(s string) string {
	s = s + " HandlerB"

	if h.next != nil {
		return h.next.Handle(s)
	}
	return s
}

func NewHandlerB(handler Handler) Handler {
	return &handlerB{handler}
}
