package animals

import "fmt"

// Structure contained the data of ducks
type duck struct {
	Name, Color, Sound string
	Age                int
}

func (d *duck) SetData(name, color, sound string, age int) {
	d.Name = name
	d.Color = color
	d.Sound = sound
	d.Age = age
}

func (d *duck) GetData() (string, string, string, int) {
	return d.Name, d.Color, d.Sound, d.Age
}

// Voice displays a "Meow"
func (d *duck) Voice() {
	fmt.Println(d.Sound)
}

func newDuck() Animal {
	return &duck{}
}
