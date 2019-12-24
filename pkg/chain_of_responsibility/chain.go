package chain_of_resp

// Handler is interface, contains method Handle
type Handler interface {
	Handle(string) string
}
