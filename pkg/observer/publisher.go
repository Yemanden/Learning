package observer

// Publisher is interface, contains methods Publish, AddObserver and RemoveObserver.
type Publisher interface {
	Publish(string) []string
	AddObserver(Observer)
	RemoveObserver(Observer)
}
