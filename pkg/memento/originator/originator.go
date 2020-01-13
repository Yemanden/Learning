package originator

import "github.com/Yemanden/Learning/pkg/memento/caretaker"

type memento interface {
	GetState() string
}

// Originator is interface, contains methods ChangeState, Rollback and GetState
type Originator interface {
	ChangeState(memento)
	Rollback() error
	GetState() string
}

type originator struct {
	state  memento
	keeper caretaker.CareTaker
}

// ChangeState sends current state in careTaker and changes state
func (o *originator) ChangeState(state memento) {
	o.keeper.Push(o.state)
	o.state = state
}

// Rollback returns previous state instead current state
func (o *originator) Rollback() error {
	state, err := o.keeper.Pop()
	if err != nil {
		return err
	}
	o.state = state
	return nil
}

// GetState returns current state
func (o *originator) GetState() string {
	return o.state.GetState()
}

// NewOriginator returns new originator with state s
func NewOriginator(state memento) Originator {
	return &originator{
		state:  state,
		keeper: caretaker.NewCareTaker(),
	}
}
