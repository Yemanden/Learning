package islands

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	testDFSPass1Name = "One easy island"
	testDFSPass2Name = "One hard island"
	testDFSPass3Name = "Very long island"
	testDFSPass4Name = "Five islands"

	testBFSPass1Name = "One easy island"
	testBFSPass2Name = "One hard island"
	testBFSPass3Name = "Very long island"
	testBFSPass4Name = "Five islands"
)

var (
	testDFSPass1Input = [][]byte{{1, 1, 0}, {1, 0, 0}, {0, 0, 0}}
	testDFSPass1Want  = 1
	testDFSPass2Input = [][]byte{{1, 1, 1}, {0, 1, 0}, {1, 1, 1}}
	testDFSPass2Want  = 1
	testDFSPass3Input = [][]byte{{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1}} // 300
	testDFSPass3Want  = 1
	testDFSPass4Input = [][]byte{{1, 1, 0, 0, 1}, {1, 0, 1, 0, 1}, {0, 0, 1, 0, 0}, {1, 0, 1, 0, 1}, {1, 0, 1, 0, 1}}
	testDFSPass4Want  = 5

	testBFSPass1Input = [][]byte{{1, 1, 0}, {1, 0, 0}, {0, 0, 0}}
	testBFSPass1Want  = 1
	testBFSPass2Input = [][]byte{{1, 1, 1}, {0, 1, 0}, {1, 1, 1}}
	testBFSPass2Want  = 1
	testBFSPass3Input = [][]byte{{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1}} // 300
	testBFSPass3Want  = 1
	testBFSPass4Input = [][]byte{{1, 1, 0, 0, 1}, {1, 0, 1, 0, 1}, {0, 0, 1, 0, 0}, {1, 0, 1, 0, 1}, {1, 0, 1, 0, 1}}
	testBFSPass4Want  = 5
)

// TestGetIslandCountDFS ...
func TestGetIslandCountDFS(t *testing.T) {
	t.Run(testDFSPass1Name, func(t *testing.T) {
		want := testDFSPass1Want

		ic := NewIslandCounter()
		got := ic.GetIslandCountDFS(testDFSPass1Input)

		assert.EqualValues(t, want, got)
	})

	t.Run(testDFSPass2Name, func(t *testing.T) {
		want := testDFSPass2Want

		ic := NewIslandCounter()
		got := ic.GetIslandCountDFS(testDFSPass2Input)

		assert.EqualValues(t, want, got)
	})

	t.Run(testDFSPass3Name, func(t *testing.T) {
		want := testDFSPass3Want

		ic := NewIslandCounter()
		got := ic.GetIslandCountDFS(testDFSPass3Input)

		assert.EqualValues(t, want, got)
	})
	t.Run(testDFSPass4Name, func(t *testing.T) {
		want := testDFSPass4Want

		ic := NewIslandCounter()
		got := ic.GetIslandCountDFS(testDFSPass4Input)

		assert.EqualValues(t, want, got)
	})
}

// TestGetIslandCountBFS ...
func TestGetIslandCountBFS(t *testing.T) {
	t.Run(testBFSPass1Name, func(t *testing.T) {
		want := testBFSPass1Want

		ic := NewIslandCounter()
		got := ic.GetIslandCountBFS(testBFSPass1Input)

		assert.EqualValues(t, want, got)
	})

	t.Run(testBFSPass2Name, func(t *testing.T) {
		want := testBFSPass2Want

		ic := NewIslandCounter()
		got := ic.GetIslandCountBFS(testBFSPass2Input)

		assert.EqualValues(t, want, got)
	})

	t.Run(testBFSPass3Name, func(t *testing.T) {
		want := testBFSPass3Want

		ic := NewIslandCounter()
		got := ic.GetIslandCountBFS(testBFSPass3Input)

		assert.EqualValues(t, want, got)
	})
	t.Run(testBFSPass4Name, func(t *testing.T) {
		want := testBFSPass4Want

		ic := NewIslandCounter()
		got := ic.GetIslandCountBFS(testBFSPass4Input)

		assert.EqualValues(t, want, got)
	})
}
