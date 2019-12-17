package Animals

import (
	"fmt"
	"testing"
)

func TestNewCreator(t *testing.T) {
	got := NewCreator()
	want := &creator{}

	if got != want {
		fmt.Printf("%s: got [%v] want [%v]", "NewCreator", got, want)
	}
}

func TestCreator_Create(t *testing.T) {

	t.Run("dog", func(t *testing.T) {
		temp := NewCreator().Create("Dog")
		temp.SetData("name", "black", 3)

		got := temp.GetAge()
		want := 3

		if got != want {
			fmt.Printf("%s: got [%v] want [%v]", "Create Dog", got, want)
		}
	})
	t.Run("Cat", func(t *testing.T) {
		temp := NewCreator().Create("Cat")
		temp.SetData("name", "black", 3)

		got := temp.GetAge()
		want := 3

		if got != want {
			fmt.Printf("%s: got [%v] want [%v]", "Create Cat", got, want)
		}
	})
	t.Run("Duck", func(t *testing.T) {
		temp := NewCreator().Create("Duck")
		temp.SetData("name", "black", 3)

		got := temp.GetAge()
		want := 3

		if got != want {
			fmt.Printf("%s: got [%v] want [%v]", "Create Duck", got, want)
		}
	})
}

func TestAnimal_Voice(t *testing.T) {
	NewCreator().Create("animal").Voice()
}

func TestCat_Voice(t *testing.T) {
	NewCreator().Create("Cat").Voice()
}

func TestDog_Voice(t *testing.T) {
	NewCreator().Create("Dog").Voice()
}

func TestDuck_Voice(t *testing.T) {
	NewCreator().Create("Duck").Voice()
}
