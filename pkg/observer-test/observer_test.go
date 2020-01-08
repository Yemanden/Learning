package observer_test

import (
	"github.com/Yemanden/Learning/pkg/observer"
	"github.com/Yemanden/Learning/pkg/publisher"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	testObserverInput     = "Hi!"
	testObserverPass1Name = "Two observers"
	testObserverPass2Name = "One observer"
	testObserverPass3Name = "Remove Observer"
)

var (
	testObserverPass1Want = []string{"I am ObserverA Hi!", "I am ObserverB Hi!"}
	testObserverPass2Want = []string{"I am ObserverA Hi!"}
	testObserverPass3Want = []string{"I am ObserverB Hi!"}
)

// TestObserver ...
func TestObserver(t *testing.T) {
	t.Run(testObserverPass1Name, func(t *testing.T) {
		publ := publisher.NewConcretePublisher()
		a := observer.NewConcreteObserverA()
		b := observer.NewConcreteObserverB()
		publ.AddObserver(a)
		publ.AddObserver(b)

		got := publ.Publish(testObserverInput)
		want := testObserverPass1Want

		assert.EqualValues(t, got, want)
	})

	t.Run(testObserverPass2Name, func(t *testing.T) {
		publisher := publisher.NewConcretePublisher()
		a := observer.NewConcreteObserverA()
		publisher.AddObserver(a)

		got := publisher.Publish(testObserverInput)
		want := testObserverPass2Want

		assert.EqualValues(t, got, want)
	})
	t.Run(testObserverPass3Name, func(t *testing.T) {
		publisher := publisher.NewConcretePublisher()
		a := observer.NewConcreteObserverA()
		b := observer.NewConcreteObserverB()
		publisher.AddObserver(a)
		publisher.AddObserver(b)
		publisher.RemoveObserver(a)

		got := publisher.Publish(testObserverInput)
		want := testObserverPass3Want

		assert.EqualValues(t, got, want)
	})
}
