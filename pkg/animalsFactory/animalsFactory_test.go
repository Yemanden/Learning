package animalsFactory

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	testNewCreatorName = "NewCreator"

	testCreatePass1Name = "Create Cat"

	testCreatePass2Name = "Create Dog"

	testCreatePass3Name = "Create Duck"

	testGetDataPass1Name  = "GetData Cat"
	testGetDataPass1Input = "Murka"

	testGetDataPass2Name  = "GetData Dog"
	testGetDataPass2Input = "Sharik"

	testGetDataPass3Name  = "GetData Duck"
	testGetDataPass3Input = "GaGa"
)

// TestNewCreator ...
func TestNewCreator(t *testing.T) {
	t.Run(testNewCreatorName, func(t *testing.T) {
		got := NewCreator()
		want := &creator{}

		assert.EqualValues(t, got, want)
	})
}

// TestCreate ...
func TestCreate(t *testing.T) {

	t.Run(testCreatePass1Name, func(t *testing.T) {
		got := NewCreator().Create("Cat")
		want := &cat{}

		assert.EqualValues(t, got, want)
	})

	t.Run(testCreatePass2Name, func(t *testing.T) {
		got := NewCreator().Create("Dog")
		want := &dog{}

		assert.EqualValues(t, got, want)
	})

	t.Run(testCreatePass3Name, func(t *testing.T) {
		got := NewCreator().Create("Duck")
		want := &duck{}

		assert.EqualValues(t, got, want)
	})
}

// TestGetData ...
func TestGetData(t *testing.T) {

	t.Run(testGetDataPass1Name, func(t *testing.T) {
		temp := NewCreator().Create("Cat")
		temp.SetData(testGetDataPass1Input, "", "", 0)
		got, _, _, _ := temp.GetData()
		want := testGetDataPass1Input

		assert.EqualValues(t, want, got)
	})

	t.Run(testGetDataPass2Name, func(t *testing.T) {
		temp := NewCreator().Create("Dog")
		temp.SetData(testGetDataPass2Input, "", "", 0)
		got, _, _, _ := temp.GetData()
		want := testGetDataPass2Input

		assert.EqualValues(t, want, got)
	})

	t.Run(testGetDataPass3Name, func(t *testing.T) {
		temp := NewCreator().Create("Duck")
		temp.SetData(testGetDataPass3Input, "", "", 0)
		got, _, _, _ := temp.GetData()
		want := testGetDataPass3Input

		assert.EqualValues(t, want, got)
	})
}
