package client

// Clienter is interface, contains methods of client
type Clienter interface {
	Request() string
}

type client struct {
	answer string
}

// Request is unique method of Clienter. Returns string c.answer
func (c *client) Request() string {
	return c.answer
}

// NewClient returns a new object, implemented Clienter
func NewClient(answer string) Clienter {
	return &client{answer}
}
