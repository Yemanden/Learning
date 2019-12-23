package set

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	TestGetValueName     = "GetValues"
	TestNewSetName       = "NewSet"
	TestAddPass1Name     = "Add Missing"
	TestAddPass2Name     = "Add Existing"
	TestRemovePass1Name  = "Remove Missing"
	TestRemovePass2Name  = "Remove Existing"
	TestUnionName        = "Union"
	TestDifferenceName   = "Difference"
	TestIntersectionName = "Intersection"
	TestSubsetPass1Name  = "Subset Missing"
	TestSubsetPass2Name  = "Subset Existing"
)

func TestGetValues(t *testing.T) {
	set := NewSet(1, 2, 3, 4, 5, 6, 7, 8, 9)

	t.Run(TestGetValueName, func(t *testing.T) {
		got := set.GetValues()
		want := map[int]bool{1: true, 2: true, 3: true, 4: true, 5: true, 6: true, 7: true, 8: true, 9: true}

		assert.EqualValues(t, got, want)
	})
}

func TestNewSet(t *testing.T) {
	t.Run(TestNewSetName, func(t *testing.T) {
		got := NewSet(1, 2, 3, 3, 2, 4).GetValues()
		want := map[int]bool{1: true, 2: true, 3: true, 4: true}

		assert.EqualValues(t, got, want)
	})
}

func TestAdd(t *testing.T) {
	set := NewSet(1, 2, 3, 4, 5, 6, 7, 9)

	t.Run(TestAddPass1Name, func(t *testing.T) {
		set.Add(8)
		got := set.GetValues()
		want := map[int]bool{1: true, 2: true, 3: true, 4: true, 5: true, 6: true, 7: true, 9: true, 8: true}

		assert.EqualValues(t, got, want)
	})

	set = NewSet(1, 2, 3, 4, 5, 6, 7, 9)

	t.Run(TestAddPass2Name, func(t *testing.T) {
		set.Add(1)
		got := set.GetValues()
		want := map[int]bool{1: true, 2: true, 3: true, 4: true, 5: true, 6: true, 7: true, 9: true}

		assert.EqualValues(t, got, want)
	})
}

func TestRemove(t *testing.T) {
	set := NewSet(1, 2, 3, 4, 5, 6, 7, 9)

	t.Run(TestRemovePass1Name, func(t *testing.T) {
		got := set.Remove(8)
		want := false

		assert.EqualValues(t, got, want)
	})
	t.Run(TestRemovePass2Name, func(t *testing.T) {
		got := set.Remove(1)
		want := true

		assert.EqualValues(t, got, want)
	})
}

func TestUnion(t *testing.T) {
	set1 := NewSet(1, 2, 3, 4)
	set2 := NewSet(3, 4, 5, 6)

	t.Run(TestUnionName, func(t *testing.T) {
		got := Union(set1, set2).GetValues()
		want := map[int]bool{1: true, 2: true, 3: true, 4: true, 5: true, 6: true}

		assert.EqualValues(t, got, want)
	})
}

func TestDifference(t *testing.T) {
	set1 := NewSet(1, 2, 3, 4)
	set2 := NewSet(3, 4, 5, 6)

	t.Run(TestDifferenceName, func(t *testing.T) {
		got := Difference(set1, set2).GetValues()
		want := map[int]bool{1: true, 2: true}

		assert.EqualValues(t, got, want)
	})
}

func TestIntersection(t *testing.T) {
	set1 := NewSet(1, 2, 3, 4)
	set2 := NewSet(3, 4, 5, 6)

	t.Run(TestIntersectionName, func(t *testing.T) {
		got := Intersection(set1, set2).GetValues()
		want := map[int]bool{3: true, 4: true}

		assert.EqualValues(t, got, want)
	})
}

func TestSubset(t *testing.T) {
	set1 := NewSet(1, 2, 3, 4)
	set2 := NewSet(3, 4, 5, 6)
	set3 := NewSet(1, 2, 3, 4, 5)

	t.Run(TestSubsetPass1Name, func(t *testing.T) {
		got := Subset(set1, set2)
		want := false

		assert.EqualValues(t, got, want)
	})
	t.Run(TestSubsetPass2Name, func(t *testing.T) {
		got := Subset(set1, set3)
		want := true

		assert.EqualValues(t, got, want)
	})
}
