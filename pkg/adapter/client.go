package adapter

// Client is interface, contains methods of client
type Client interface {
	Request() string
}

type client struct{}

// Request is unique method of Client. Returns string "Client!"
func (c *client) Request() string {
	return "Client!"
}

// NewClient returns a new object, implemented Client
func NewClient() Client {
	return &client{}
}
