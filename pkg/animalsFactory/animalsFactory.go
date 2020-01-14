package animalsFactory

// Creator is interface, contains method Create
type Creator interface {
	Create(string) Animal
}

// An empty structure that has a Create method
type creator struct{}

// Create takes the name of the structure to create.
// Returns a new object implements interface Animal
func (c *creator) Create(name string) Animal {
	var anim Animal
	switch name {
	case "Cat":
		{
			anim = newCat()
			break
		}
	case "Dog":
		{
			anim = newDog()
			break
		}
	case "Duck":
		{
			anim = newDuck()
			break
		}
	}

	return anim
}

// NewCreator returns a new object of creator struct
func NewCreator() Creator {
	return &creator{}
}
