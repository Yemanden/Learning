package proxy

// realObject is a real object
type realObject struct {
	data int
}

// SetData sets data in realObject s
func (s *realObject) SetData(i int) {
	s.data = i
}

// GetData returns data from realObject s
func (s *realObject) GetData() int {
	return s.data
}

func newRealObject() Proxer {
	return &realObject{}
}
