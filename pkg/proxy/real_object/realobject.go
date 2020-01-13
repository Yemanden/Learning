package realobject

type proxer interface {
	setter
	getter
}

// Setter ...
type setter interface {
	SetData(int)
}

// Getter ...
type getter interface {
	GetData() int
}

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

func NewRealObject() proxer {
	return &realObject{}
}
