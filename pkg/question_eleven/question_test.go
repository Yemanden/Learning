package questioneleven

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	testNumIslandPass1Name = "One easy island"
	testNumIslandPass2Name = "One hard island"
	testNumIslandPass3Name = "Very long island"
	testNumIslandPass4Name = "Five islands"
)

var (
	testNumIslandPass1Input = [][]byte{{1, 1, 0}, {1, 0, 0}, {0, 0, 0}}
	testNumIslandPass1Want  = 1
	testNumIslandPass2Input = [][]byte{{1, 1, 1}, {0, 1, 0}, {1, 1, 1}}
	testNumIslandPass2Want  = 1
	testNumIslandPass3Input = [][]byte{{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1}} // 300
	testNumIslandPass3Want  = 1
	testNumIslandPass4Input = [][]byte{{1, 1, 0, 0, 1}, {1, 0, 1, 0, 1}, {0, 0, 1, 0, 0}, {1, 0, 1, 0, 1}, {1, 0, 1, 0, 1}}
	testNumIslandPass4Want  = 5
)

func TestNumIslands(t *testing.T) {
	t.Run(testNumIslandPass1Name, func(t *testing.T) {
		got := numIslands(testNumIslandPass1Input)
		want := testNumIslandPass1Want

		assert.EqualValues(t, got, want)
	})

	t.Run(testNumIslandPass2Name, func(t *testing.T) {
		got := numIslands(testNumIslandPass2Input)
		want := testNumIslandPass2Want

		assert.EqualValues(t, got, want)
	})

	t.Run(testNumIslandPass3Name, func(t *testing.T) {
		got := numIslands(testNumIslandPass3Input)
		want := testNumIslandPass3Want

		assert.EqualValues(t, got, want)
	})
	t.Run(testNumIslandPass4Name, func(t *testing.T) {
		got := numIslands(testNumIslandPass4Input)
		want := testNumIslandPass4Want

		assert.EqualValues(t, got, want)
	})
}
