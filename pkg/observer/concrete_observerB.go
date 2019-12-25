package observer

type concreteObserverB struct {
}

// Observe returns changed massage s, received from Publisher
func (c *concreteObserverB) Observe(s string) string {
	return "I am ObserverB " + s
}

// NewConcreteObserverA is constructor, returns a new object of concreteObserverB structure
func NewConcreteObserverB() Observer {
	return &concreteObserverB{}
}