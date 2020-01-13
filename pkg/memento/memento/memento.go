package memento

type Memento interface {
	GetState() string
}

type memento struct {
	data string
}

func (m *memento) GetState() string {
	return m.data
}

func NewMemento(state string) Memento {
	return &memento{state}
}
