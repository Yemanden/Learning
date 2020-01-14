package service

// Servicer is interface, contains methods of service
type Servicer interface {
	SpecificRequest() string
}

type service struct {
	answer string
}

// SpecificRequest is unique method of Servicer. Returns string s.answer
func (s *service) SpecificRequest() string {
	return s.answer
}

// NewService returns a new object, implemented Servicer
func NewService(answer string) Servicer {
	return &service{answer}
}
