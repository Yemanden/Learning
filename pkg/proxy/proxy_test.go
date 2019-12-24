package proxy

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	TestSetGetDataName = "SetGetData"
)

func TestSetGetData(t *testing.T) {
	t.Run(TestSetGetDataName, func(t *testing.T) {
		value := []int{1, 4, 7, 10, 123, 55}

		for _, v := range value {
			proxy := NewProxy()
			proxy.SetData(v)
			got := proxy.GetData()
			want := v

			assert.EqualValues(t, got, want)
		}
	})
}
