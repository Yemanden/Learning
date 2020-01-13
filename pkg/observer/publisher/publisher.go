package publisher

type observer interface {
	Observe(string) string
}

// Publisher is interface, contains methods Publish, AddObserver and RemoveObserver.
type Publisher interface {
	Publish(string) []string
	AddObserver(observer)
	RemoveObserver(observer)
}
