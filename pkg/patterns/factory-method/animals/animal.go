package animals

// Animal is interface, describing the behavior of all animals
type Animal interface {
	Voice()
	SetData(string, string, string, int)
	GetData() (string, string, string, int)
}
