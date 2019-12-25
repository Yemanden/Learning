package observer

type concreteObserverA struct{}

// Observe returns changed massage s, received from Publisher
func (c *concreteObserverA) Observe(s string) string {
	return "I am ObserverA " + s
}

// NewConcreteObserverA is constructor, returns a new object of concreteObserverA structure
func NewConcreteObserverA() Observer {
	return &concreteObserverA{}
}
