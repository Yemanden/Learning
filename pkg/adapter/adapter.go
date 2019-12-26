package adapter

// Adapter is interface contains interfaces Client and Service
type Adapter interface {
	Client
	Service
}

type adapter struct {
	client  Client
	service Service
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
func NewAdapter(c Client, s Service) Adapter {
	return &adapter{
		client:  c,
		service: s,
	}
}
