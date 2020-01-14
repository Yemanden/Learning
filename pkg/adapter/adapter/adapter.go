package adapter

type client interface {
	Request() string
}

type service interface {
	SpecificRequest() string
}

// Adapter is interface contains interfaces Client and Service
type Adapter interface {
	client
	service
}

type adapter struct {
	client  client
	service service
}

// Request is adaptive method, returns answer from Service
func (a *adapter) Request() string {
	return a.service.SpecificRequest()
}

// SpecificRequest is adaptive method, returns answer from Client
func (a *adapter) SpecificRequest() string {
	return a.client.Request()
}

// NewAdapter returns a new object, implemented Adapter
func NewAdapter(c client, s service) Adapter {
	return &adapter{
		client:  c,
		service: s,
	}
}
