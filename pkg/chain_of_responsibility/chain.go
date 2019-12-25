package chainofresp

// Handler is interface, contains method Handle
type Handler interface {
	Handle(string) string
}
