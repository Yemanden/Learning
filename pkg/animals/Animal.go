package Animals

import "fmt"

// Interface describing the behavior of all animals
type iAnimal interface {
	Voice()
	SetData(string, string, int)
	GetAge() int
}

// Structure contained the data of all animals
type animal struct {
	Name  string
	Color string
	Age   int
	iAnimal
}

// Structure contained the data of cats
type cat struct {
	animal
}

// Structure contained the data of ducks
type duck struct {
	animal
}

// Structure contained the data of dogs
type dog struct {
	animal
}

// SetData arguments is name, color and age of animal. Sets data, describes in struct animal
func (obj *animal) SetData(name string, color string, age int) {
	obj.Name = name
	obj.Color = color
	obj.Age = age
}

// GetAge return Age of this animal
func (obj *animal) GetAge() int {
	return obj.Age
}

// Voice is default method for all animals
func (obj *animal) Voice() {
	fmt.Println("animals sound")
}

// Voice displays a "Meow"
func (obj *cat) Voice() {
	fmt.Println("Meow")
}

// Voice displays a "Woof woof"
func (obj *dog) Voice() {
	fmt.Println("Woof woof")
}

// Voice displays a "Quack quack"
func (obj *duck) Voice() {
	fmt.Println("Quack quack")
}
