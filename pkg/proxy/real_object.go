package proxy

// realObject is a real object
type realSubject struct {
	data int
}

// SetData sets data in realSubject s
func (s *realSubject) SetData(i int) {
	s.data = i
}

// GetData returns data from realSubject s
func (s *realSubject) GetData() int {
	return s.data
}

func newRealObject() Proxy {
	return &realSubject{}
}
