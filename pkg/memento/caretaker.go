package memento

// CareTaker ...
type CareTaker interface {
	Push(*memento)
	Pop() *memento
}

type careTaker struct {
	mementos []*memento
}

// Push ...
func (c *careTaker) Push(m *memento) {
	c.mementos = append(c.mementos, m)
}

// Pop ...
func (c *careTaker) Pop() *memento {
	if len(c.mementos) == 0 {
		return &memento{}
	}

	tmp := c.mementos[len(c.mementos)-1]
	c.mementos = c.mementos[:len(c.mementos)-1]

	return tmp
}
