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

		for i := 0; i < len(value); i++ {
			proxy := NewProxy()
			proxy.SetData(value[i])
			got := proxy.GetData()
			want := value[i]

			assert.EqualValues(t, got, want)
		}
	})
}
