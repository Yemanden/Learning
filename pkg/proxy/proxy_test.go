package proxy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	testSetGetDataName = "SetGetData"
)

func TestSetGetData(t *testing.T) {
	t.Run(testSetGetDataName, func(t *testing.T) {
		value := []int{1, 4, 7, 10, 123, 55}

		for i := 0; i < len(value); i++ {
			p := NewProxy()
			p.SetData(value[i])
			got := p.GetData()
			want := value[i]

			assert.EqualValues(t, want, got)
		}
	})
}
