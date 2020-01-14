package service

// Service is interface, contains methods of service
type Service interface {
	SpecificRequest() string
}

type service struct {
	answer string
}

// SpecificRequest is unique method of Service. Returns string s.answer
func (s *service) SpecificRequest() string {
	return s.answer
}

// NewService returns a new object, implemented Service
func NewService(answer string) Service {
	return &service{answer}
}
