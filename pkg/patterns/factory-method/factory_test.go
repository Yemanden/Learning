package factory_method

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Yemanden/Learning/pkg/patterns/factory-method/factory"
)

const (
	testGetDataPass1Name  = "GetData Cat"
	testGetDataPass1Input = "Murka"

	testGetDataPass2Name  = "GetData Dog"
	testGetDataPass2Input = "Sharik"

	testGetDataPass3Name  = "GetData Duck"
	testGetDataPass3Input = "GaGa"
)

// TestGetData ...
func TestGetData(t *testing.T) {

	t.Run(testGetDataPass1Name, func(t *testing.T) {
		temp := factory.NewCreator().Create("Cat")
		temp.SetData(testGetDataPass1Input, "", "", 0)
		got, _, _, _ := temp.GetData()
		want := testGetDataPass1Input

		assert.EqualValues(t, want, got)
	})

	t.Run(testGetDataPass2Name, func(t *testing.T) {
		temp := factory.NewCreator().Create("Dog")
		temp.SetData(testGetDataPass2Input, "", "", 0)
		got, _, _, _ := temp.GetData()
		want := testGetDataPass2Input

		assert.EqualValues(t, want, got)
	})

	t.Run(testGetDataPass3Name, func(t *testing.T) {
		temp := factory.NewCreator().Create("Duck")
		temp.SetData(testGetDataPass3Input, "", "", 0)
		got, _, _, _ := temp.GetData()
		want := testGetDataPass3Input

		assert.EqualValues(t, want, got)
	})
}
