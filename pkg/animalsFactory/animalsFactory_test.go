package animalsFactory

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	TestNewCreatorName    = "NewCreator"
	TestCreatePass1Name   = "Create Cat"
	TestCreatePass2Name   = "Create Dog"
	TestCreatePass3Name   = "Create Duck"
	TestGetDataPass1Name  = "GetData Cat"
	TestGetDataPass1Input = "Murka"
	TestGetDataPass2Name  = "GetData Dog"
	TestGetDataPass2Input = "Sharik"
	TestGetDataPass3Name  = "GetData Duck"
	TestGetDataPass3Input = "GaGa"
)

func TestNewCreator(t *testing.T) {
	t.Run(TestNewCreatorName, func(t *testing.T) {
		got := NewCreator()
		want := &creator{}

		assert.EqualValues(t, got, want)
	})
}

func TestCreate(t *testing.T) {

	t.Run(TestCreatePass1Name, func(t *testing.T) {
		got := NewCreator().Create("Cat")
		want := &cat{}

		assert.EqualValues(t, got, want)
	})
	t.Run(TestCreatePass2Name, func(t *testing.T) {
		got := NewCreator().Create("Dog")
		want := &dog{}

		assert.EqualValues(t, got, want)
	})
	t.Run(TestCreatePass3Name, func(t *testing.T) {
		got := NewCreator().Create("Duck")
		want := &duck{}

		assert.EqualValues(t, got, want)
	})
}

func TestGetData(t *testing.T) {

	t.Run(TestGetDataPass1Name, func(t *testing.T) {
		temp := NewCreator().Create("Cat")
		temp.SetData(TestGetDataPass1Input, "", "", 0)
		got, _, _, _ := temp.GetData()
		want := TestGetDataPass1Input

		assert.EqualValues(t, got, want)
	})
	t.Run(TestGetDataPass2Name, func(t *testing.T) {
		temp := NewCreator().Create("Dog")
		temp.SetData(TestGetDataPass2Input, "", "", 0)
		got, _, _, _ := temp.GetData()
		want := TestGetDataPass2Input

		assert.EqualValues(t, got, want)
	})
	t.Run(TestGetDataPass3Name, func(t *testing.T) {
		temp := NewCreator().Create("Duck")
		temp.SetData(TestGetDataPass3Input, "", "", 0)
		got, _, _, _ := temp.GetData()
		want := TestGetDataPass3Input

		assert.EqualValues(t, got, want)
	})
}
