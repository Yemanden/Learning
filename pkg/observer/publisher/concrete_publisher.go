package publisher

type concretePublisher struct {
	observers []observer
}

// Publish sends string s for all observers, receives their answers and returns them
func (c *concretePublisher) Publish(s string) []string {
	answer := []string{}
	for _, item := range c.observers {
		answer = append(answer, item.Observe(s))
	}
	return answer
}

// AddObserver appends new observer in observers list
func (c *concretePublisher) AddObserver(observer observer) {
	c.observers = append(c.observers, observer)
}

// RemoveObserver deletes observer from observers list
func (c *concretePublisher) RemoveObserver(observer observer) {
	for i, obs := range c.observers {
		if obs == observer {
			c.observers = append(c.observers[:i], c.observers[i+1:]...)
		}
	}
}

// NewConcretePublisher is constructor a new object of concretePublisher structure
func NewConcretePublisher() Publisher {
	return &concretePublisher{}
}
