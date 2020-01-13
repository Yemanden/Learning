package caretaker

import "errors"

type memento interface {
	GetState() string
}

// CareTaker ...
type CareTaker interface {
	Push(memento)
	Pop() (memento, error)
}

type careTaker struct {
	mementos []memento
}

// Push ...
func (c *careTaker) Push(m memento) {
	c.mementos = append(c.mementos, m)
}

// Pop ...
func (c *careTaker) Pop() (memento, error) {
	if len(c.mementos) == 0 {
		return nil, errors.New("no old state")
	}

	tmp := c.mementos[len(c.mementos)-1]
	c.mementos = c.mementos[:len(c.mementos)-1]

	return tmp, nil
}

// NewCareTaker ...
func NewCareTaker() CareTaker {
	return &careTaker{}
}
