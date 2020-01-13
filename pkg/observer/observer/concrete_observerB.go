package observer

type concreteObserverB struct{}

// Observe returns changed message s, received from Publisher
func (c *concreteObserverB) Observe(s string) string {
	return "I am ObserverB " + s
}

// NewConcreteObserverB is constructor, returns a new object of concreteObserverB structure
func NewConcreteObserverB() Observer {
	return &concreteObserverB{}
}
