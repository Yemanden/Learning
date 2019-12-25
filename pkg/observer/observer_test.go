package observer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	testObserverPass1Name = "Two observers"
	testObserverPass2Name = "One observer"
	testObserverInput     = "Hi!"
)

func TestObserver(t *testing.T) {
	t.Run(testObserverPass1Name, func(t *testing.T) {
		publisher := NewConcretePublisher()
		a := NewConcreteObserverA()
		b := NewConcreteObserverB()
		publisher.AddObserver(a)
		publisher.AddObserver(b)

		got := publisher.Publish(testObserverInput)
		want := []string{"I am ObserverA Hi!", "I am ObserverB Hi!"}

		assert.EqualValues(t, got, want)
	})

	t.Run(testObserverPass2Name, func(t *testing.T) {
		publisher := NewConcretePublisher()
		a := NewConcreteObserverA()
		publisher.AddObserver(a)

		got := publisher.Publish(testObserverInput)
		want := []string{"I am ObserverA Hi!"}

		assert.EqualValues(t, got, want)
	})
}
