package Singleton

import (
	"fmt"
	"testing"
)

func TestGetInstance(t *testing.T) {

	got := GetInstance()
	want := instance

	if got != want {
		fmt.Printf("%s: got [%v] want [%v]", "Instance", got, want)
	}
}

func TestSingleton_Add(t *testing.T) {

	testVariables := []int{-500, 1, 3, 5, 7, 10, 1567}

	for _, v := range testVariables {
		instance.data = 0

		single := GetInstance()
		single.Add(v)
		got := single.data

		want := v

		if got != want {
			fmt.Printf("%s: got [%v] want [%v] \n", "Add", got, want)
		}
	}
}

func TestSingleton_GetData(t *testing.T) {

	testVariables := []int{-500, 1, 3, 5, 7, 10, 1567}

	for v := range testVariables {
		instance.data = 0

		single := GetInstance()
		single.Add(v)
		got := single.GetData()

		want := v

		if got != want {
			fmt.Printf("%s: got [%v] want [%v]\n", "GetData", got, want)
		}
	}
}
