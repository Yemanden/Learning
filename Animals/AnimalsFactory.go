package Animals

// NewCreator returns a new object of creator struct
func NewCreator() *creator {
	return &creator{}
}

// An empty structure that has a Create method
type creator struct {
}

// Create takes the name of the structure to create.
// Returns a new object implements interface iAnimal
func (c *creator) Create(name string) iAnimal {
	var anim iAnimal
	switch name {
	case "Cat":
		{
			anim = &cat{}
			break
		}
	case "Dog":
		{
			anim = &dog{}
			break
		}
	case "Duck":
		{
			anim = &duck{}
			break
		}
	default:
		anim = &animal{}
	}

	return anim
}
