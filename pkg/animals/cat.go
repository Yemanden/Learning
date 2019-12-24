package animals

import "fmt"

// Structure contained the data of cats
type cat struct {
	Name, Color, Sound string
	Age                int
}

func (c *cat) SetData(name, color, sound string, age int) {
	c.Name = name
	c.Color = color
	c.Sound = sound
	c.Age = age
}

func (c *cat) GetData() (string, string, string, int) {
	return c.Name, c.Color, c.Sound, c.Age
}

// Voice displays a "Meow"
func (c *cat) Voice() {
	fmt.Println(c.Sound)
}

func newCat() Animal {
	return &cat{}
}
