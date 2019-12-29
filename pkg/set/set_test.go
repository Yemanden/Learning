package set

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	testGetValueName = "GetValues"
	testNewSetName   = "NewSet"

	testAddPass1Name = "Add Missing"
	testAddPass2Name = "Add Existing"

	testRemovePass1Name = "Remove Missing"
	testRemovePass1Want = false

	testRemovePass2Name = "Remove Existing"
	testRemovePass2Want = true

	testUnionName = "Union"

	testDifferenceName = "Difference"

	testIntersectionName = "Intersection"

	testSubsetPass1Name = "Subset Missing"
	testSubsetPass1Want = false

	testSubsetPass2Name = "Subset Existing"
	testSubsetPass2Want = true
)

func TestGetValues(t *testing.T) {
	set := NewSet()
	for i := 1; i < 10; i++ {
		set.Add(i)
	}

	t.Run(testGetValueName, func(t *testing.T) {
		got := set.GetValues()
		want := map[int]bool{1: true, 2: true, 3: true, 4: true, 5: true, 6: true, 7: true, 8: true, 9: true}

		assert.EqualValues(t, got, want)
	})
}

func TestNewSet(t *testing.T) {
	t.Run(testNewSetName, func(t *testing.T) {
		got := NewSet()
		want := &set{}

		assert.EqualValues(t, got, want)
	})
}

func TestAdd(t *testing.T) {
	set := NewSet()
	for i := 1; i < 10; i++ {
		set.Add(i)
	}

	t.Run(testAddPass1Name, func(t *testing.T) {
		set.Add(10)
		got := set.GetValues()
		want := map[int]bool{1: true, 2: true, 3: true, 4: true, 5: true, 6: true, 7: true, 8: true, 9: true, 10: true}

		assert.EqualValues(t, got, want)
	})

	set = NewSet()
	for i := 1; i < 10; i++ {
		set.Add(i)
	}

	t.Run(testAddPass2Name, func(t *testing.T) {
		set.Add(1)
		got := set.GetValues()
		want := map[int]bool{1: true, 2: true, 3: true, 4: true, 5: true, 6: true, 7: true, 8: true, 9: true}

		assert.EqualValues(t, got, want)
	})
}

func TestRemove(t *testing.T) {
	set := NewSet()
	for i := 1; i < 10; i++ {
		set.Add(i)
	}

	t.Run(testRemovePass1Name, func(t *testing.T) {
		got := set.Remove(10)
		want := testRemovePass1Want

		assert.EqualValues(t, got, want)
	})
	t.Run(testRemovePass2Name, func(t *testing.T) {
		got := set.Remove(1)
		want := testRemovePass2Want

		assert.EqualValues(t, got, want)
	})
}

func TestUnion(t *testing.T) {
	set1 := NewSet()
	set2 := NewSet()
	for i := 1; i < 4; i++ {
		set1.Add(i)
		set2.Add(2 * i)
	}

	t.Run(testUnionName, func(t *testing.T) {
		got := set1.Union(set2).GetValues()
		want := map[int]bool{1: true, 2: true, 3: true, 4: true, 6: true}

		assert.EqualValues(t, got, want)
	})
}

func TestDifference(t *testing.T) {
	set1 := NewSet()
	set2 := NewSet()
	for i := 1; i < 4; i++ {
		set1.Add(i)
		set2.Add(2 * i)
	}

	t.Run(testDifferenceName, func(t *testing.T) {
		got := set1.Difference(set2).GetValues()
		want := map[int]bool{1: true, 3: true}

		assert.EqualValues(t, got, want)
	})
}

func TestIntersection(t *testing.T) {
	set1 := NewSet()
	set2 := NewSet()
	for i := 1; i < 4; i++ {
		set1.Add(i)
		set2.Add(2 * i)
	}

	t.Run(testIntersectionName, func(t *testing.T) {
		got := set1.Intersection(set2).GetValues()
		want := map[int]bool{2: true}

		assert.EqualValues(t, got, want)
	})
}

func TestSubset(t *testing.T) {
	set1 := NewSet()
	set2 := NewSet()
	for i := 1; i < 4; i++ {
		set1.Add(i)
		set2.Add(2 * i)
	}
	set3 := NewSet()
	for i := 1; i < 10; i++ {
		set3.Add(i)
	}

	t.Run(testSubsetPass1Name, func(t *testing.T) {
		got := set1.Subset(set2)
		want := testSubsetPass1Want

		assert.EqualValues(t, got, want)
	})
	t.Run(testSubsetPass2Name, func(t *testing.T) {
		got := set1.Subset(set3)
		want := testSubsetPass2Want

		assert.EqualValues(t, got, want)
	})
}
