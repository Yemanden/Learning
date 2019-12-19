package proxy

import (
	"fmt"
	"testing"
)

func TestProxy_SetData(t *testing.T) {
	value := []int{1, 4, 7, 10, 123, 55}

	for _, v := range value {
		prox := Proxy{}
		prox.SetData(v)
		got := prox.object.data
		want := v

		if got != want {
			fmt.Printf("%s: got [%v] want [%v]", "SetData", got, want)
		}
	}
}

func TestProxy_GetData(t *testing.T) {
	value := []int{1, 4, 7, 10, 123, 55}

	for _, v := range value {
		prox := Proxy{}
		prox.SetData(v)
		got := prox.GetData()
		want := v

		if got != want {
			fmt.Printf("%s: got [%v] want [%v]", "GetData", got, want)
		}
	}
}
