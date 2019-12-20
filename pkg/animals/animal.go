package animals

import "fmt"

// Interface describing the behavior of all animals
type Animaler interface {
	Voice()
	SetData(string, string, int)
	GetAge() int
}

// Structure contained the data of all animals
type animal struct {
	Name, Color string
	Age         int
	Animaler
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
func (a *animal) SetData(name string, color string, age int) {
	a.Name = name
	a.Color = color
	a.Age = age
}

// GetAge return Age of this animal
func (a *animal) GetAge() int {
	return a.Age
}

// Voice is default method for all animals
func (a *animal) Voice() {
	fmt.Println("animals sound")
}

// Voice displays a "Meow"
func (c *cat) Voice() {
	fmt.Println("Meow")
}

// Voice displays a "Woof woof"
func (d *dog) Voice() {
	fmt.Println("Woof woof")
}

// Voice displays a "Quack quack"
func (d *duck) Voice() {
	fmt.Println("Quack quack")
}
