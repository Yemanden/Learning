package animalsFactory

import (
	"fmt"
)

// Structure contained the data of ducks
type duck struct {
	Name, Color, Sound string
	Age                int
}

// SetData sets ducks name, color, sound and age
func (d *duck) SetData(name, color, sound string, age int) {
	d.Name = name
	d.Color = color
	d.Sound = sound
	d.Age = age
}

// GetData returns all data of duck
func (d *duck) GetData() (string, string, string, int) {
	return d.Name, d.Color, d.Sound, d.Age
}

// Voice displays a ducks sound
func (d *duck) Voice() {
	fmt.Println(d.Sound)
}

func newDuck() Animal {
	return &duck{}
}