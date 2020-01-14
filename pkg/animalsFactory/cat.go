package animalsFactory

import (
	"fmt"
)

// Structure contained the data of cats
type cat struct {
	Name, Color, Sound string
	Age                int
}

// SetData sets cats name, color, sound and age
func (c *cat) SetData(name, color, sound string, age int) {
	c.Name = name
	c.Color = color
	c.Sound = sound
	c.Age = age
}

// GetData returns all data of cat
func (c *cat) GetData() (string, string, string, int) {
	return c.Name, c.Color, c.Sound, c.Age
}

// Voice displays a cats sound
func (c *cat) Voice() {
	fmt.Println(c.Sound)
}

func newCat() Animal {
	return &cat{}
}
