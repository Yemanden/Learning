package observer

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Yemanden/Learning/pkg/observer/publisher"
)

const (
	testObserverPass1Name = "Two observers"
	testObserverPass2Name = "One observer"
	testObserverPass3Name = "Remove Observer"

	testObserverInput = "Hi!"
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
		a := NewConcreteObserverA()
		b := NewConcreteObserverB()
		publ.AddObserver(a)
		publ.AddObserver(b)

		got := publ.Publish(testObserverInput)
		want := testObserverPass1Want

		assert.EqualValues(t, got, want)
	})

	t.Run(testObserverPass2Name, func(t *testing.T) {
		publ := publisher.NewConcretePublisher()
		a := NewConcreteObserverA()
		publ.AddObserver(a)

		got := publ.Publish(testObserverInput)
		want := testObserverPass2Want

		assert.EqualValues(t, got, want)
	})
	t.Run(testObserverPass3Name, func(t *testing.T) {
		publ := publisher.NewConcretePublisher()
		a := NewConcreteObserverA()
		b := NewConcreteObserverB()
		publ.AddObserver(a)
		publ.AddObserver(b)
		publ.RemoveObserver(a)

		got := publ.Publish(testObserverInput)
		want := testObserverPass3Want

		assert.EqualValues(t, got, want)
	})
}
