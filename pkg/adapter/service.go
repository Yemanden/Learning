package adapter

// Service is interface, contains methods of service
type Service interface {
	SpecificRequest() string
}

type service struct{}

// SpecificRequest is unique method of Service. Returns string "Service!"
func (s *service) SpecificRequest() string {
	return "Service!"
}

// NewService returns a new object, implemented Service
func NewService() Service {
	return &service{}
}
