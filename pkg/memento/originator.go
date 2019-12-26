package memento

// Originator is interface, contains methods ChangeState, Rollback and GetState
type Originator interface {
	ChangeState(string)
	Rollback()
	GetState() string
}

type originator struct {
	state  *memento
	keeper *careTaker
}

// ChangeState sends current state in careTaker and changes state
func (o *originator) ChangeState(s string) {
	o.keeper.Push(o.state)
	memento := &memento{s}
	o.state = memento
}

// Rollback returns previous state instead current state
func (o *originator) Rollback() {
	o.state = o.keeper.Pop()
}

// GetState returns current state
func (o *originator) GetState() string {
	return o.state.data
}

// NewOriginator returns new originator with state s
func NewOriginator(s string) Originator {
	return &originator{
		state:  &memento{s},
		keeper: &careTaker{},
	}
}
