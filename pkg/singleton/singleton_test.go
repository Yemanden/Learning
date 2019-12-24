package singleton

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	TestGetInstanceName = "GetInstance"
	TestAddName         = "Add"
	TestGetDataName     = "GetData"
)

func TestGetInstance(t *testing.T) {
	t.Run(TestGetInstanceName, func(t *testing.T) {
		got := GetInstance()
		want := instance

		assert.EqualValues(t, got, want)
	})
}

func TestAdd(t *testing.T) {
	t.Run(TestAddName, func(t *testing.T) {
		testVariables := []int{-500, 1, 3, 5, 7, 10, 1567}
		for _, v := range testVariables {
			instance = &singleton{}
			instance.data = 0

			single := GetInstance()
			single.Add(v)

			got := instance.data
			want := v

			assert.EqualValues(t, got, want)
		}
	})
}

func TestGetData(t *testing.T) {
	t.Run(TestGetDataName, func(t *testing.T) {
		testVariables := []int{-500, 1, 3, 5, 7, 10, 1567}
		for v := range testVariables {
			instance = &singleton{}
			instance.data = v

			single := GetInstance()
			got := single.GetData()

			want := v

			assert.EqualValues(t, got, want)
		}
	})
}
