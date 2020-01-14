package factory

import "github.com/Yemanden/Learning/pkg/patterns/factory-method/animals"

// Creator is interface, contains method Create
type Creator interface {
	Create(string) animals.Animal
}

// An empty structure that has a Create method
type creator struct{}

// Create takes the name of the structure to create.
// Returns a new object implements interface Animal
func (c *creator) Create(name string) animals.Animal {
	var anim animals.Animal
	switch name {
	case "Cat":
		{
			anim = animals.NewCat()
			break
		}
	case "Dog":
		{
			anim = animals.NewDog()
			break
		}
	case "Duck":
		{
			anim = animals.NewDuck()
			break
		}
	}

	return anim
}

// NewCreator returns a new object of creator struct
func NewCreator() Creator {
	return &creator{}
}
