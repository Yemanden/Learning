package observer

// Observer is interface, contains method Observe
type Observer interface {
	Observe(string) string
}
