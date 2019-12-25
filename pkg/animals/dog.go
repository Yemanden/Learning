package animals

import "fmt"

// Structure contained the data of dogs
type dog struct {
	Name, Color, Sound string
	Age                int
}

// SetData sets dogs name, color. sound and age
func (d *dog) SetData(name, color, sound string, age int) {
	d.Name = name
	d.Color = color
	d.Sound = sound
	d.Age = age
}

// GetData returns all data of dog
func (d *dog) GetData() (string, string, string, int) {
	return d.Name, d.Color, d.Sound, d.Age
}

// Voice displays a dogs sound
func (d *dog) Voice() {
	fmt.Println(d.Sound)
}

func newDog() Animal {
	return &dog{}
}
