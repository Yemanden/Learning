package publisher

import "github.com/Yemanden/Learning/pkg/observer"

// Publisher is interface, contains methods Publish, AddObserver and RemoveObserver.
type Publisher interface {
	Publish(string) []string
	AddObserver(observer.Observer)
	RemoveObserver(observer.Observer)
}
