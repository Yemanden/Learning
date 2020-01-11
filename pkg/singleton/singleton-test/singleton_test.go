package singleton_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Yemanden/Learning/pkg/singleton/singleton"
)

const (
	testGetInstanceName = "Test GetInstance"

	testDataName  = "Test Data"
	testDataInput = 5
)

// TestGetInstance ...
func TestGetInstance(t *testing.T) {
	t.Run(testGetInstanceName, func(t *testing.T) {
		got := singleton.GetInstance()
		got.Add(3)
		want := singleton.GetInstance()

		assert.EqualValues(t, want, got)
	})
}

// TestData ...
func TestData(t *testing.T) {
	t.Run(testDataName, func(t *testing.T) {
		single := singleton.GetInstance()
		tmp := single.GetData()
		single.Add(testDataInput)

		got := single.GetData()
		want := tmp + testDataInput

		assert.EqualValues(t, want, got)
	})
}
