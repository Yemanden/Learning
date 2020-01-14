package client

// Client is interface, contains methods of client
type Client interface {
	Request() string
}

type client struct {
	answer string
}

// Request is unique method of Client. Returns string c.answer
func (c *client) Request() string {
	return c.answer
}

// NewClient returns a new object, implemented Client
func NewClient(answer string) Client {
	return &client{answer}
}
