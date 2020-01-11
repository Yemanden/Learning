package Factory_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Yemanden/Learning/pkg/animalsFactory/Factory"
)

const (
	/*	TestNewCreatorName    = "NewCreator"
		TestCreatePass1Name   = "Create Cat"
		TestCreatePass2Name   = "Create Dog"
		TestCreatePass3Name   = "Create Duck"*/
	TestGetDataPass1Name  = "GetData Cat"
	TestGetDataPass1Input = "Murka"
	TestGetDataPass2Name  = "GetData Dog"
	TestGetDataPass2Input = "Sharik"
	TestGetDataPass3Name  = "GetData Duck"
	TestGetDataPass3Input = "GaGa"
)

/* todo: разобраться, нужно ли держать тесты в отдельном пакете
func TestNewCreator(t *testing.T) {
	t.Run(TestNewCreatorName, func(t *testing.T) {
		got := Factory.NewCreator()
		want := &Factory.creator{}

		assert.EqualValues(t, got, want)
	})
}

func TestCreate(t *testing.T) {

	t.Run(TestCreatePass1Name, func(t *testing.T) {
		got := Factory.NewCreator().Create("Cat")
		want := &cats.cat{}

		assert.EqualValues(t, got, want)
	})
	t.Run(TestCreatePass2Name, func(t *testing.T) {
		got := Factory.NewCreator().Create("Dog")
		want := &dogs.dog{}

		assert.EqualValues(t, got, want)
	})
	t.Run(TestCreatePass3Name, func(t *testing.T) {
		got := Factory.NewCreator().Create("Duck")
		want := &ducks.duck{}

		assert.EqualValues(t, got, want)
	})
}*/

func TestGetData(t *testing.T) {

	t.Run(TestGetDataPass1Name, func(t *testing.T) {
		temp := Factory.NewCreator().Create("Cat")
		temp.SetData(TestGetDataPass1Input, "", "", 0)
		got, _, _, _ := temp.GetData()
		want := TestGetDataPass1Input

		assert.EqualValues(t, got, want)
	})
	t.Run(TestGetDataPass2Name, func(t *testing.T) {
		temp := Factory.NewCreator().Create("Dog")
		temp.SetData(TestGetDataPass2Input, "", "", 0)
		got, _, _, _ := temp.GetData()
		want := TestGetDataPass2Input

		assert.EqualValues(t, got, want)
	})
	t.Run(TestGetDataPass3Name, func(t *testing.T) {
		temp := Factory.NewCreator().Create("Duck")
		temp.SetData(TestGetDataPass3Input, "", "", 0)
		got, _, _, _ := temp.GetData()
		want := TestGetDataPass3Input

		assert.EqualValues(t, got, want)
	})
}
