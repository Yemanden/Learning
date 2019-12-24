package mediator

// Performer is interface, contains methods Send and receive
type Performer interface {
	Send(string) string
	receive(string) string
}
